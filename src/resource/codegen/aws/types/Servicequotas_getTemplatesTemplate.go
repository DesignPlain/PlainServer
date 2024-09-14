package types

type Servicequotas_getTemplatesTemplate struct {
	// Unit of measurement.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// (Required) The new, increased value for the quota.
	Value float64 `json:"value,omitempty" yaml:"value,omitempty"`

	// Indicates whether the quota is global.
	GlobalQuota bool `json:"globalQuota,omitempty" yaml:"globalQuota,omitempty"`

	// Quota identifier.
	QuotaCode string `json:"quotaCode,omitempty" yaml:"quotaCode,omitempty"`

	// Quota name.
	QuotaName string `json:"quotaName,omitempty" yaml:"quotaName,omitempty"`

	// AWS Region to which the quota increases apply.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// (Required) Service identifier.
	ServiceCode string `json:"serviceCode,omitempty" yaml:"serviceCode,omitempty"`

	// Service name.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}
