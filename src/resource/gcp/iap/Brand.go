package iap

type Brand struct {
	/*
	   Application name displayed on OAuth consent screen.


	   - - -
	*/
	ApplicationTitle string `json:"applicationTitle,omitempty" yaml:"applicationTitle,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Support email displayed on the OAuth consent screen. Can be either a
	   user or group email. When a user email is specified, the caller must
	   be the user with the associated email address. When a group email is
	   specified, the caller can be either a user or a service account which
	   is an owner of the specified group in Cloud Identity.
	*/
	SupportEmail string `json:"supportEmail,omitempty" yaml:"supportEmail,omitempty"`
}
