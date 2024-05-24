package imagebuilder

import types "DesignSphere_Server/src/resource/aws/types"

type DistributionConfiguration struct {
	// Description of the distribution configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   One or more configuration blocks with distribution settings. Detailed below.

	   The following arguments are optional:
	*/
	Distributions []types.Imagebuilder_DistributionConfigurationDistribution `json:"distributions,omitempty" yaml:"distributions,omitempty"`

	// Name of the distribution configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags for the distribution configuration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
