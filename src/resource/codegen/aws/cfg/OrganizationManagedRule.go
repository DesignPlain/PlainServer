package cfg

type OrganizationManagedRule struct {
	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts []string `json:"excludedAccounts,omitempty" yaml:"excludedAccounts,omitempty"`

	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters string `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" yaml:"maximumExecutionFrequency,omitempty"`

	// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
	RuleIdentifier string `json:"ruleIdentifier,omitempty" yaml:"ruleIdentifier,omitempty"`

	// Description of the rule
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the rule
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Identifier of the AWS resource to evaluate
	ResourceIdScope string `json:"resourceIdScope,omitempty" yaml:"resourceIdScope,omitempty"`

	// List of types of AWS resources to evaluate
	ResourceTypesScopes []string `json:"resourceTypesScopes,omitempty" yaml:"resourceTypesScopes,omitempty"`

	// Tag key of AWS resources to evaluate
	TagKeyScope string `json:"tagKeyScope,omitempty" yaml:"tagKeyScope,omitempty"`

	// Tag value of AWS resources to evaluate
	TagValueScope string `json:"tagValueScope,omitempty" yaml:"tagValueScope,omitempty"`
}
