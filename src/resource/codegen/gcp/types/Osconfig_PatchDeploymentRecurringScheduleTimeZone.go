package types

type Osconfig_PatchDeploymentRecurringScheduleTimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// IANA Time Zone Database version number, e.g. "2019a".
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
