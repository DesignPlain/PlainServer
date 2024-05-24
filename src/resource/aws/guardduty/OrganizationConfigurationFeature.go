package guardduty

import types "DesignSphere_Server/src/resource/aws/types"

type OrganizationConfigurationFeature struct {
	// The additional information that will be configured for the organization See below.
	AdditionalConfigurations []types.Guardduty_OrganizationConfigurationFeatureAdditionalConfiguration `json:"additionalConfigurations,omitempty" yaml:"additionalConfigurations,omitempty"`

	// The status of the feature that is configured for the member accounts within the organization. Valid values: `NEW`, `ALL`, `NONE`.
	AutoEnable string `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// The ID of the detector that configures the delegated administrator.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// The name of the feature that will be configured for the organization. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
