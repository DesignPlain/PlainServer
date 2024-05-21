package kms

import types "DesignSphere_Server/src/resource/gcp/types"

type CryptoKey struct {
	/*
	   Labels with user-defined metadata to apply to this resource.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The immutable purpose of this CryptoKey. See the
	   [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	   for possible inputs.
	   Default value is "ENCRYPT_DECRYPT".
	*/
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`

	/*
	   Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	   The first rotation will take place after the specified period. The rotation period has
	   the format of a decimal number with up to 9 fractional digits, followed by the
	   letter `s` (seconds). It must be greater than a day (ie, 86400).
	*/
	RotationPeriod string `json:"rotationPeriod,omitempty" yaml:"rotationPeriod,omitempty"`

	/*
	   If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	   You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
	*/
	SkipInitialVersionCreation bool `json:"skipInitialVersionCreation,omitempty" yaml:"skipInitialVersionCreation,omitempty"`

	/*
	   A template describing settings for new crypto key versions.
	   Structure is documented below.
	*/
	VersionTemplate types.Kms_CryptoKeyVersionTemplate `json:"versionTemplate,omitempty" yaml:"versionTemplate,omitempty"`

	/*
	   The KeyRing that this key belongs to.
	   Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.


	   - - -
	*/
	KeyRing string `json:"keyRing,omitempty" yaml:"keyRing,omitempty"`

	// Whether this key may contain imported versions only.
	ImportOnly bool `json:"importOnly,omitempty" yaml:"importOnly,omitempty"`

	// The resource name for the CryptoKey.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	   If not specified at creation time, the default duration is 24 hours.
	*/
	DestroyScheduledDuration string `json:"destroyScheduledDuration,omitempty" yaml:"destroyScheduledDuration,omitempty"`
}
