package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes struct {
	/*
	   InfoTypes to apply the transformation to. Leaving this empty will apply the transformation to apply to
	   all findings that correspond to infoTypes that were requested in InspectConfig.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypesInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`
}
