package types



type Bigquery_getDatasetAccessDataset struct {
	// The dataset this entry applies to
	Datasets []Bigquery_getDatasetAccessDatasetDataset `json:"datasets,omitempty" yaml:"datasets,omitempty"`

	/*
	   Which resources in the dataset this entry applies to. Currently, only views are supported,
	   but additional target types may be added in the future. Possible values: VIEWS
	*/
	TargetTypes []string `json:"targetTypes,omitempty" yaml:"targetTypes,omitempty"`
}
