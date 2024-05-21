package types

type Clouddeploy_AutomationSelectorTarget struct {
	// ID of the `Target`. The value of this field could be one of the following: - The last segment of a target name. It only needs the ID to determine which target is being referred to - "-", all targets in a location.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Target labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
