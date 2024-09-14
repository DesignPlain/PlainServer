package apigateway

import types "libds/aws/types"

type UsagePlan struct {
	// Throttling limits of the usage plan.
	ThrottleSettings types.Apigateway_UsagePlanThrottleSettings `json:"throttleSettings,omitempty" yaml:"throttleSettings,omitempty"`

	// Associated API stages of the usage plan.
	ApiStages []types.Apigateway_UsagePlanApiStage `json:"apiStages,omitempty" yaml:"apiStages,omitempty"`

	// Description of a usage plan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the usage plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode string `json:"productCode,omitempty" yaml:"productCode,omitempty"`

	// Quota of the usage plan.
	QuotaSettings types.Apigateway_UsagePlanQuotaSettings `json:"quotaSettings,omitempty" yaml:"quotaSettings,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
