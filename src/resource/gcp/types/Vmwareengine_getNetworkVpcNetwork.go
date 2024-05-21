package types

type Vmwareengine_getNetworkVpcNetwork struct {
	/*
	   The relative resource name of the service VPC network this VMware Engine network is attached to.
	   For example: projects/123123/global/networks/my-network
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// Type of VPC network (INTRANET, INTERNET, or GOOGLE_CLOUD)
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
