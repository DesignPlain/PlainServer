package types

type Dataplex_LakeMetastore struct {
	// Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: `projects/{project_id}/locations/{location_id}/services/{service_id}`
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
