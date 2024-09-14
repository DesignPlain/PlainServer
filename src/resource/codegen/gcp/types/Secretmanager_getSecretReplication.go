package types

type Secretmanager_getSecretReplication struct {
	// The Secret will automatically be replicated without any restrictions.
	Autos []Secretmanager_getSecretReplicationAuto `json:"autos,omitempty" yaml:"autos,omitempty"`

	// The Secret will be replicated to the regions specified by the user.
	UserManageds []Secretmanager_getSecretReplicationUserManaged `json:"userManageds,omitempty" yaml:"userManageds,omitempty"`
}
