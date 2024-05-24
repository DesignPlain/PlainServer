package signer

import types "DesignSphere_Server/src/resource/aws/types"

type SigningJob struct {
	// The S3 bucket in which to save your signed object. See Destination below for details.
	Destination types.Signer_SigningJobDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
	IgnoreSigningJobFailure bool `json:"ignoreSigningJobFailure,omitempty" yaml:"ignoreSigningJobFailure,omitempty"`

	// The name of the profile to initiate the signing operation.
	ProfileName string `json:"profileName,omitempty" yaml:"profileName,omitempty"`

	// The S3 bucket that contains the object to sign. See Source below for details.
	Source types.Signer_SigningJobSource `json:"source,omitempty" yaml:"source,omitempty"`
}
