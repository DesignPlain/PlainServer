package types

type Compute_NetworkAttachmentConnectionEndpoint struct {
	/*
	   (Output)
	   The IPv4 address assigned to the producer instance network interface. This value will be a range in case of Serverless.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   (Output)
	   The project id or number of the interface to which the IP was assigned.
	*/
	ProjectIdOrNum string `json:"projectIdOrNum,omitempty" yaml:"projectIdOrNum,omitempty"`

	/*
	   (Output)
	   Alias IP ranges from the same subnetwork.
	*/
	SecondaryIpCidrRanges string `json:"secondaryIpCidrRanges,omitempty" yaml:"secondaryIpCidrRanges,omitempty"`

	/*
	   (Output)
	   The status of a connected endpoint to this network attachment.
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	/*
	   (Output)
	   The subnetwork used to assign the IP to the producer instance network interface.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`
}
