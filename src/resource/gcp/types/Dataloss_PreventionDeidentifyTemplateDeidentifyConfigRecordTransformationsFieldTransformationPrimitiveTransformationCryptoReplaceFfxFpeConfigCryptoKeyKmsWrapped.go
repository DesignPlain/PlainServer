package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped struct {
	/*
	   The wrapped data crypto key.
	   A base64-encoded string.
	*/
	WrappedKey string `json:"wrappedKey,omitempty" yaml:"wrappedKey,omitempty"`

	// The resource name of the KMS CryptoKey to use for unwrapping.
	CryptoKeyName string `json:"cryptoKeyName,omitempty" yaml:"cryptoKeyName,omitempty"`
}
