package types

type Ssm_PatchBaselineSource struct {
	// Value of the yum repo configuration. For information about other options available for your yum repository configuration, see the [`dnf.conf` documentation](https://man7.org/linux/man-pages/man5/dnf.conf.5.html)
	Configuration string `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Name specified to identify the patch source.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specific operating system versions a patch repository applies to, such as `"Ubuntu16.04"`, `"AmazonLinux2016.09"`, `"RedhatEnterpriseLinux7.2"` or `"Suse12.7"`. For lists of supported product values, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html).
	Products []string `json:"products,omitempty" yaml:"products,omitempty"`
}
