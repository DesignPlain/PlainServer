package directconnect

type HostedConnection struct {
	// The dedicated VLAN provisioned to the hosted connection.
	Vlan int `json:"vlan,omitempty" yaml:"vlan,omitempty"`

	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	Bandwidth string `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`

	// The ID of the interconnect or LAG.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// The name of the connection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ID of the AWS account of the customer for the connection.
	OwnerAccountId string `json:"ownerAccountId,omitempty" yaml:"ownerAccountId,omitempty"`
}
