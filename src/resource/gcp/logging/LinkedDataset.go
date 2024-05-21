package logging

import types "DesignSphere_Server/src/resource/gcp/types"

type LinkedDataset struct {
	// The id of the linked dataset.
	LinkId string `json:"linkId,omitempty" yaml:"linkId,omitempty"`

	// The location of the linked dataset.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The parent of the linked dataset.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along
	   with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery
	   Views corresponding to the LogViews in the bucket.
	   Structure is documented below.
	*/
	BigqueryDatasets []types.Logging_LinkedDatasetBigqueryDataset `json:"bigqueryDatasets,omitempty" yaml:"bigqueryDatasets,omitempty"`

	/*
	   The bucket to which the linked dataset is attached.


	   - - -
	*/
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Describes this link. The maximum length of the description is 8000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
