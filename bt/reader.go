package bt

import (
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os/exec"
	"strings"

	"cloud.google.com/go/bigtable"
	"github.com/golang/protobuf/proto"
	"github.com/klauspost/compress/zstd"
	pbsolv1 "github.com/streamingfast/firehose-solana/pb/sf/solana/type/v1"
	"go.uber.org/zap"
)

type RowType string

const (
	RowTypeProto RowType = "proto"
	RowTypeBin   RowType = "bin"
)

func explodeRow(row bigtable.Row) (*big.Int, RowType, []byte) {
	el := row["x"][0]
	var rowType RowType
	if strings.HasSuffix(el.Column, "proto") {
		rowType = RowTypeProto
	} else {
		rowType = RowTypeBin
	}
	blockNum, _ := new(big.Int).SetString(el.Row, 16)
	return blockNum, rowType, el.Value
}

func (r *Client) processRow(row bigtable.Row) (*pbsolv1.Block, *zap.Logger, error) {
	blockNum, rowType, rowCnt := explodeRow(row)
	zlogger := r.logger.With(
		zap.Uint64("block_num", blockNum.Uint64()),
		zap.String("row_type", string(rowType)),
		zap.String("row_key", row.Key()),
	)

	var cnt []byte
	var err error

	switch rowType {
	case RowTypeBin:
		cnt, err = externalBinToProto(rowCnt, "solana-bigtable-decoder", "--hex")
		if err != nil {
			return nil, zlogger, fmt.Errorf("unable get decode bin with external command 'solana-bigtable-decoder'  %s: %w", blockNum.String(), err)
		}
	default:
		cnt, err = r.decompress(rowCnt)
		if err != nil {
			return nil, zlogger, fmt.Errorf("unable to decompress block %s (uncompresse length %d): %w", blockNum.String(), len(rowCnt), err)
		}

	}
	zlogger.Debug("found bigtable row",
		zap.Stringer("blk_num", blockNum),
		zap.String("key", row.Key()),
		zap.Int("compressed_length", len(rowCnt)),
		zap.Int("uncompressed_length", len(cnt)),
		zap.String("row_type", string(rowType)),
	)

	blk := &pbsolv1.Block{}
	if err := proto.Unmarshal(cnt, blk); err != nil {
		return nil, zlogger, fmt.Errorf("unable to unmarshall confirmed block: %w", err)
	}
	blk.Slot = blockNum.Uint64()

	// horrible tweaks
	switch blk.Blockhash {
	case "Goi3t9JjgDkyULZbM2TzE5QqHP1fPeMcHNaXNFBCBv1v":
		zlogger.Warn("applying horrible tweak to block Goi3t9JjgDkyULZbM2TzE5QqHP1fPeMcHNaXNFBCBv1v")
		if blk.PreviousBlockhash == "11111111111111111111111111111111" {
			blk.PreviousBlockhash = "HQEr9qcbUVBt2okfu755FdJvJrPYTSpzzmmyeWTj5oau"
		}
	case "6UFQveZ94DUKGbcLFoyayn1QwthVfD3ZqvrM2916pHCR":
		zlogger.Warn("applying horrible tweak to block 63,072,071")
		if blk.PreviousBlockhash == "11111111111111111111111111111111" {
			blk.PreviousBlockhash = "7cLQx2cZvyKbGoMuutXEZ3peg3D21D5qbX19T5V1XEiK"
		}
	case "Fqbm7QvCTYnToXWcCw6nbkWhMmXx2Nv91LsXBrKraB43":
		zlogger.Warn("applying horrible tweak to block 53,135,959")
		if blk.PreviousBlockhash == "11111111111111111111111111111111" {
			blk.PreviousBlockhash = "RfXUrekgajPSb1R4CGFJWNaHTnB6p53Tzert4gouj2u"
		}
	case "ABp9G2NaPzM6kQbeyZYCYgdzL8JN9AxSSbCQG2X1K9UF":
		zlogger.Warn("applying horrible tweak to block 46,223,993")
		if blk.PreviousBlockhash == "11111111111111111111111111111111" {
			blk.PreviousBlockhash = "9F2C7TGqUpFu6krd8vQbUv64BskrneBSgY7U2QfrGx96"
		}
	case "ByUxmGuaT7iQS9qGS8on5xHRjiHXcGxvwPPaTGZXQyz7":
		zlogger.Warn("applying horrible tweak to block 61,328,766")
		if blk.PreviousBlockhash == "11111111111111111111111111111111" {
			blk.PreviousBlockhash = "J6rRToKMK5DQDzVLqo7ibL3snwBYtqkYnRnQ7vXoUSEc"
		}
	}

	return blk, zlogger, nil
}

func externalBinToProto(in []byte, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to run command: %w", err)
	}

	inString := hex.EncodeToString(in)

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, inString)
	}()

	outCntHex, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed get command output: %w", err)
	}
	outHex := string(outCntHex)
	if strings.HasPrefix(outHex, "0x") {
		outHex = outHex[2:]
	}
	outHex = strings.TrimRight(outHex, "\n")
	cnt, err := hex.DecodeString(outHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode output string %q: %w", outHex, err)
	}
	return cnt, nil
}

func (r *Client) decompress(in []byte) (out []byte, err error) {
	switch in[0] {
	case 0:
		r.logger.Debug("no compression found")
		out = in[4:]
	case 1:
		r.logger.Debug("bzip2 compression")
		// bzip2
		out, err = ioutil.ReadAll(bzip2.NewReader(bytes.NewBuffer(in[4:])))
		if err != nil {
			return nil, fmt.Errorf("bzip2 decompress: %w", err)
		}
	case 2:
		r.logger.Debug("gzip compression")
		// gzip
		reader, err := gzip.NewReader(bytes.NewBuffer(in[4:]))
		if err != nil {
			return nil, fmt.Errorf("gzip reader: %w", err)
		}
		out, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("gzip decompress: %w", err)
		}
	case 3:
		r.logger.Debug("zstd compression")
		// zstd
		var dec *zstd.Decoder
		dec, err = zstd.NewReader(nil)
		if err != nil {
			return nil, fmt.Errorf("zstd reader: %w", err)
		}
		out, err = dec.DecodeAll(in[4:], out)
		if err != nil {
			return nil, fmt.Errorf("zstd decompress: %w", err)

		}
	default:
		return nil, fmt.Errorf("unsupported compression scheme for a block %d", in[0])
	}
	return
}
