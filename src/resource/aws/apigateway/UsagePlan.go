package apigateway

import types "DesignSphere_Server/src/resource/aws/types"

type UsagePlan struct {
	// Name of the usage plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode string `json:"productCode,omitempty" yaml:"productCode,omitempty"`

	// The quota settings of the usage plan.
	QuotaSettings types.Apigateway_UsagePlanQuotaSettings `json:"quotaSettings,omitempty" yaml:"quotaSettings,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The throttling limits of the usage plan.
	ThrottleSettings types.Apigateway_UsagePlanThrottleSettings `json:"throttleSettings,omitempty" yaml:"throttleSettings,omitempty"`

	// Associated API stages of the usage plan.
	ApiStages []types.Apigateway_UsagePlanApiStage `json:"apiStages,omitempty" yaml:"apiStages,omitempty"`

	// Description of a usage plan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
