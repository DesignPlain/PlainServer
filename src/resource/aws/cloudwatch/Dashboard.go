package cloudwatch

type Dashboard struct {
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody string `json:"dashboardBody,omitempty" yaml:"dashboardBody,omitempty"`

	// The name of the dashboard.
	DashboardName string `json:"dashboardName,omitempty" yaml:"dashboardName,omitempty"`
}
