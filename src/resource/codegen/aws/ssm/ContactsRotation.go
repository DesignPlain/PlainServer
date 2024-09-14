package ssm

import types "libds/aws/types"

type ContactsRotation struct {
	// The time zone to base the rotation’s activity on in Internet Assigned Numbers Authority (IANA) format.
	TimeZoneId string `json:"timeZoneId,omitempty" yaml:"timeZoneId,omitempty"`

	// Amazon Resource Names (ARNs) of the contacts to add to the rotation. The order in which you list the contacts is their shift order in the rotation schedule.
	ContactIds []string `json:"contactIds,omitempty" yaml:"contactIds,omitempty"`

	// The name for the rotation.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Information about when an on-call rotation is in effect and how long the rotation period lasts. Exactly one of either `daily_settings`, `monthly_settings`, or `weekly_settings` must be populated. See Recurrence for more details.

	   The following arguments are optional:
	*/
	Recurrence types.Ssm_ContactsRotationRecurrence `json:"recurrence,omitempty" yaml:"recurrence,omitempty"`

	// The date and time, in RFC 3339 format, that the rotation goes into effect.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
