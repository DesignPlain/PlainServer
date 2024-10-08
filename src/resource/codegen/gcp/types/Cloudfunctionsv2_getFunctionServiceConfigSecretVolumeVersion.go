package types

type Cloudfunctionsv2_getFunctionServiceConfigSecretVolumeVersion struct {
	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mountPath as '/etc/secrets' and path as secret_foo would mount the secret value file at /etc/secrets/secret_foo.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
