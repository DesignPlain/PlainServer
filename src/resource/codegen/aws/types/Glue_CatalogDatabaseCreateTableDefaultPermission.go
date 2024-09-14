package types

type Glue_CatalogDatabaseCreateTableDefaultPermission struct {
	// The permissions that are granted to the principal.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// The principal who is granted permissions.. See `principal` below.
	Principal Glue_CatalogDatabaseCreateTableDefaultPermissionPrincipal `json:"principal,omitempty" yaml:"principal,omitempty"`
}
