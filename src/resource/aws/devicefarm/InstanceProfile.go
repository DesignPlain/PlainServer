package devicefarm

type InstanceProfile struct {
	// When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
	RebootAfterUse bool `json:"rebootAfterUse,omitempty" yaml:"rebootAfterUse,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the instance profile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
	ExcludeAppPackagesFromCleanups []string `json:"excludeAppPackagesFromCleanups,omitempty" yaml:"excludeAppPackagesFromCleanups,omitempty"`

	// The name for the instance profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
	PackageCleanup bool `json:"packageCleanup,omitempty" yaml:"packageCleanup,omitempty"`
}
