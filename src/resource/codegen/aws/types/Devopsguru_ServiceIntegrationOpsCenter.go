package types

type Devopsguru_ServiceIntegrationOpsCenter struct {
	// Specifies if DevOps Guru is enabled to create an AWS Systems Manager OpsItem for each created insight. Valid values are `DISABLED` and `ENABLED`.
	OptInStatus string `json:"optInStatus,omitempty" yaml:"optInStatus,omitempty"`
}
