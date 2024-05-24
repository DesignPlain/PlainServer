package cfg

type OrganizationCustomRule struct {
	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" yaml:"maximumExecutionFrequency,omitempty"`

	// The name of the rule
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of types of AWS resources to evaluate
	ResourceTypesScopes []string `json:"resourceTypesScopes,omitempty" yaml:"resourceTypesScopes,omitempty"`

	// Tag key of AWS resources to evaluate
	TagKeyScope string `json:"tagKeyScope,omitempty" yaml:"tagKeyScope,omitempty"`

	// Description of the rule
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts []string `json:"excludedAccounts,omitempty" yaml:"excludedAccounts,omitempty"`

	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters string `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// Amazon Resource Name (ARN) of the rule Lambda Function
	LambdaFunctionArn string `json:"lambdaFunctionArn,omitempty" yaml:"lambdaFunctionArn,omitempty"`

	// Identifier of the AWS resource to evaluate
	ResourceIdScope string `json:"resourceIdScope,omitempty" yaml:"resourceIdScope,omitempty"`

	// Tag value of AWS resources to evaluate
	TagValueScope string `json:"tagValueScope,omitempty" yaml:"tagValueScope,omitempty"`

	// List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
	TriggerTypes []string `json:"triggerTypes,omitempty" yaml:"triggerTypes,omitempty"`
}
