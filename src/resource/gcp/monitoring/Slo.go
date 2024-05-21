package monitoring

import types "DesignSphere_Server/src/resource/gcp/types"

type Slo struct {
	/*
	   ID of the service to which this SLO belongs.


	   - - -
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId string `json:"sloId,omitempty" yaml:"sloId,omitempty"`

	/*
	   Basic Service-Level Indicator (SLI) on a well-known service type.
	   Performance will be computed on the basis of pre-defined metrics.
	   SLIs are used to measure and calculate the quality of the Service's
	   performance with respect to a single aspect of service quality.
	   Exactly one of the following must be set:
	   `basic_sli`, `request_based_sli`, `windows_based_sli`
	   Structure is documented below.
	*/
	BasicSli types.Monitoring_SloBasicSli `json:"basicSli,omitempty" yaml:"basicSli,omitempty"`

	/*
	   The fraction of service that must be good in order for this objective
	   to be met. 0 < goal <= 0.999
	*/
	Goal float64 `json:"goal,omitempty" yaml:"goal,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A rolling time period, semantically "in the past X days".
	   Must be between 1 to 30 days, inclusive.
	*/
	RollingPeriodDays int `json:"rollingPeriodDays,omitempty" yaml:"rollingPeriodDays,omitempty"`

	/*
	   This field is intended to be used for organizing and identifying the AlertPolicy
	   objects.The field can contain up to 64 entries. Each key and value is limited
	   to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	   can contain only lowercase letters, numerals, underscores, and dashes. Keys
	   must begin with a letter.
	*/
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`

	/*
	   A windows-based SLI defines the criteria for time windows.
	   good_service is defined based off the count of these time windows
	   for which the provided service was of good quality.
	   A SLI describes a good service. It is used to measure and calculate
	   the quality of the Service's performance with respect to a single
	   aspect of service quality.
	   Exactly one of the following must be set:
	   `basic_sli`, `request_based_sli`, `windows_based_sli`
	   Structure is documented below.
	*/
	WindowsBasedSli types.Monitoring_SloWindowsBasedSli `json:"windowsBasedSli,omitempty" yaml:"windowsBasedSli,omitempty"`

	/*
	   A calendar period, semantically "since the start of the current
	   <calendarPeriod>".
	   Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
	*/
	CalendarPeriod string `json:"calendarPeriod,omitempty" yaml:"calendarPeriod,omitempty"`

	// Name used for UI elements listing this SLO.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   A request-based SLI defines a SLI for which atomic units of
	   service are counted directly.
	   A SLI describes a good service.
	   It is used to measure and calculate the quality of the Service's
	   performance with respect to a single aspect of service quality.
	   Exactly one of the following must be set:
	   `basic_sli`, `request_based_sli`, `windows_based_sli`
	   Structure is documented below.
	*/
	RequestBasedSli types.Monitoring_SloRequestBasedSli `json:"requestBasedSli,omitempty" yaml:"requestBasedSli,omitempty"`
}
