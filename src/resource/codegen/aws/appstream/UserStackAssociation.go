package appstream

type UserStackAssociation struct {
	// Authentication type for the user.
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// Whether a welcome email is sent to a user after the user is created in the user pool.
	SendEmailNotification bool `json:"sendEmailNotification,omitempty" yaml:"sendEmailNotification,omitempty"`

	// Name of the stack that is associated with the user.
	StackName string `json:"stackName,omitempty" yaml:"stackName,omitempty"`

	/*
	   Email address of the user who is associated with the stack.

	   The following arguments are optional:
	*/
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
