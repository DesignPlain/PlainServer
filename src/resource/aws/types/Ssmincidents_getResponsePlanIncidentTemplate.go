package types

type Ssmincidents_getResponsePlanIncidentTemplate struct {
	// The impact value of a generated incident. The following values are supported:
	Impact int `json:"impact,omitempty" yaml:"impact,omitempty"`

	// The tags assigned to an incident template. When an incident starts, Incident Manager assigns the tags specified in the template to the incident.
	IncidentTags map[string]string `json:"incidentTags,omitempty" yaml:"incidentTags,omitempty"`

	// The Amazon Simple Notification Service (Amazon SNS) targets that this incident notifies when it is updated. The `notification_target` configuration block supports the following argument:
	NotificationTargets []Ssmincidents_getResponsePlanIncidentTemplateNotificationTarget `json:"notificationTargets,omitempty" yaml:"notificationTargets,omitempty"`

	// The summary of an incident.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// The title of a generated incident.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// A string used to stop Incident Manager from creating multiple incident records for the same incident.
	DedupeString string `json:"dedupeString,omitempty" yaml:"dedupeString,omitempty"`
}
