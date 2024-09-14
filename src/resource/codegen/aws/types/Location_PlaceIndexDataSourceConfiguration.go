package types

type Location_PlaceIndexDataSourceConfiguration struct {
	// Specifies how the results of an operation will be stored by the caller. Valid values: `SingleUse`, `Storage`. Default: `SingleUse`.
	IntendedUse string `json:"intendedUse,omitempty" yaml:"intendedUse,omitempty"`
}
