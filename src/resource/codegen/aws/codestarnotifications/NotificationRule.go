package codestarnotifications

import types "libds/aws/types"

type NotificationRule struct {
	// The ARN of the resource to associate with the notification rule.
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`

	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Targets []types.Codestarnotifications_NotificationRuleTarget `json:"targets,omitempty" yaml:"targets,omitempty"`

	// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
	DetailType string `json:"detailType,omitempty" yaml:"detailType,omitempty"`

	/*
	   A list of event types associated with this notification rule.
	   For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
	*/
	EventTypeIds []string `json:"eventTypeIds,omitempty" yaml:"eventTypeIds,omitempty"`

	// The name of notification rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
