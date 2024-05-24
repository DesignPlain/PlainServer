package appstream

type User struct {
	// Last name, or surname, of the user.
	LastName string `json:"lastName,omitempty" yaml:"lastName,omitempty"`

	// Send an email notification.
	SendEmailNotification bool `json:"sendEmailNotification,omitempty" yaml:"sendEmailNotification,omitempty"`

	/*
	   Email address of the user.

	   The following arguments are optional:
	*/
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// Authentication type for the user. You must specify USERPOOL. Valid values: `API`, `SAML`, `USERPOOL`
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// Whether the user in the user pool is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// First name, or given name, of the user.
	FirstName string `json:"firstName,omitempty" yaml:"firstName,omitempty"`
}
