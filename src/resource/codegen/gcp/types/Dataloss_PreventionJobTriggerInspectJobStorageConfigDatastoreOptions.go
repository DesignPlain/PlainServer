package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigDatastoreOptions struct {
	/*
	   A representation of a Datastore kind.
	   Structure is documented below.
	*/
	Kind Dataloss_PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKind `json:"kind,omitempty" yaml:"kind,omitempty"`

	/*
	   Datastore partition ID. A partition ID identifies a grouping of entities. The grouping
	   is always by project and namespace, however the namespace ID may be empty.
	   Structure is documented below.
	*/
	PartitionId Dataloss_PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId `json:"partitionId,omitempty" yaml:"partitionId,omitempty"`
}
