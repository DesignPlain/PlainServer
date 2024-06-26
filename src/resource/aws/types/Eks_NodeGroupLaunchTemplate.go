package types

type Eks_NodeGroupLaunchTemplate struct {
	// Identifier of the EC2 Launch Template. Conflicts with `name`.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the EC2 Launch Template. Conflicts with `id`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g., `1`) on read and the provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
