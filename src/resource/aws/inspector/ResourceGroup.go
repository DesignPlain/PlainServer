package inspector

type ResourceGroup struct {
	// Key-value map of tags that are used to select the EC2 instances to be included in an Amazon Inspector assessment target.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
