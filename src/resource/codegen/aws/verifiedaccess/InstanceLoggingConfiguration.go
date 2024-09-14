package verifiedaccess

import types "libds/aws/types"

type InstanceLoggingConfiguration struct {
	// A block that specifies the configuration options for Verified Access instances. Detailed below.
	AccessLogs types.Verifiedaccess_InstanceLoggingConfigurationAccessLogs `json:"accessLogs,omitempty" yaml:"accessLogs,omitempty"`

	// The ID of the Verified Access instance.
	VerifiedaccessInstanceId string `json:"verifiedaccessInstanceId,omitempty" yaml:"verifiedaccessInstanceId,omitempty"`
}
