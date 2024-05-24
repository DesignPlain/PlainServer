package ec2

type NetworkInsightsAnalysis struct {
	// A list of ARNs for resources the path must traverse.
	FilterInArns []string `json:"filterInArns,omitempty" yaml:"filterInArns,omitempty"`

	/*
	   ID of the Network Insights Path to run an analysis on.

	   The following arguments are optional:
	*/
	NetworkInsightsPathId string `json:"networkInsightsPathId,omitempty" yaml:"networkInsightsPathId,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// If enabled, the resource will wait for the Network Insights Analysis status to change to `succeeded` or `failed`. Setting this to `false` will skip the process. Default: `true`.
	WaitForCompletion bool `json:"waitForCompletion,omitempty" yaml:"waitForCompletion,omitempty"`
}
