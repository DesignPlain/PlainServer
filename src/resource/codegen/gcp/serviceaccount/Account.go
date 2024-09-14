package serviceaccount

type Account struct {
	/*
	   The ID of the project that the service account will be created in.
	   Defaults to the provider project configuration.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The account id that is used to generate the service
	   account email address and a stable unique id. It is unique within a project,
	   must be 6-30 characters long, and match the regular expression `a-z`
	   to comply with RFC1035. Changing this forces a new service account to be created.
	*/
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists bool `json:"createIgnoreAlreadyExists,omitempty" yaml:"createIgnoreAlreadyExists,omitempty"`

	/*
	   A text description of the service account.
	   Must be less than or equal to 256 UTF-8 bytes.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
	   Must be set after creation to disable a service account.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   The display name for the service account.
	   Can be updated without creating a new resource.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
