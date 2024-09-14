package types

type Ssmincidents_getResponsePlanIntegration struct {
	// Details about the PagerDuty configuration for a response plan. The following values are supported:
	Pagerduties []Ssmincidents_getResponsePlanIntegrationPagerduty `json:"pagerduties,omitempty" yaml:"pagerduties,omitempty"`
}
