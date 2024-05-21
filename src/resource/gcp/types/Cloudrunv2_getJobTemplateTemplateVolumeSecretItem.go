package types

type Cloudrunv2_getJobTemplateTemplateVolumeSecretItem struct {
	// Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal). If 0 or not set, the Volume's default mode will be used.
	Mode int `json:"mode,omitempty" yaml:"mode,omitempty"`

	// The relative path of the secret in the container.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
