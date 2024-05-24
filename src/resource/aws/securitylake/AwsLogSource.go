package securitylake

import types "DesignSphere_Server/src/resource/aws/types"

type AwsLogSource struct {
	// Specify the natively-supported AWS service to add as a source in Security Lake.
	Source types.Securitylake_AwsLogSourceSource `json:"source,omitempty" yaml:"source,omitempty"`
}
