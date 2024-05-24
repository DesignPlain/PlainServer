package types

type Quicksight_getDataSetPermission struct {
	//
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`

	//
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
