package types

type Compute_SecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader struct {
	/*
	   A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation.
	   The field value must be given if the field `operator` is not `EQUALS_ANY`, and cannot be given if the field `operator` is `EQUALS_ANY`.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// You can specify an exact match or a partial match by using a field operator and a field value.
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`
}
