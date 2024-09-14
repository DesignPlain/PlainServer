package types

type Quicksight_DashboardParametersIntegerParameter struct {
	// Display name for the dashboard.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []int `json:"values,omitempty" yaml:"values,omitempty"`
}
