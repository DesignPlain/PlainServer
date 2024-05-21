package types

type Datastream_StreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset struct {
	/*
	   Dataset ID in the format projects/{project}/datasets/{dataset_id} or
	   {project}:{dataset_id}
	*/
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`
}
