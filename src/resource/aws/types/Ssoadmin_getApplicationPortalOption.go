package types

type Ssoadmin_getApplicationPortalOption struct {
	//
	SignInOptions []Ssoadmin_getApplicationPortalOptionSignInOption `json:"signInOptions,omitempty" yaml:"signInOptions,omitempty"`

	//
	Visibility string `json:"visibility,omitempty" yaml:"visibility,omitempty"`
}
