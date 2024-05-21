package types

type Secretmanager_getSecretsSecretReplication struct {
	/*
	   The Secret will automatically be replicated without any restrictions.
	   Structure is documented below.
	*/
	Autos []Secretmanager_getSecretsSecretReplicationAuto `json:"autos,omitempty" yaml:"autos,omitempty"`

	/*
	   The Secret will be replicated to the regions specified by the user.
	   Structure is documented below.
	*/
	UserManageds []Secretmanager_getSecretsSecretReplicationUserManaged `json:"userManageds,omitempty" yaml:"userManageds,omitempty"`
}
