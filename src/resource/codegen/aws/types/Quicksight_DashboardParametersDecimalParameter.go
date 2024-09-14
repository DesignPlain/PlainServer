package types

type Quicksight_DashboardParametersDecimalParameter struct {
	// Display name for the dashboard.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []float64 `json:"values,omitempty" yaml:"values,omitempty"`
}
