package containeranalysis

import types "libds/gcp/types"

type Occurence struct {
	/*
	   Occurrence that represents a single "attestation". The authenticity
	   of an attestation can be verified using the attached signature.
	   If the verifier trusts the public key of the signer, then verifying
	   the signature is sufficient to establish trust. In this circumstance,
	   the authority to which this attestation is attached is primarily
	   useful for lookup (how to find this attestation if you already
	   know the authority and artifact to be verified) and intent (for
	   which authority this attestation was intended to sign.
	   Structure is documented below.
	*/
	Attestation types.Containeranalysis_OccurenceAttestation `json:"attestation,omitempty" yaml:"attestation,omitempty"`

	/*
	   The analysis note associated with this occurrence, in the form of
	   projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	   filter in list requests.
	*/
	NoteName string `json:"noteName,omitempty" yaml:"noteName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A description of actions that can be taken to remedy the note.
	Remediation string `json:"remediation,omitempty" yaml:"remediation,omitempty"`

	/*
	   Required. Immutable. A URI that represents the resource for which
	   the occurrence applies. For example,
	   https://gcr.io/project/image@sha256:123abc for a Docker image.
	*/
	ResourceUri string `json:"resourceUri,omitempty" yaml:"resourceUri,omitempty"`
}
