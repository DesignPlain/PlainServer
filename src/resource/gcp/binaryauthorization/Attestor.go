package binaryauthorization

import types "DesignSphere_Server/src/resource/gcp/types"

type Attestor struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	   Structure is documented below.
	*/
	AttestationAuthorityNote types.Binaryauthorization_AttestorAttestationAuthorityNote `json:"attestationAuthorityNote,omitempty" yaml:"attestationAuthorityNote,omitempty"`

	/*
	   A descriptive comment. This field may be updated. The field may be
	   displayed in chooser dialogs.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The resource name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
