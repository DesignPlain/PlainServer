package guardduty

import types "DesignSphere_Server/src/resource/aws/types"

type DetectorFeature struct {
	// The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Additional feature configuration block. See below.
	AdditionalConfigurations []types.Guardduty_DetectorFeatureAdditionalConfiguration `json:"additionalConfigurations,omitempty" yaml:"additionalConfigurations,omitempty"`

	// Amazon GuardDuty detector ID.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
