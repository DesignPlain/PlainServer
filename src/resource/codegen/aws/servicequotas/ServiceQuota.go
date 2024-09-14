package servicequotas

type ServiceQuota struct {
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value float64 `json:"value,omitempty" yaml:"value,omitempty"`

	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode string `json:"quotaCode,omitempty" yaml:"quotaCode,omitempty"`

	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode string `json:"serviceCode,omitempty" yaml:"serviceCode,omitempty"`
}
