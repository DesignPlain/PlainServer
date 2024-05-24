package types

type Quicksight_DataSetDataSetUsageConfiguration struct {
	// Controls whether a child dataset of a direct query can use this dataset as a source.
	DisableUseAsDirectQuerySource bool `json:"disableUseAsDirectQuerySource,omitempty" yaml:"disableUseAsDirectQuerySource,omitempty"`

	// Controls whether a child dataset that's stored in QuickSight can use this dataset as a source.
	DisableUseAsImportedSource bool `json:"disableUseAsImportedSource,omitempty" yaml:"disableUseAsImportedSource,omitempty"`
}
