package types

type Databasemigrationservice_ConnectionProfileAlloydbSettingsInitialUser struct {
	/*
	   The initial password for the user.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	/*
	   (Output)
	   Output only. Indicates if the initialUser.password field has been set.
	*/
	PasswordSet bool `json:"passwordSet,omitempty" yaml:"passwordSet,omitempty"`

	// The database username.
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
