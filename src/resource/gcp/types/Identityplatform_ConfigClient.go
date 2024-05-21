package types

type Identityplatform_ConfigClient struct {
	/*
	   (Output)
	   API key that can be used when making requests for this project.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	/*
	   (Output)
	   Firebase subdomain.
	*/
	FirebaseSubdomain string `json:"firebaseSubdomain,omitempty" yaml:"firebaseSubdomain,omitempty"`

	/*
	   Configuration related to restricting a user's ability to affect their account.
	   Structure is documented below.
	*/
	Permissions Identityplatform_ConfigClientPermissions `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}
