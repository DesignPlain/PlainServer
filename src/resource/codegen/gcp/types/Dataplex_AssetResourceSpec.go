package types

type Dataplex_AssetResourceSpec struct {
	// Optional. Determines how read permissions are handled for each asset and their associated tables. Only available to storage buckets assets. Possible values: DIRECT, MANAGED
	ReadAccessMode string `json:"readAccessMode,omitempty" yaml:"readAccessMode,omitempty"`

	/*
	   Required. Immutable. Type of resource. Possible values: STORAGE_BUCKET, BIGQUERY_DATASET

	   - - -
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Immutable. Relative name of the cloud resource that contains the data that is being managed within a lake. For example: `projects/{project_number}/buckets/{bucket_id}` `projects/{project_number}/datasets/{dataset_id}`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
