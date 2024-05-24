package cfg

type OrganizationCustomPolicyRule struct {
	// Description of the rule
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters string `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// runtime system for your organization AWS Config Custom Policy rules
	PolicyRuntime string `json:"policyRuntime,omitempty" yaml:"policyRuntime,omitempty"`

	// policy definition containing the logic for your organization AWS Config Custom Policy rule
	PolicyText string `json:"policyText,omitempty" yaml:"policyText,omitempty"`

	// Tag key of AWS resources to evaluate
	TagKeyScope string `json:"tagKeyScope,omitempty" yaml:"tagKeyScope,omitempty"`

	/*
	   List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`

	   The following arguments are optional:
	*/
	TriggerTypes []string `json:"triggerTypes,omitempty" yaml:"triggerTypes,omitempty"`

	// List of AWS account identifiers to exclude from the rule
	DebugLogDeliveryAccounts []string `json:"debugLogDeliveryAccounts,omitempty" yaml:"debugLogDeliveryAccounts,omitempty"`

	// Maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" yaml:"maximumExecutionFrequency,omitempty"`

	// name of the rule
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Identifier of the AWS resource to evaluate
	ResourceIdScope string `json:"resourceIdScope,omitempty" yaml:"resourceIdScope,omitempty"`

	// List of types of AWS resources to evaluate
	ResourceTypesScopes []string `json:"resourceTypesScopes,omitempty" yaml:"resourceTypesScopes,omitempty"`

	// Tag value of AWS resources to evaluate
	TagValueScope string `json:"tagValueScope,omitempty" yaml:"tagValueScope,omitempty"`

	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts []string `json:"excludedAccounts,omitempty" yaml:"excludedAccounts,omitempty"`
}
