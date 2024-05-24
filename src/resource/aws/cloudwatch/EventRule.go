package cloudwatch

type EventRule struct {
	// The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. --Note--: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
	EventPattern string `json:"eventPattern,omitempty" yaml:"eventPattern,omitempty"`

	// The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. --Note--: Due to the length of the generated suffix, must be 38 characters or less.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The scheduling expression. For example, `cron(0 20 - - ? -)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`

	// The description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name or ARN of the event bus to associate with this rule.
	   If you omit this, the `default` event bus is used.
	*/
	EventBusName string `json:"eventBusName,omitempty" yaml:"eventBusName,omitempty"`

	/*
	   Whether the rule should be enabled.
	   Defaults to `true`.
	   Conflicts with `state`.
	*/
	IsEnabled bool `json:"isEnabled,omitempty" yaml:"isEnabled,omitempty"`

	/*
	   State of the rule.
	   Valid values are `DISABLED`, `ENABLED`, and `ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS`.
	   When state is `ENABLED`, the rule is enabled for all events except those delivered by CloudTrail.
	   To also enable the rule for events delivered by CloudTrail, set `state` to `ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS`.
	   Defaults to `ENABLED`.
	   Conflicts with `is_enabled`.

	   --NOTE:-- The rule state  `ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS` cannot be used in conjunction with the `schedule_expression` argument.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
