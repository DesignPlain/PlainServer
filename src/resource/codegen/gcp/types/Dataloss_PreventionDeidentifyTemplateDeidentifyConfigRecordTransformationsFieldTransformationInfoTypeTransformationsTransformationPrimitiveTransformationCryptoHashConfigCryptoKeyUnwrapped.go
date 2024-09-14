package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformationPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped struct {
	/*
	   A 128/192/256 bit key.
	   A base64-encoded string.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
