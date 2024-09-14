package networkservices

import types "libds/gcp/types"

type EdgeCacheKeyset struct {
	/*
	   An ordered list of Ed25519 public keys to use for validating signed requests.
	   You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
	   You may specify no more than one Google-managed public key.
	   If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
	   Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
	   Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
	   Structure is documented below.
	*/
	PublicKeys []types.Networkservices_EdgeCacheKeysetPublicKey `json:"publicKeys,omitempty" yaml:"publicKeys,omitempty"`

	/*
	   An ordered list of shared keys to use for validating signed requests.
	   Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
	   You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
	   You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
	   Structure is documented below.
	*/
	ValidationSharedKeys []types.Networkservices_EdgeCacheKeysetValidationSharedKey `json:"validationSharedKeys,omitempty" yaml:"validationSharedKeys,omitempty"`

	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Set of label tags associated with the EdgeCache resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is created.
	   The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]- which means the first character must be a letter,
	   and all following characters must be a dash, underscore, letter or digit.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
