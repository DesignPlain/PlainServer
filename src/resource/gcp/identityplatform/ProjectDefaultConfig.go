package identityplatform

import types "DesignSphere_Server/src/resource/gcp/types"

type ProjectDefaultConfig struct {
	/*
	   Configuration related to local sign in methods.
	   Structure is documented below.
	*/
	SignIn types.Identityplatform_ProjectDefaultConfigSignIn `json:"signIn,omitempty" yaml:"signIn,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
