module github.com/streamingfast/firehose-solana

go 1.15

require (
	cloud.google.com/go/bigtable v1.13.0
	cloud.google.com/go/storage v1.30.1
	github.com/Azure/go-autorest/autorest/adal v0.9.20 // indirect
	github.com/Azure/go-autorest/autorest/mocks v0.4.2 // indirect
	github.com/ShinyTrinkets/overseer v0.3.0
	github.com/abourget/llerrgroup v0.2.0
	github.com/dustin/go-humanize v1.0.0
	github.com/golang-jwt/jwt/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.3
	github.com/klauspost/compress v1.15.12
	github.com/lithammer/dedent v1.1.0
	github.com/lorenzosaino/go-sysctl v0.1.1
	github.com/manifoldco/promptui v0.9.0
	github.com/mholt/archiver/v3 v3.5.0
	github.com/mr-tron/base58 v1.2.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.10.1
	github.com/streamingfast/bstream v0.0.2-0.20230510131449-6b591d74130d
	github.com/streamingfast/cli v0.0.4-0.20220630165922-bc58c6666fc8
	github.com/streamingfast/dauth v0.0.0-20230726175303-fc1d7198cb33
	github.com/streamingfast/dbin v0.9.1-0.20220513054835-1abebbb944ad // indirect
	github.com/streamingfast/derr v0.0.0-20230515163924-8570aaa43fe1
	github.com/streamingfast/dgrpc v0.0.0-20230621153617-bc715cdb9fd1
	github.com/streamingfast/dlauncher v0.0.0-20220909121534-7a9aa91dbb32
	github.com/streamingfast/dmetering v0.0.0-20230706141508-cd783a0fb671
	github.com/streamingfast/dmetrics v0.0.0-20230516031116-28fcfeb4b9ed
	github.com/streamingfast/dstore v0.1.1-0.20230620124109-3924b3b36c77
	github.com/streamingfast/firehose v0.1.1-0.20230717171430-1d7a06ed55c5
	github.com/streamingfast/firehose-solana/types v0.0.0-20230126211203-a2d17ce9f8b9
	github.com/streamingfast/jsonpb v0.0.0-20210811021341-3670f0aa02d0
	github.com/streamingfast/kvdb v0.0.2-0.20210811194032-09bf862bd2e3
	github.com/streamingfast/logging v0.0.0-20220813175024-b4fbb0e893df
	github.com/streamingfast/merger v0.0.3-0.20221123202507-445dfd357868
	github.com/streamingfast/node-manager v0.0.2-0.20221115101723-d9823ffd7ad5
	github.com/streamingfast/pbgo v0.0.6-0.20221020131607-255008258d28
	github.com/streamingfast/relayer v0.0.2-0.20220909122435-e67fbc964fd9
	github.com/streamingfast/sf-tools v0.0.0-20230217210059-2ab577f61e8e
	github.com/streamingfast/shutter v1.5.0
	github.com/streamingfast/solana-go v0.5.1-0.20220502224452-432fbe84aee8
	github.com/streamingfast/substreams v1.1.11-0.20230727201625-837a63b339ef
	github.com/stretchr/testify v1.8.3
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125 // indirect
	github.com/test-go/testify v1.1.4
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.6.0
	google.golang.org/api v0.114.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

replace (
	github.com/ShinyTrinkets/overseer => github.com/dfuse-io/overseer v0.2.1-0.20191024193921-39856397cf3f
	github.com/bytecodealliance/wasmtime-go/v4 => github.com/streamingfast/wasmtime-go/v4 v4.0.0-freemem3
	github.com/graph-gophers/graphql-go => github.com/streamingfast/graphql-go v0.0.0-20210204202750-0e485a040a3c
)
