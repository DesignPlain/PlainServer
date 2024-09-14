package types

type Ssmincidents_getResponsePlanIntegrationPagerduty struct {
	// The name of the PagerDuty configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ID of the AWS Secrets Manager secret that stores your PagerDuty key &mdash; either a General Access REST API Key or User Token REST API Key &mdash; and other user credentials.
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`

	// The ID of the PagerDuty service that the response plan associates with an incident when it launches.
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
}
