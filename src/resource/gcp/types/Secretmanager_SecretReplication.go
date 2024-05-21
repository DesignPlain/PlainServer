package types

type Secretmanager_SecretReplication struct {
	/*
	   The Secret will automatically be replicated without any restrictions.
	   Structure is documented below.
	*/
	Auto Secretmanager_SecretReplicationAuto `json:"auto,omitempty" yaml:"auto,omitempty"`

	/*
	   The Secret will be replicated to the regions specified by the user.
	   Structure is documented below.
	*/
	UserManaged Secretmanager_SecretReplicationUserManaged `json:"userManaged,omitempty" yaml:"userManaged,omitempty"`
}
