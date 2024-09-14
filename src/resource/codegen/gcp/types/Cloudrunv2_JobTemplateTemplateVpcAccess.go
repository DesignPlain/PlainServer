package types

type Cloudrunv2_JobTemplateTemplateVpcAccess struct {
	/*
	   Direct VPC egress settings. Currently only single network interface is supported.
	   Structure is documented below.
	*/
	NetworkInterfaces []Cloudrunv2_JobTemplateTemplateVpcAccessNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	// VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.
	Connector string `json:"connector,omitempty" yaml:"connector,omitempty"`

	/*
	   Traffic VPC egress settings.
	   Possible values are: `ALL_TRAFFIC`, `PRIVATE_RANGES_ONLY`.
	*/
	Egress string `json:"egress,omitempty" yaml:"egress,omitempty"`
}
