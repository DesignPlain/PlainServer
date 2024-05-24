package types

type Datasync_LocationHdfsQopConfiguration struct {
	// The data transfer protection setting configured on the HDFS cluster. This setting corresponds to your dfs.data.transfer.protection setting in the hdfs-site.xml file on your Hadoop cluster. Valid values are `DISABLED`, `AUTHENTICATION`, `INTEGRITY` and `PRIVACY`.
	DataTransferProtection string `json:"dataTransferProtection,omitempty" yaml:"dataTransferProtection,omitempty"`

	// The RPC protection setting configured on the HDFS cluster. This setting corresponds to your hadoop.rpc.protection setting in your core-site.xml file on your Hadoop cluster. Valid values are `DISABLED`, `AUTHENTICATION`, `INTEGRITY` and `PRIVACY`.
	RpcProtection string `json:"rpcProtection,omitempty" yaml:"rpcProtection,omitempty"`
}
