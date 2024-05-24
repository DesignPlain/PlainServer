package types

type Quicksight_DataSourcePermission struct {
	// Set of IAM actions to grant or revoke permissions on. Max of 16 items.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`

	// The Amazon Resource Name (ARN) of the principal.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
