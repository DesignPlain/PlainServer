package alloydb

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	/*
	   User-defined labels for the alloydb instance.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
	   Structure is documented below.
	*/
	ReadPoolConfig types.Alloydb_InstanceReadPoolConfig `json:"readPoolConfig,omitempty" yaml:"readPoolConfig,omitempty"`

	/*
	   Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   Client connection specific configurations.
	   Structure is documented below.
	*/
	ClientConnectionConfig types.Alloydb_InstanceClientConnectionConfig `json:"clientConnectionConfig,omitempty" yaml:"clientConnectionConfig,omitempty"`

	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone string `json:"gceZone,omitempty" yaml:"gceZone,omitempty"`

	/*
	   The ID of the alloydb instance.


	   - - -
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the
	   'depends_on' meta-data attribute. If the instance type is SECONDARY, point to the cluster_type of the associated
	   secondary cluster instead of mentioning SECONDARY. Example: {instance_type =
	   google_alloydb_cluster.<secondary_cluster_name>.cluster_type} instead of {instance_type = SECONDARY} If the instance
	   type is SECONDARY, the terraform delete instance operation does not delete the secondary instance but abandons it
	   instead. Use deletion_policy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete
	   the secondary cluster as well its associated secondary instance. Users can undo the delete secondary instance action by
	   importing the deleted secondary instance by calling terraform import. Possible values: ["PRIMARY", "READ_POOL",
	   "SECONDARY"]
	*/
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	/*
	   Configuration for query insights.
	   Structure is documented below.
	*/
	QueryInsightsConfig types.Alloydb_InstanceQueryInsightsConfig `json:"queryInsightsConfig,omitempty" yaml:"queryInsightsConfig,omitempty"`

	/*
	   'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
	   Note that primary and read instances can have different availability types.
	   Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
	   Zone is automatically chosen from the list of zones in the region specified.
	   Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
	   can have regional availability (nodes are present in 2 or more zones in a region).'
	   Possible values are: `AVAILABILITY_TYPE_UNSPECIFIED`, `ZONAL`, `REGIONAL`.
	*/
	AvailabilityType string `json:"availabilityType,omitempty" yaml:"availabilityType,omitempty"`

	/*
	   Identifies the alloydb cluster. Must be in the format
	   'projects/{project}/locations/{location}/clusters/{cluster_id}'
	*/
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// Database flags. Set at instance level. - They are copied from primary instance on read instance creation. - Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty" yaml:"databaseFlags,omitempty"`

	// User-settable and human-readable display name for the Instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Configurations for the machines that host the underlying database engine.
	   Structure is documented below.
	*/
	MachineConfig types.Alloydb_InstanceMachineConfig `json:"machineConfig,omitempty" yaml:"machineConfig,omitempty"`
}
