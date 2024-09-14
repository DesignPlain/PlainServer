package types

type Quicksight_getDataSetFieldFolder struct {
	//
	Columns []string `json:"columns,omitempty" yaml:"columns,omitempty"`

	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	FieldFoldersId string `json:"fieldFoldersId,omitempty" yaml:"fieldFoldersId,omitempty"`
}
