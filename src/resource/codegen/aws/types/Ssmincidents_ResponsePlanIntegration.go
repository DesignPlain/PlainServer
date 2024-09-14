package types

type Ssmincidents_ResponsePlanIntegration struct {
	// Details about the PagerDuty configuration for a response plan. The following values are supported:
	Pagerduties []Ssmincidents_ResponsePlanIntegrationPagerduty `json:"pagerduties,omitempty" yaml:"pagerduties,omitempty"`
}
