package types

type Customerprofiles_DomainMatching struct {
	// A block that specifies the configuration about the auto-merging process. Documented below.
	AutoMerging Customerprofiles_DomainMatchingAutoMerging `json:"autoMerging,omitempty" yaml:"autoMerging,omitempty"`

	// The flag that enables the matching process of duplicate profiles.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A block that specifies the configuration for exporting Identity Resolution results. Documented below.
	ExportingConfig Customerprofiles_DomainMatchingExportingConfig `json:"exportingConfig,omitempty" yaml:"exportingConfig,omitempty"`

	// A block that specifies the day and time when you want to start the Identity Resolution Job every week. Documented below.
	JobSchedule Customerprofiles_DomainMatchingJobSchedule `json:"jobSchedule,omitempty" yaml:"jobSchedule,omitempty"`
}
