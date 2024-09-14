package monitoring

import types "libds/gcp/types"

type AlertPolicy struct {
	/*
	   Identifies the notification channels to which notifications should be
	   sent when incidents are opened or closed or when new violations occur
	   on an already opened incident. Each element of this array corresponds
	   to the name field in each of the NotificationChannel objects that are
	   returned from the notificationChannels.list method. The syntax of the
	   entries in this field is
	   `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
	*/
	NotificationChannels []string `json:"notificationChannels,omitempty" yaml:"notificationChannels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The severity of an alert policy indicates how important incidents generated
	   by that policy are. The severity level will be displayed on the Incident
	   detail page and in notifications.
	   Possible values are: `CRITICAL`, `ERROR`, `WARNING`.
	*/
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`

	/*
	   A list of conditions for the policy. The conditions are combined by
	   AND or OR according to the combiner field. If the combined conditions
	   evaluate to true, then an incident is created. A policy can have from
	   one to six conditions.
	   Structure is documented below.
	*/
	Conditions []types.Monitoring_AlertPolicyCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	/*
	   A short name or phrase used to identify the policy in
	   dashboards, notifications, and incidents. To avoid confusion, don't use
	   the same display name for multiple policies in the same project. The
	   name is limited to 512 Unicode characters.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Documentation that is included with notifications and incidents related
	   to this policy. Best practice is for the documentation to include information
	   to help responders understand, mitigate, escalate, and correct the underlying
	   problems detected by the alerting policy. Notification channels that have
	   limited capacity might not show this documentation.
	   Structure is documented below.
	*/
	Documentation types.Monitoring_AlertPolicyDocumentation `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Whether or not the policy is enabled. The default is true.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   This field is intended to be used for organizing and identifying the AlertPolicy
	   objects.The field can contain up to 64 entries. Each key and value is limited
	   to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	   can contain only lowercase letters, numerals, underscores, and dashes. Keys
	   must begin with a letter.
	*/
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`

	/*
	   Control over how this alert policy's notification channels are notified.
	   Structure is documented below.
	*/
	AlertStrategy types.Monitoring_AlertPolicyAlertStrategy `json:"alertStrategy,omitempty" yaml:"alertStrategy,omitempty"`

	/*
	   How to combine the results of multiple conditions to
	   determine if an incident should be opened.
	   Possible values are: `AND`, `OR`, `AND_WITH_MATCHING_RESOURCE`.
	*/
	Combiner string `json:"combiner,omitempty" yaml:"combiner,omitempty"`
}
