package bigtable

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	   It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	   and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	   `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	   is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	*/
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	/*
	   A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance. Must be 6-33 characters and must only contain hyphens, lowercase letters and numbers.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A block of cluster configuration options. This can be specified at least once, and up
	   to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
	   to default to the backend value. See structure below.

	   -----
	*/
	Clusters []types.Bigtable_InstanceCluster `json:"clusters,omitempty" yaml:"clusters,omitempty"`

	/*
	   Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	   in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
}
