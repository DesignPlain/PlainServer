package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped struct {
	// The resource name of the KMS CryptoKey to use for unwrapping.
	CryptoKeyName string `json:"cryptoKeyName,omitempty" yaml:"cryptoKeyName,omitempty"`

	/*
	   The wrapped data crypto key.
	   A base64-encoded string.
	*/
	WrappedKey string `json:"wrappedKey,omitempty" yaml:"wrappedKey,omitempty"`
}
