package types

type Billing_BudgetAllUpdatesRule struct {
	/*
	   Boolean. When set to true, disables default notifications sent
	   when a threshold is exceeded. Default recipients are
	   those with Billing Account Administrators and Billing
	   Account Users IAM roles for the target account.
	*/
	DisableDefaultIamRecipients bool `json:"disableDefaultIamRecipients,omitempty" yaml:"disableDefaultIamRecipients,omitempty"`

	/*
	   The full resource name of a monitoring notification
	   channel in the form
	   projects/{project_id}/notificationChannels/{channel_id}.
	   A maximum of 5 channels are allowed.
	*/
	MonitoringNotificationChannels []string `json:"monitoringNotificationChannels,omitempty" yaml:"monitoringNotificationChannels,omitempty"`

	/*
	   The name of the Cloud Pub/Sub topic where budget related
	   messages will be published, in the form
	   projects/{project_id}/topics/{topic_id}. Updates are sent
	   at regular intervals to the topic.
	*/
	PubsubTopic string `json:"pubsubTopic,omitempty" yaml:"pubsubTopic,omitempty"`

	/*
	   The schema version of the notification. Only "1.0" is
	   accepted. It represents the JSON schema as defined in
	   https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
	*/
	SchemaVersion string `json:"schemaVersion,omitempty" yaml:"schemaVersion,omitempty"`
}
