package types

type Bigquery_DatasetAccessAuthorizedDataset struct {
	/*
	   The dataset this entry applies to
	   Structure is documented below.
	*/
	Dataset Bigquery_DatasetAccessAuthorizedDatasetDataset `json:"dataset,omitempty" yaml:"dataset,omitempty"`

	/*
	   Which resources in the dataset this entry applies to. Currently, only views are supported,
	   but additional target types may be added in the future. Possible values: VIEWS
	*/
	TargetTypes []string `json:"targetTypes,omitempty" yaml:"targetTypes,omitempty"`
}
