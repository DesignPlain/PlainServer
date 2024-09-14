package guardduty

import types "libds/aws/types"

type DetectorFeature struct {
	// Additional feature configuration block for features`EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING`. See below.
	AdditionalConfigurations []types.Guardduty_DetectorFeatureAdditionalConfiguration `json:"additionalConfigurations,omitempty" yaml:"additionalConfigurations,omitempty"`

	// Amazon GuardDuty detector ID.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`. Only one of two features `EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING` can be added, adding both features will cause an error. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html) for the current list of supported values.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
