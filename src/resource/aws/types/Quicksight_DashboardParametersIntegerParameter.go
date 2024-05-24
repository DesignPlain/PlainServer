package types

type Quicksight_DashboardParametersIntegerParameter struct {
	//
	Values []int `json:"values,omitempty" yaml:"values,omitempty"`

	// Display name for the dashboard.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
