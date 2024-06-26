package types

type Assuredworkloads_WorkloadResource struct {
	// Resource identifier. For a project this represents projectId. If the project is already taken, the workload creation will fail. For KeyRing, this represents the keyring_id. For a folder, don't set this value as folder_id is assigned by Google.
	ResourceId int `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}
