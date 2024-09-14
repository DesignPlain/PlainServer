package identityplatform

import types "libds/gcp/types"

type ProjectDefaultConfig struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Configuration related to local sign in methods.
	   Structure is documented below.
	*/
	SignIn types.Identityplatform_ProjectDefaultConfigSignIn `json:"signIn,omitempty" yaml:"signIn,omitempty"`
}
