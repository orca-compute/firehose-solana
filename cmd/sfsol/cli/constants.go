package cli

const (
	//Protocol               pbbstream.Protocol = pbbstream.Protocol_EOS //todo
	TrxDBDSN               string = "badger://{sf-data-dir}/storage/trxdb"     //%s will be replaced by `<data-dir>`
	FluxDSN                string = "badger://{sf-data-dir}/storage/statedb"   //%s will be replaced by `<data-dir>/<flux-data-dir>
	SerumHistDSN           string = "badger://{sf-data-dir}/storage/serumhist" //%s will be replaced by `<data-dir>/<flux-data-dir>
	MergedBlocksStoreURL   string = "file://{sf-data-dir}/storage/merged-blocks"
	FilteredBlocksStoreURL string = "file://{sf-data-dir}/storage/filtered-merged-blocks"
	IndicesStoreURL        string = "file://{sf-data-dir}/storage/indexes"
	OneBlockStoreURL       string = "file://{sf-data-dir}/storage/one-blocks"
	BlockDataStoreURL      string = "file://{sf-data-dir}/storage/data-blocks"
	SnapshotsURL           string = "file://{sf-data-dir}/storage/snapshots"
	BlocksCacheDirectory   string = "{sf-data-dir}/blocks-cache"
	DmeshDSN               string = "local://"
	DmeshServiceVersion    string = "v1"
	NetworkID              string = "development"
	SFNetworkID            string = "sol-local"

	APIProxyHTTPListenAddr  string = ":8080"
	DashboardHTTPListenAddr string = ":8081"
	MetricsListenAddr       string = ":9102"

	MinerNodeHTTPServingAddr    string = ":14001"
	ReaderNodeHTTPServingAddr   string = ":14002"
	ReaderNodeGRPCAddr          string = ":14003"
	PeeringNodeHTTPServingAddr  string = ":14004"
	BackupNodeHTTPServingAddr   string = ":14005"
	RelayerServingAddr          string = ":14006"
	MergerServingAddr           string = ":14007"
	AbiServingAddr              string = ":14008"
	BlockmetaServingAddr        string = ":14009"
	ArchiveServingAddr          string = ":14010"
	ArchiveHTTPServingAddr      string = ":14011"
	LiveServingAddr             string = ":14012"
	RouterServingAddr           string = ":14013"
	RouterHTTPServingAddr       string = ":14014"
	TrxDBHTTPServingAddr        string = ":14015"
	IndexerServingAddr          string = ":14016"
	IndexerHTTPServingAddr      string = ":14017"
	DgraphqlHTTPServingAddr     string = ":14018"
	DgraphqlGRPCServingAddr     string = ":14019"
	ForkResolverServingAddr     string = ":14020"
	ForkResolverHTTPServingAddr string = ":14021"
	SolqHTTPServingAddr         string = ":14022"
	DashboardGRPCServingAddr    string = ":14023"
	FilteringRelayerServingAddr string = ":14024"
	TokenmetaGRPCServingAddr    string = ":14025"
	FirehoseGRPCServingAddr     string = ":14026"
	SerumHistoryGRPCServingAddr string = ":14027"
	SerumHistoryHTTPServingAddr string = ":14028"

	// Solana node instance port definitions
	MinerNodeRPCPort      string = "14100"
	MinerNodeRPCWSPort    string = "14101"
	MinerNodeGossipPort   string = "14110"
	MinerNodeP2PPortStart string = "14111"
	MinerNodeP2PPortEnd   string = "14199"

	ReaderNodeRPCPort      string = "14200"
	ReaderNodeRPCWSPort    string = "14201"
	ReaderNodeGossipPort   string = "14210"
	ReaderNodeP2PPortStart string = "14211"
	ReaderNodeP2PPortEnd   string = "14299"

	PeeringNodeRPCPort      string = "14300"
	PeeringNodeRPCWSPort    string = "14301"
	PeeringNodeGossipPort   string = "14310"
	PeeringNodeP2PPortStart string = "14311"
	PeeringNodeP2PPortEnd   string = "14399"

	BackupNodeRPCPort      string = "14400"
	BackupNodeRPCWSPort    string = "14401"
	BackupNodeGossipPort   string = "14410"
	BackupNodeP2PPortStart string = "14411"
	BackupNodeP2PPortEnd   string = "14499"
)
