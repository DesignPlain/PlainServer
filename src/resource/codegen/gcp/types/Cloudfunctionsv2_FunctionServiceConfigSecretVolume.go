package types

type Cloudfunctionsv2_FunctionServiceConfigSecretVolume struct {
	// The path within the container to mount the secret volume. For example, setting the mountPath as /etc/secrets would mount the secret value files under the /etc/secrets directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount path: /etc/secrets
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	Secret string `json:"secret,omitempty" yaml:"secret,omitempty"`

	/*
	   List of secret versions to mount for this secret. If empty, the latest version of the secret will be made available in a file named after the secret under the mount point.'
	   Structure is documented below.
	*/
	Versions []Cloudfunctionsv2_FunctionServiceConfigSecretVolumeVersion `json:"versions,omitempty" yaml:"versions,omitempty"`
}
