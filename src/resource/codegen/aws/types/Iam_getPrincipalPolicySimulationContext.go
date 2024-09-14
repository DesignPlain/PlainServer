package types

type Iam_getPrincipalPolicySimulationContext struct {
	/*
	   The context _condition key_ to set.

	   If you have policies containing `Condition` elements or using dynamic interpolations then you will need to provide suitable values for each condition key your policies use. See [Actions, resources, and condition keys for AWS services](https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html) to find the various condition keys that are normally provided for real requests to each action of each AWS service.
	*/
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   An IAM value type that determines how the policy simulator will interpret the strings given in `values`.

	   For more information, see the `ContextKeyType` field of [`iam.ContextEntry`](https://docs.aws.amazon.com/IAM/latest/APIReference/API_ContextEntry.html) in the underlying API.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A set of one or more values for this context entry.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
