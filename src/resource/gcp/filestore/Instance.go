package filestore

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	// KMS key name used for data encryption.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The service tier of the instance.
	   Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	// A description of the instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   File system shares on the instance. For this version, only a
	   single file share is supported.
	   Structure is documented below.
	*/
	FileShares types.Filestore_InstanceFileShares `json:"fileShares,omitempty" yaml:"fileShares,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The resource name of the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   VPC networks to which the instance is connected. For this version,
	   only a single network is supported.
	   Structure is documented below.
	*/
	Networks []types.Filestore_InstanceNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   (Optional, Deprecated)
	   The name of the Filestore zone of the instance.

	   > --Warning:-- `zone` is deprecated and will be removed in a future major release. Use `location` instead.
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
