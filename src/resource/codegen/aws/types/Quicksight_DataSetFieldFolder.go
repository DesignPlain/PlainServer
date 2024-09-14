package types

type Quicksight_DataSetFieldFolder struct {
	// Key of the field folder map.
	FieldFoldersId string `json:"fieldFoldersId,omitempty" yaml:"fieldFoldersId,omitempty"`

	// An array of column names to add to the folder. A column can only be in one folder.
	Columns []string `json:"columns,omitempty" yaml:"columns,omitempty"`

	// Field folder description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
