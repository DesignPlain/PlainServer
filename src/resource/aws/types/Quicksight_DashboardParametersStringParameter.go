package types

type Quicksight_DashboardParametersStringParameter struct {
	// Display name for the dashboard.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
