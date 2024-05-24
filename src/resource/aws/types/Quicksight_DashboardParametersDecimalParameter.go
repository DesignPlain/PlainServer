package types

type Quicksight_DashboardParametersDecimalParameter struct {
	//
	Values []float64 `json:"values,omitempty" yaml:"values,omitempty"`

	// Display name for the dashboard.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
