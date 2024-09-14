package iam

type VirtualMfaDevice struct {
	// The path for the virtual MFA device.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Map of resource tags for the virtual mfa device. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
	VirtualMfaDeviceName string `json:"virtualMfaDeviceName,omitempty" yaml:"virtualMfaDeviceName,omitempty"`
}
