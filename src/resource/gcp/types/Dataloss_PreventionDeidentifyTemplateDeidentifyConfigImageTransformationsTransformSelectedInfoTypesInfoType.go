package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypesInfoType struct {
	// Name of the information type.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Optional custom sensitivity for this InfoType. This only applies to data profiling.
	   Structure is documented below.
	*/
	SensitivityScore Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypesInfoTypeSensitivityScore `json:"sensitivityScore,omitempty" yaml:"sensitivityScore,omitempty"`

	// Version name for this InfoType.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
