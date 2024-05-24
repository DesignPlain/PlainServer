package types

type Ssmincidents_ResponsePlanIntegrationPagerduty struct {
	// The name of the response plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the AWS Secrets Manager secret that stores your PagerDuty key &mdash; either a General Access REST API Key or User Token REST API Key &mdash; and other user credentials.

	   For more information about the constraints for each field, see [CreateResponsePlan](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateResponsePlan.html) in the -AWS Systems Manager Incident Manager API Reference-.
	*/
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`

	// The ID of the PagerDuty service that the response plan associated with the incident at launch.
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
}
