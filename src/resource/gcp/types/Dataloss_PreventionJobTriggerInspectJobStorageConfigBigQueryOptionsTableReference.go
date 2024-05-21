package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference struct {
	// The dataset ID of the table.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// The name of the table.
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`
}
