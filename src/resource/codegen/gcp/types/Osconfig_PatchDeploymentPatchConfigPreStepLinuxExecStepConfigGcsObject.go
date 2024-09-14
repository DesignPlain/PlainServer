package types

type Osconfig_PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject struct {
	// Name of the Cloud Storage object.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	// Bucket of the Cloud Storage object.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	GenerationNumber string `json:"generationNumber,omitempty" yaml:"generationNumber,omitempty"`
}
