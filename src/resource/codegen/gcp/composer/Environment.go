package composer

import types "libds/gcp/types"

type Environment struct {
	/*
	   User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
	   are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
	   must conform to the following regular expression: [a-z]([-a-z0-9]-[a-z0-9])?. Label values must be between 0 and 63
	   characters long and must conform to the regular expression ([a-z]([-a-z0-9]-[a-z0-9])?)?. No more than 64 labels can be
	   associated with a given environment. Both keys and values must be <= 128 bytes in size. --Note--: This field is
	   non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	   'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Name of the environment.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The location or Compute Engine region for the environment.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Configuration options for storage used by Composer environment.
	StorageConfig types.Composer_EnvironmentStorageConfig `json:"storageConfig,omitempty" yaml:"storageConfig,omitempty"`

	// Configuration parameters for this environment.
	Config types.Composer_EnvironmentConfig `json:"config,omitempty" yaml:"config,omitempty"`
}
