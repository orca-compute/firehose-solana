package codec

import "github.com/streamingfast/logging"

var zlog, _ = logging.PackageLogger("sfsol", "github.com/streamingfast/sf-solana/node-mananager/codec")

func init() {
	logging.InstantiateLoggers()
}

type ObjectReader func() (interface{}, error)