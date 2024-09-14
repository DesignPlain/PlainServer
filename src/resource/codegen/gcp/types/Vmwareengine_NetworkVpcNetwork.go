package types

type Vmwareengine_NetworkVpcNetwork struct {
	/*
	   (Output)
	   The relative resource name of the service VPC network this VMware Engine network is attached to.
	   For example: projects/123123/global/networks/my-network
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   VMware Engine network type.
	   Possible values are: `LEGACY`, `STANDARD`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
