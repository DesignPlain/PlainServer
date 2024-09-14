package types

type Datasync_LocationHdfsNameNode struct {
	// The hostname of the NameNode in the HDFS cluster. This value is the IP address or Domain Name Service (DNS) name of the NameNode. An agent that's installed on-premises uses this hostname to communicate with the NameNode in the network.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// The port that the NameNode uses to listen to client requests.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
