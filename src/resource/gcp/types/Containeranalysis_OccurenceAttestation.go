package types

type Containeranalysis_OccurenceAttestation struct {
	/*
	   The serialized payload that is verified by one or
	   more signatures. A base64-encoded string.
	*/
	SerializedPayload string `json:"serializedPayload,omitempty" yaml:"serializedPayload,omitempty"`

	/*
	   One or more signatures over serializedPayload.
	   Verifier implementations should consider this attestation
	   message verified if at least one signature verifies
	   serializedPayload. See Signature in common.proto for more
	   details on signature structure and verification.
	   Structure is documented below.
	*/
	Signatures []Containeranalysis_OccurenceAttestationSignature `json:"signatures,omitempty" yaml:"signatures,omitempty"`
}
