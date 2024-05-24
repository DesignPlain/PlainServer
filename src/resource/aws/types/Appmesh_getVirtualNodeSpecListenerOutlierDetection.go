package types

type Appmesh_getVirtualNodeSpecListenerOutlierDetection struct {
	//
	MaxEjectionPercent int `json:"maxEjectionPercent,omitempty" yaml:"maxEjectionPercent,omitempty"`

	//
	MaxServerErrors int `json:"maxServerErrors,omitempty" yaml:"maxServerErrors,omitempty"`

	//
	BaseEjectionDurations []Appmesh_getVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration `json:"baseEjectionDurations,omitempty" yaml:"baseEjectionDurations,omitempty"`

	//
	Intervals []Appmesh_getVirtualNodeSpecListenerOutlierDetectionInterval `json:"intervals,omitempty" yaml:"intervals,omitempty"`
}
