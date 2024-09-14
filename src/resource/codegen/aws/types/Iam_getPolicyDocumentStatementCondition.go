package types

type Iam_getPolicyDocumentStatementCondition struct {
	// Name of the [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html) to evaluate.
	Test string `json:"test,omitempty" yaml:"test,omitempty"`

	// Values to evaluate the condition against. If multiple values are provided, the condition matches if at least one of them applies. That is, AWS evaluates multiple values as though using an "OR" boolean operation.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Name of a [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys) to apply the condition to. Context variables may either be standard AWS variables starting with `aws:` or service-specific variables prefixed with the service name.
	Variable string `json:"variable,omitempty" yaml:"variable,omitempty"`
}
