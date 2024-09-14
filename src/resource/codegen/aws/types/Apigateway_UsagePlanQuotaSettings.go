package types

type Apigateway_UsagePlanQuotaSettings struct {
	// Maximum number of requests that can be made in a given time period.
	Limit int `json:"limit,omitempty" yaml:"limit,omitempty"`

	// Number of requests subtracted from the given limit in the initial time period.
	Offset int `json:"offset,omitempty" yaml:"offset,omitempty"`

	// Time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
	Period string `json:"period,omitempty" yaml:"period,omitempty"`
}
