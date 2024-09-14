package signer

import types "libds/aws/types"

type SigningProfile struct {
	// A list of tags associated with the signing profile. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The ID of the platform that is used by the target signing profile.
	PlatformId string `json:"platformId,omitempty" yaml:"platformId,omitempty"`

	// The validity period for a signing job. See `signature_validity_period` Block below for details.
	SignatureValidityPeriod types.Signer_SigningProfileSignatureValidityPeriod `json:"signatureValidityPeriod,omitempty" yaml:"signatureValidityPeriod,omitempty"`

	// The AWS Certificate Manager certificate that will be used to sign code with the new signing profile. See `signing_material` Block below for details.
	SigningMaterial types.Signer_SigningProfileSigningMaterial `json:"signingMaterial,omitempty" yaml:"signingMaterial,omitempty"`
}
