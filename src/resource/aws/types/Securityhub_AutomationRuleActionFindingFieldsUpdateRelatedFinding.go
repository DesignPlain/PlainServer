package types

type Securityhub_AutomationRuleActionFindingFieldsUpdateRelatedFinding struct {
	// The product-generated identifier for a related finding.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The ARN of the product that generated a related finding.
	ProductArn string `json:"productArn,omitempty" yaml:"productArn,omitempty"`
}
