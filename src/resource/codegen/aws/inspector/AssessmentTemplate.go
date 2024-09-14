package inspector

import types "libds/aws/types"

type AssessmentTemplate struct {
	// The duration of the inspector run.
	Duration int `json:"duration,omitempty" yaml:"duration,omitempty"`

	// A block that enables sending notifications about a specified assessment template event to a designated SNS topic. See Event Subscriptions for details.
	EventSubscriptions []types.Inspector_AssessmentTemplateEventSubscription `json:"eventSubscriptions,omitempty" yaml:"eventSubscriptions,omitempty"`

	// The name of the assessment template.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The rules to be used during the run.
	RulesPackageArns []string `json:"rulesPackageArns,omitempty" yaml:"rulesPackageArns,omitempty"`

	// Key-value map of tags for the Inspector assessment template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The assessment target ARN to attach the template to.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`
}
