package amp

type RuleGroupNamespace struct {
	// the rule group namespace data that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-Ruler.html).
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// The name of the rule group namespace
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ID of the prometheus workspace the rule group namespace should be linked to
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
}
