package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType struct {
	// Optional version name for this InfoType.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Optional custom sensitivity for this InfoType. This only applies to data profiling.
	   Structure is documented below.
	*/
	SensitivityScore Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoTypeSensitivityScore `json:"sensitivityScore,omitempty" yaml:"sensitivityScore,omitempty"`
}
