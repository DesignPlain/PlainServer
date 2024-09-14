package monitoring

type Dashboard struct {
	/*
	   The JSON representation of a dashboard, following the format at
	   https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	*/
	DashboardJson string `json:"dashboardJson,omitempty" yaml:"dashboardJson,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
