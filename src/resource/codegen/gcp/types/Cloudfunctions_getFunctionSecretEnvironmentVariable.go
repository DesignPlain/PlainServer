package types

type Cloudfunctions_getFunctionSecretEnvironmentVariable struct {
	// Name of the environment variable.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret string `json:"secret,omitempty" yaml:"secret,omitempty"`

	// Version of the secret (version number or the string "latest"). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new clones start.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
