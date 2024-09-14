package servicequotas

type Template struct {
	// Quota identifier. To find the quota code for a specific quota, use the aws.servicequotas.ServiceQuota data source.
	QuotaCode string `json:"quotaCode,omitempty" yaml:"quotaCode,omitempty"`

	// AWS Region to which the template applies.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Service identifier. To find the service code value for an AWS service, use the aws.servicequotas.getService data source.
	ServiceCode string `json:"serviceCode,omitempty" yaml:"serviceCode,omitempty"`

	// The new, increased value for the quota.
	Value float64 `json:"value,omitempty" yaml:"value,omitempty"`
}
