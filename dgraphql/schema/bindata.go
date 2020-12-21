// Code generated by go-bindata. DO NOT EDIT.
// sources:
// query.graphql (815B)
// query_alpha.graphql (14B)
// registered_token.graphql (208B)
// schema.graphql (59B)
// serum_fill.graphql (1.463kB)
// serum_instruction.graphql (2.797kB)
// serum_market.graphql (110B)
// subscription.graphql (94B)
// subscription_alpha.graphql (21B)
// token.graphql (180B)

package schema

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x4d\x6b\x1b\x41\x0c\xbd\xef\xaf\x78\xde\x4b\x13\x08\x21\x85\xd2\x83\xaf\x81\xd2\x0f\xe8\x87\x9d\x9e\x4a\x0f\xe3\x1d\xad\x2d\x3c\x1e\x6d\x25\xad\xcd\x52\xf2\xdf\xcb\xac\x99\x40\x4c\x4a\x6e\x23\xad\xde\x7b\x7a\x6f\xe5\xd3\x40\xf8\x31\x92\x4e\xf8\xdb\x00\x2e\x7b\xca\xb6\xc4\xaf\x87\xf2\x58\x91\x0d\x92\x8d\x16\xbf\x17\x4d\xfd\x78\x15\x62\x54\x32\x5b\x62\xed\xca\x79\xbb\xb8\x5e\xe2\xd9\x70\x99\x54\xda\xb2\x39\x29\xc5\x87\x4a\xb8\x7a\xde\xba\xa0\xbe\x00\xbc\x24\xf2\x1f\x82\x82\x6e\xdb\x76\x45\x3e\x6a\x46\x0a\xe6\x78\x7b\x77\x87\x9e\x53\x32\xf4\xa2\xf0\x1d\x61\xcb\x47\xca\x08\x5d\x27\x63\xf6\x37\x86\x61\xdc\x24\xee\xb0\xa7\xe9\x16\x67\xa4\x81\x55\xe9\x48\x6a\xbc\x49\x04\xc9\x69\x42\x0c\x1e\x60\x82\xc4\x47\x3a\x17\x59\x1c\x13\x39\x86\xa0\x0e\xe9\x11\x60\x49\x1c\xbe\x0b\x0e\x36\x70\x9e\xd5\x54\xc4\x61\xe4\x38\x71\x4a\x33\x66\x43\xd0\x59\x86\x22\x4e\x3b\xca\xe8\x42\x4a\x14\x6f\xdb\xb6\x6d\x00\x23\x1d\x0f\x1f\x38\xa5\x8f\x6c\x2e\x3a\x5d\x0d\xe3\xe6\x0b\x4d\x4f\xe6\x6f\x70\x08\xba\x27\xaf\x8d\xeb\x25\xd6\x15\x72\x2f\x39\x53\xe7\x2c\x79\xd1\x3c\x36\x4d\xdb\xb6\xf7\xa3\x9a\xa8\x41\xe9\xcf\xc8\x4a\x11\x2e\xe8\x24\x3b\xe7\x91\x40\xec\x3b\xd2\x92\xcb\x29\x68\x84\x28\x36\xa1\xdb\x97\xb7\xa1\x57\x39\x20\x20\xb1\xcd\xde\x86\xb0\xe5\x1c\x9c\x22\x28\xd1\x81\xb2\x5b\xd9\x76\xbe\x97\xef\x61\x4b\x9f\x72\x2f\xe5\x62\xe6\xf4\xbb\x59\xb3\xa0\x8a\xff\x9e\xd5\xbc\xa2\x6a\xb3\xd0\xde\x60\x34\x02\x7b\x59\xc9\x28\x68\xb7\xab\x91\xc9\x30\x88\xb1\x13\x22\xeb\xd9\xce\x39\x1a\xc0\x3c\xa8\x9f\x3d\x3d\x05\xd2\xbc\xac\x3b\xff\xfc\x57\x64\x6b\x12\x95\x9e\x72\xbc\x24\x7f\x6c\x1a\xeb\x42\x0a\x8a\x9f\x9c\xfd\xfd\xbb\x5a\x7d\x5e\x7f\xfb\xfa\x2f\x00\x00\xff\xff\x65\x72\xfd\x54\x2f\x03\x00\x00")

func queryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphql,
		"query.graphql",
	)
}

func queryGraphql() (*asset, error) {
	bytes, err := queryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphql", size: 815, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x25, 0x1a, 0x7, 0xca, 0x4c, 0x13, 0x65, 0x47, 0x8b, 0xe7, 0x95, 0xed, 0x31, 0x7b, 0x90, 0x28, 0x6d, 0xf, 0xb9, 0xa4, 0xa3, 0xb1, 0xce, 0xaa, 0x7a, 0x3b, 0xdd, 0xe1, 0x73, 0xbb, 0x28, 0x4b}}
	return a, nil
}

var _query_alphaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\x54\xa8\xae\xe5\x02\x04\x00\x00\xff\xff\x76\xca\x60\x3d\x0e\x00\x00\x00")

func query_alphaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_query_alphaGraphql,
		"query_alpha.graphql",
	)
}

func query_alphaGraphql() (*asset, error) {
	bytes, err := query_alphaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query_alpha.graphql", size: 14, mode: os.FileMode(0644), modTime: time.Unix(1608153826, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0x58, 0xf, 0x86, 0xb0, 0x43, 0x44, 0x23, 0x5f, 0x97, 0xd3, 0xde, 0x25, 0xbd, 0x4b, 0x29, 0x22, 0xad, 0x9b, 0x95, 0xef, 0x8, 0x81, 0x45, 0x11, 0x3a, 0x12, 0x62, 0xab, 0x4c, 0x93, 0xf8}}
	return a, nil
}

var _registered_tokenGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8b\x31\x0e\xc2\x30\x0c\x45\xf7\x9e\xc2\xdc\x01\x31\x74\x63\x64\x2d\x70\x80\x96\x7c\x82\x45\x62\x47\xb1\x2b\x14\x10\x77\x47\x30\x65\x60\x7c\xff\xbd\xef\xad\x80\x26\x44\x36\x47\x45\x38\xe9\x1d\x32\xc1\x8a\x8a\x81\x5e\x03\xd1\x1c\x42\x85\xd9\x48\x47\xaf\x2c\x71\x33\x10\x65\x16\xdf\xaf\x7e\xd3\xca\xde\x7a\x71\xad\xc0\x13\x7f\x95\xad\xa5\xa4\x36\xd2\x99\xc5\x77\xdb\xef\x12\x70\xe1\x3c\x27\x1b\xe9\x20\xfe\x4b\x5a\x5e\x34\xf5\x27\x99\x33\x7a\x4e\x1a\xb5\xe7\x07\x16\x63\xef\x92\xf7\xf0\x09\x00\x00\xff\xff\x65\x08\x1e\x7a\xd0\x00\x00\x00")

func registered_tokenGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_registered_tokenGraphql,
		"registered_token.graphql",
	)
}

func registered_tokenGraphql() (*asset, error) {
	bytes, err := registered_tokenGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "registered_token.graphql", size: 208, mode: os.FileMode(0644), modTime: time.Unix(1608153758, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa7, 0xad, 0xf8, 0x27, 0xfd, 0xab, 0x5a, 0x68, 0x60, 0x15, 0x89, 0x72, 0xbe, 0xcc, 0xcf, 0x4f, 0xcf, 0x8c, 0x4c, 0x3b, 0xe6, 0xcb, 0x89, 0x83, 0xe7, 0x3c, 0x5d, 0x1a, 0xf6, 0xe1, 0xec, 0xb5}}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xe2\xd2\xa4\xe2\xe4\xa2\xcc\x82\x92\xcc\xfc\x3c\x2b\x85\x60\x24\x1e\x57\x2d\x17\x20\x00\x00\xff\xff\x52\xd9\x58\xe5\x3b\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 59, mode: os.FileMode(0644), modTime: time.Unix(1608156028, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5b, 0xba, 0x8f, 0x74, 0x29, 0x4e, 0xaf, 0x41, 0x66, 0x2b, 0x4e, 0x31, 0x85, 0x84, 0x19, 0x59, 0x50, 0x81, 0xda, 0x72, 0x50, 0x56, 0xaf, 0xe3, 0xd8, 0xb7, 0x35, 0xbc, 0xd1, 0x85, 0xf3, 0xc6}}
	return a, nil
}

var _serum_fillGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4d\x4f\xc4\x36\x14\xbc\xe7\x57\x3c\x72\x01\xa4\xed\xae\xd4\x52\x0e\xb9\x2d\x94\x56\x88\xae\x54\xc8\x56\x3d\x20\x0e\x4e\x3c\x49\x5c\x1c\x3b\xd8\xcf\x6c\xa3\x8a\xff\x5e\x39\xc9\xb2\x1f\x2c\xea\xc9\xbb\xf6\xcc\xbc\x99\xe7\x17\x27\x69\x9a\xae\x1b\xd0\xad\x35\x06\x25\x2b\x6b\x88\xfb\x0e\x54\x59\x47\x82\x72\xb8\xd0\xfe\xaa\xb4\x9e\x53\x0e\x10\x37\xa0\xe7\xdf\x9c\xe8\x9a\xc7\xdf\xa9\xb4\x0e\x54\x5a\x53\xa2\x63\xff\x72\xd1\x30\x77\x3e\x5b\x2c\xa4\x2d\xfd\x5c\x56\xc1\x63\xae\xec\xa2\x0e\x4a\xc2\x2f\x22\xf6\x87\x2d\x76\x51\x47\x85\x37\xbd\xb8\x4c\xd3\x34\x19\xaa\x7d\xd6\xd9\xb3\xf1\x6f\x42\x94\x2e\x49\x2b\xcf\x64\x2b\x82\xac\xe1\x89\xed\x0e\x9b\x26\x34\xee\x66\xf4\xfc\xb9\x79\x27\x6b\xbc\x9c\x25\x91\x7b\x6f\x2a\xeb\x5a\x31\x66\xb2\x24\x94\xa4\x4e\xd4\xca\x0c\x3b\x91\xdc\x89\x1a\x11\x94\xd1\x1f\xd3\xaf\xb3\xe4\x23\x89\x1d\x49\x96\xe4\x95\xa9\xf5\x9e\x33\x82\x46\x0b\x33\x78\x89\x7d\x28\x77\x4e\xa3\xc5\x79\xf2\x35\x4c\xf4\x32\xc6\x48\xd3\xc7\x00\xd7\x53\x19\x9c\xb7\x6e\x46\x42\x6b\xbb\x51\xa6\xa6\xde\x86\xe8\xad\xb4\x86\x95\x09\xa0\x0a\x5c\x36\xf1\xc0\xc1\x07\xcd\x9e\xf0\x0e\x43\xa2\x62\x38\x52\x86\xe1\x5c\xe8\x86\x9a\xb6\x8a\x5c\xb7\x67\x63\x46\x1b\xc5\x8d\x0d\x4c\xad\xf2\xd1\x3d\x09\x2a\x20\x78\x1e\x8d\xd1\x54\x3a\xa3\x9c\x9d\x32\xf5\xd8\xa1\xf1\xea\x77\x11\x6d\xf1\x37\xca\x2d\xc1\x58\x89\x6c\x77\x18\x3b\x73\x18\x6f\x1b\x2d\x6a\x58\x27\xe1\x28\x18\xf5\x16\x40\x4a\xc2\xb0\xaa\x14\xdc\x30\x46\x9b\x46\x95\x0d\x71\xa3\xfc\x1e\x57\x79\x2a\x6c\x30\x92\xd8\x4e\xf5\x06\x89\x7b\x79\xca\xa1\xdd\x18\xb8\x73\x4f\x5d\x28\xb4\x2a\xe9\x15\xfd\x20\x7c\x28\x39\x23\x35\xc7\x7c\xb8\x1b\x6b\xe2\xac\x0a\xa6\x22\xf4\x64\x1d\x79\x68\xbd\x63\xb4\xc2\xbd\x62\x9b\xb2\x0b\xc5\x03\xfa\x53\x45\x47\xd8\xb7\x09\xce\xfd\x14\xfa\x6b\x92\x91\x39\xf5\x6e\x35\xfc\xd9\x0a\xff\x35\x28\x79\x25\x71\xa2\x21\x55\x1c\x0d\x28\x6e\xe0\xe8\xe6\xfe\x17\xba\x88\xbe\xe1\x2e\x63\x84\x65\xfe\x40\x17\x45\xe8\xe1\x2e\xa7\x2a\x51\x24\xa3\x5c\x49\xac\xfb\x0e\xfb\xce\xd9\xbe\xc2\x1c\xeb\x7b\x16\x8e\x3d\x55\xce\xb6\x93\x40\x21\x3c\xd6\x11\x9a\xd1\xb0\xfc\xaf\x04\x8c\xf4\xbb\x94\x6f\xc1\xf2\xb7\x7c\x13\xda\x02\x2e\x4e\xa9\xb6\x7c\xac\x53\x29\xad\x21\xc7\x9b\x1a\x5a\x18\x07\x77\x46\x0e\x9d\x83\x87\x61\x3f\x1c\x75\x70\x25\x0c\x8b\x1a\xdb\x4f\x6e\x04\x1f\x89\x49\x25\x27\x43\xda\xf2\xad\x0d\x86\x33\xfa\x53\x19\xbe\xbe\xda\xf7\xd3\x39\x55\x82\xba\xf8\x06\x7c\xce\x01\xfe\x11\x25\xef\x3d\x71\xd3\x40\x44\xe4\x29\x89\x11\x08\xac\xe3\x5c\x6b\xbc\x43\x1f\xe9\x55\x3b\x91\x6a\x84\x65\x07\xa4\xe1\x71\x81\x09\xed\xa1\x54\xfc\x8a\x6e\x96\xf9\x5d\x42\x94\x3f\xad\x7e\x1c\x97\x9f\xc6\xe5\x6a\x5c\x7e\x1e\x97\xeb\x84\x68\x95\x3f\xad\x92\x8f\xe4\xbf\x00\x00\x00\xff\xff\x59\x1e\x62\xc0\xb7\x05\x00\x00")

func serum_fillGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_fillGraphql,
		"serum_fill.graphql",
	)
}

func serum_fillGraphql() (*asset, error) {
	bytes, err := serum_fillGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_fill.graphql", size: 1463, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0xa5, 0xfc, 0x36, 0x8b, 0x6b, 0xb3, 0x5b, 0x83, 0x8e, 0x2, 0xee, 0xc7, 0x9f, 0xda, 0x67, 0x27, 0xb6, 0xc0, 0xa6, 0xfb, 0x65, 0xf7, 0x9a, 0x81, 0x15, 0x1b, 0x60, 0xcf, 0x0, 0x58, 0x6d}}
	return a, nil
}

var _serum_instructionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdd\x4e\xe3\x3a\x10\xbe\xcf\x53\x98\x77\x38\x3a\x17\x91\xce\x45\xff\x90\x22\xfa\x47\x53\x0e\x5a\x21\x84\xdc\x64\x68\x2d\x1c\x3b\xd8\x63\xa0\xec\xf2\xee\x2b\x3b\x4d\x63\x27\x69\x97\xdd\x25\x17\x95\xf2\xcd\x78\x32\xf3\x7d\x33\xe3\x46\xb8\x2f\x81\xa4\xa0\x4c\x91\x08\x8d\xca\x64\xc8\xa4\x58\x81\x2e\xa5\xd0\x40\xbe\x47\x84\xa0\x7a\x4b\xd9\x56\x50\x34\x0a\x62\x92\xa2\x62\x62\x7b\x51\xe1\x13\xa5\xa4\x8a\xc9\x50\x4a\x0e\x54\x58\x90\x35\x41\xe2\x4e\xd8\xe8\x23\x8a\x8c\x60\x52\x74\x2c\xe4\x3f\x72\x23\x72\xc8\x64\x0e\xb9\x0f\xff\xa8\x3d\x19\x32\xca\xd9\x3b\xcc\xa8\x7a\x02\xac\xf1\x39\xbc\x2e\x54\x0e\xaa\x7e\x9f\x51\xcc\x76\x01\x32\x92\x42\x9b\x02\x26\x2f\x20\x50\x1f\x41\x2a\x32\xe0\x81\x5f\x0a\x88\x1c\x2e\x8d\xc8\xfb\xbc\x86\xfb\x11\x67\x20\x30\xc9\xa3\x8a\xb0\xde\x64\x2d\x59\xa5\x92\x5b\x45\x8b\x64\x9c\x88\x1c\xde\x62\x72\xc3\x04\xfe\xfb\x8f\x65\x86\x66\x99\x34\x02\x47\xf6\x27\x26\xa4\x6b\xd1\x16\xb5\xcf\xdd\xc1\x74\x6f\x8d\x39\x45\x3a\x05\xb1\xc5\x5d\x65\x6e\x8e\x59\x4b\x7d\xc4\x3e\x8d\x32\x50\xc9\xd2\xb1\x7c\x44\x81\xdc\x21\xa7\x83\x43\x16\xae\x8c\xc2\x41\xbe\xda\xba\xe4\x23\xc9\xc4\x5a\x3e\x81\x68\xe1\x4b\xc5\x32\xe8\x18\x32\xc9\xc4\x8c\x89\x20\x48\x69\x3d\x43\xf0\x6c\x4e\x2e\x97\x0d\xd5\x30\x95\x98\xb2\x77\xf0\xf9\x7c\x36\x12\xfb\xf0\x47\x80\x15\x45\x18\x96\xda\x47\x5f\xa8\xe1\x68\xdb\x18\xd4\x5c\x8a\xac\x1b\x69\x6c\x34\xae\x77\x0a\xf4\x4e\xf2\xbc\xb1\x7a\xf2\x90\xf8\x3c\x71\x4d\x2d\x07\xa4\x6a\x08\xb3\xe1\x2c\xbb\x82\x7d\xec\x2b\xc4\x74\x95\x8b\x15\xc9\x9f\x1e\x7d\xab\x18\xd2\x0d\x07\x6f\xa8\x02\x86\xea\x8e\xef\x55\xeb\x00\xda\x48\xb2\x04\xe1\x1c\x75\x00\x2b\x78\x36\xa0\xf1\xda\x80\x81\xc0\x50\xd2\xbd\xcd\xc6\x8f\xf0\x2a\x5a\x88\x55\xf4\x7f\xcb\x63\x78\x32\xeb\x62\xba\xe4\xae\x1f\x96\xd5\x34\xb4\x52\x10\x2d\x67\x55\x8c\x99\xce\xaa\xb9\x38\xe0\xfd\x45\xbb\x62\x35\xcb\xe1\xd8\xdb\x29\xcb\x61\xbd\x2f\xc1\xc6\xe1\xac\x60\xe8\x7a\x31\xf6\xc7\xa4\xa0\x6f\xd7\x86\x0a\x64\xb8\xf7\x45\x97\x36\xa0\x3d\x6a\x63\x2d\xea\x17\x57\x66\x35\xea\xe3\xd8\x1f\xb7\x60\x4c\x7b\x95\x68\x09\xd5\xac\xa2\x5f\x49\x75\x52\x13\xb0\x4b\xab\x0b\x6f\x58\x1e\x8a\x4a\xf5\x93\xee\x08\x75\x09\xb0\x82\x0c\xd8\x4b\xd5\x4d\x81\x60\xa7\x6c\x27\x0a\x70\x89\x3b\x76\xcf\xce\x45\xb7\xe2\x56\xc4\x60\x17\x07\xac\xf8\xdd\x7a\x57\xa7\x73\x7f\xd1\xcf\xd7\x09\x5a\xbe\xa4\xe8\xf0\xba\x38\x5f\x77\x7c\xa6\xaa\x76\xd8\xe6\x2a\xf9\xe3\x6e\x68\xcf\xe3\xa9\x0f\x78\x43\xe2\x4f\x87\xeb\xf7\x24\xf7\xb7\xb1\xcf\x7a\x0f\x9a\x72\x89\x3d\x17\xd8\xb1\xee\x6e\x4d\xad\xa4\xbc\x3b\xb5\xf6\xf8\xbd\x6d\xf5\x37\x2b\xc8\x7a\xde\x52\xce\xa1\xed\xda\x03\xea\xc3\x2e\xfe\xf4\x06\x7b\x04\xa5\x40\x2d\x47\xad\x58\x27\xcb\x77\x9a\xb4\xe7\xa5\x87\x9e\xd3\xaa\x36\xff\x40\xbe\x7a\xf3\x7f\xb6\xaf\x9a\x0c\xdc\x97\x9b\x2d\xf9\x99\x0e\xe9\xa6\xef\x3e\x04\xc2\x14\xc7\x2e\x75\x61\x07\xe9\x55\x44\xc8\x30\x19\x47\x84\xdc\xcc\xaf\xe6\x8b\xdb\xf9\xd1\xf1\xb8\xa5\x9d\xe7\x34\x99\x25\xeb\x88\x90\x64\x36\x9b\x8c\x93\xc1\x7a\xf2\xb0\x58\x3d\x8c\x06\xf3\xd1\x64\x1a\x11\xb2\x5c\xa4\xeb\x87\xc5\x7c\xfa\xcd\x8f\xf3\x33\x00\x00\xff\xff\xc7\x80\xb1\xed\xed\x0a\x00\x00")

func serum_instructionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_instructionGraphql,
		"serum_instruction.graphql",
	)
}

func serum_instructionGraphql() (*asset, error) {
	bytes, err := serum_instructionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_instruction.graphql", size: 2797, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x0, 0x97, 0x87, 0xf5, 0xf2, 0x68, 0x9e, 0x91, 0xaa, 0x70, 0xd6, 0x26, 0x51, 0x2a, 0x72, 0x35, 0xb3, 0x2e, 0xd7, 0x53, 0xc, 0xc, 0x18, 0xc0, 0x7f, 0xb, 0xfc, 0xd1, 0x82, 0xca, 0xbc, 0xbc}}
	return a, nil
}

var _serum_marketGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x4e\x2d\x2a\xcd\xf5\x4d\x2c\xca\x4e\x2d\x51\x50\xa8\xe6\x52\x50\x48\x4c\x49\x29\x4a\x2d\x2e\xb6\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\xc8\x4b\xcc\x4d\x45\xf0\xb9\x94\x15\x14\x14\xf4\xf5\x15\x0a\x32\x12\x8b\x53\x15\x8c\xc0\xdc\xe4\xfc\xcc\xbc\x90\xfc\xec\xd4\x3c\x05\x30\x09\x16\x2b\x48\x46\x16\xa9\x05\x04\x00\x00\xff\xff\x43\x18\x5c\x95\x6e\x00\x00\x00")

func serum_marketGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_marketGraphql,
		"serum_market.graphql",
	)
}

func serum_marketGraphql() (*asset, error) {
	bytes, err := serum_marketGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_market.graphql", size: 110, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcd, 0x9f, 0xd3, 0x7, 0x2c, 0x57, 0x3c, 0xd4, 0x62, 0x3b, 0x83, 0xb4, 0xfc, 0x4b, 0x33, 0xe0, 0x26, 0x15, 0x4a, 0x7a, 0x5c, 0x6b, 0x7e, 0x97, 0x2e, 0x5, 0x41, 0xd3, 0x3a, 0x29, 0x57, 0x44}}
	return a, nil
}

var _subscriptionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2e\x4d\x2a\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xe6\x52\x50\x50\x50\x28\x4e\x2d\x2a\xcd\xf5\xcc\x2b\x2e\x29\x2a\x4d\x06\x09\x7b\x64\x16\x97\xe4\x17\x55\x6a\x24\x26\x27\xe7\x97\xe6\x95\x58\x29\x04\x97\x14\x65\xe6\xa5\x2b\x6a\x5a\x29\x04\xa3\x29\x0d\x4a\x2d\x2e\xc8\xcf\x2b\x4e\xe5\xaa\xe5\x02\x04\x00\x00\xff\xff\xba\xf2\xbd\x88\x5e\x00\x00\x00")

func subscriptionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_subscriptionGraphql,
		"subscription.graphql",
	)
}

func subscriptionGraphql() (*asset, error) {
	bytes, err := subscriptionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "subscription.graphql", size: 94, mode: os.FileMode(0644), modTime: time.Unix(1608153851, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x54, 0x96, 0xf5, 0x98, 0x89, 0xb2, 0xd0, 0xf9, 0x8a, 0x9b, 0x7f, 0xcc, 0xb5, 0x9e, 0xee, 0xf6, 0x5f, 0xf7, 0x11, 0xe4, 0xcc, 0x82, 0x8b, 0x4a, 0x4f, 0x67, 0x9a, 0x59, 0xee, 0x8f, 0x2b}}
	return a, nil
}

var _subscription_alphaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2e\x4d\x2a\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xae\xe5\x02\x04\x00\x00\xff\xff\x4d\xe9\x40\xe8\x15\x00\x00\x00")

func subscription_alphaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_subscription_alphaGraphql,
		"subscription_alpha.graphql",
	)
}

func subscription_alphaGraphql() (*asset, error) {
	bytes, err := subscription_alphaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "subscription_alpha.graphql", size: 21, mode: os.FileMode(0644), modTime: time.Unix(1608153609, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x20, 0x13, 0xe7, 0x39, 0xf6, 0x4e, 0x44, 0x50, 0x61, 0x33, 0x62, 0x95, 0x35, 0x1d, 0x5c, 0x43, 0x43, 0xb0, 0xe4, 0xe2, 0xf9, 0xa2, 0x94, 0xdd, 0xcc, 0xad, 0x3b, 0x5d, 0x4d, 0xce, 0x16}}
	return a, nil
}

var _tokenGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x3b\xae\x02\x31\x10\x04\x73\x9f\xa2\xdf\x1d\x9e\x08\x9c\x11\x92\xf2\x39\x80\x85\x1b\x18\xb1\x1e\x5b\x9e\xd9\xc0\x20\xee\x8e\x36\x5a\x02\x08\xbb\x4b\x25\x95\x8f\x46\x1c\xeb\x9d\x8a\x67\x00\x52\xce\x9d\x66\x11\x07\xef\xa2\xd7\xbf\x00\x68\x2a\x5c\xf7\x2b\x84\x55\xd9\xd3\x5a\x55\xe3\x2f\xb5\x88\xfa\x76\xf6\x5b\xed\xe2\xe3\x13\x5c\x3a\xf9\xe0\x57\x64\x73\x6b\xd3\x88\x38\x89\xfa\xe6\x7f\x79\x32\xcf\x52\xd2\x64\x11\x3b\xf5\xa5\xe0\x1d\x00\x00\xff\xff\x26\x2a\xf1\x44\xb4\x00\x00\x00")

func tokenGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_tokenGraphql,
		"token.graphql",
	)
}

func tokenGraphql() (*asset, error) {
	bytes, err := tokenGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "token.graphql", size: 180, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2b, 0x2c, 0x9d, 0x30, 0xbf, 0x9f, 0xfb, 0x37, 0xc8, 0x67, 0xe1, 0x9e, 0xc6, 0xb, 0xc4, 0xa3, 0xd9, 0xb7, 0x8f, 0x7b, 0x43, 0xa4, 0xec, 0x19, 0x34, 0x6, 0x65, 0x6c, 0x22, 0xbb, 0x2c, 0x9}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"query.graphql": queryGraphql,

	"query_alpha.graphql": query_alphaGraphql,

	"registered_token.graphql": registered_tokenGraphql,

	"schema.graphql": schemaGraphql,

	"serum_fill.graphql": serum_fillGraphql,

	"serum_instruction.graphql": serum_instructionGraphql,

	"serum_market.graphql": serum_marketGraphql,

	"subscription.graphql": subscriptionGraphql,

	"subscription_alpha.graphql": subscription_alphaGraphql,

	"token.graphql": tokenGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"query.graphql":              &bintree{queryGraphql, map[string]*bintree{}},
	"query_alpha.graphql":        &bintree{query_alphaGraphql, map[string]*bintree{}},
	"registered_token.graphql":   &bintree{registered_tokenGraphql, map[string]*bintree{}},
	"schema.graphql":             &bintree{schemaGraphql, map[string]*bintree{}},
	"serum_fill.graphql":         &bintree{serum_fillGraphql, map[string]*bintree{}},
	"serum_instruction.graphql":  &bintree{serum_instructionGraphql, map[string]*bintree{}},
	"serum_market.graphql":       &bintree{serum_marketGraphql, map[string]*bintree{}},
	"subscription.graphql":       &bintree{subscriptionGraphql, map[string]*bintree{}},
	"subscription_alpha.graphql": &bintree{subscription_alphaGraphql, map[string]*bintree{}},
	"token.graphql":              &bintree{tokenGraphql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}