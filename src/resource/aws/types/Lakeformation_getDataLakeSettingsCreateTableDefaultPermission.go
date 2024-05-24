package types

type Lakeformation_getDataLakeSettingsCreateTableDefaultPermission struct {
	// List of permissions granted to the principal.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Principal who is granted permissions.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
