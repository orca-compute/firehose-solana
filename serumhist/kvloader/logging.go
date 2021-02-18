package kvloader

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var zlog *zap.Logger

var traceEnabled = logging.IsTraceEnabled("serumhist", "github.com/dfuse-io/dfuse-solana/serumhist/kvdbloader")
var logEveryXSlot = uint64(10)

func init() {
	logging.Register("github.com/dfuse-io/dfuse-solana/serumhist/kvdbloader", &zlog)
}