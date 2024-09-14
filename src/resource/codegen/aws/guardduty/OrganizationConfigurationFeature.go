package guardduty

import types "libds/aws/types"

type OrganizationConfigurationFeature struct {
	// The name of the feature that will be configured for the organization. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`. Only one of two features `EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING` can be added, adding both features will cause an error. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html) for the current list of supported values.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Additional feature configuration block for features `EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING`. See below.
	AdditionalConfigurations []types.Guardduty_OrganizationConfigurationFeatureAdditionalConfiguration `json:"additionalConfigurations,omitempty" yaml:"additionalConfigurations,omitempty"`

	// The status of the feature that is configured for the member accounts within the organization. Valid values: `NEW`, `ALL`, `NONE`.
	AutoEnable string `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// The ID of the detector that configures the delegated administrator.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`
}
