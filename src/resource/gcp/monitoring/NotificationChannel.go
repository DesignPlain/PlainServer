package monitoring

import types "DesignSphere_Server/src/resource/gcp/types"

type NotificationChannel struct {
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   If true, the notification channel will be deleted regardless
	   of its use in alert policies (the policies will be updated
	   to remove the channel). If false, channels that are still
	   referenced by an existing alerting policy will fail to be
	   deleted in a delete operation.
	*/
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	/*
	   Configuration fields that define the channel and its behavior. The
	   permissible and required labels are specified in the
	   NotificationChannelDescriptor corresponding to the type field.
	   Labels with sensitive data are obfuscated by the API and therefore the provider cannot
	   determine if there are upstream changes to these fields. They can also be configured via
	   the sensitive_labels block, but cannot be configured in both places.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Different notification type behaviors are configured primarily using the the `labels` field on this
	   resource. This block contains the labels which contain secrets or passwords so that they can be marked
	   sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	   in the `labels` map in the api request.
	   Credentials may not be specified in both locations and will cause an error. Changing from one location
	   to a different credential configuration in the config will require an apply to update state.
	   Structure is documented below.
	*/
	SensitiveLabels types.Monitoring_NotificationChannelSensitiveLabels `json:"sensitiveLabels,omitempty" yaml:"sensitiveLabels,omitempty"`

	/*
	   The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...


	   - - -
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`
}
