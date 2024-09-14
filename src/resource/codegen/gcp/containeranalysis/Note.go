package containeranalysis

import types "libds/gcp/types"

type Note struct {
	/*
	   URLs associated with this note and related metadata.
	   Structure is documented below.
	*/
	RelatedUrls []types.Containeranalysis_NoteRelatedUrl `json:"relatedUrls,omitempty" yaml:"relatedUrls,omitempty"`

	// A one sentence description of the note.
	ShortDescription string `json:"shortDescription,omitempty" yaml:"shortDescription,omitempty"`

	/*
	   Note kind that represents a logical attestation "role" or "authority".
	   For example, an organization might have one AttestationAuthority for
	   "QA" and one for "build". This Note is intended to act strictly as a
	   grouping mechanism for the attached Occurrences (Attestations). This
	   grouping mechanism also provides a security boundary, since IAM ACLs
	   gate the ability for a principle to attach an Occurrence to a given
	   Note. It also provides a single point of lookup to find all attached
	   Attestation Occurrences, even if they don't all live in the same
	   project.
	   Structure is documented below.
	*/
	AttestationAuthority types.Containeranalysis_NoteAttestationAuthority `json:"attestationAuthority,omitempty" yaml:"attestationAuthority,omitempty"`

	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	// A detailed description of the note
	LongDescription string `json:"longDescription,omitempty" yaml:"longDescription,omitempty"`

	// The name of the note.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Names of other notes related to this note.
	RelatedNoteNames []string `json:"relatedNoteNames,omitempty" yaml:"relatedNoteNames,omitempty"`
}
