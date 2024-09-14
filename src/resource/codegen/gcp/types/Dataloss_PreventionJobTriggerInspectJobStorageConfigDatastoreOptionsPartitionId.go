package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId struct {
	// If not empty, the ID of the namespace to which the entities belong.
	NamespaceId string `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`

	// The ID of the project to which the entities belong.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
