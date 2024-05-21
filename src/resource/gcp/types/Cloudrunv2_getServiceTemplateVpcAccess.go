package types



type Cloudrunv2_getServiceTemplateVpcAccess struct {
	// VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.
	Connector string `json:"connector,omitempty" yaml:"connector,omitempty"`

	// Traffic VPC egress settings. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"]
	Egress string `json:"egress,omitempty" yaml:"egress,omitempty"`

	// Direct VPC egress settings. Currently only single network interface is supported.
	NetworkInterfaces []Cloudrunv2_getServiceTemplateVpcAccessNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`
}
