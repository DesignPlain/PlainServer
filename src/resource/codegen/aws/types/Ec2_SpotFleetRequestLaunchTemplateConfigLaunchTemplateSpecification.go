package types

type Ec2_SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecification struct {
	// The ID of the launch template. Conflicts with `name`.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The name of the launch template. Conflicts with `id`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Template version. Unlike the autoscaling equivalent, does not support `$Latest` or `$Default`, so use the launch_template resource's attribute, e.g., `"${aws_launch_template.foo.latest_version}"`. It will use the default version if omitted.

	   --Note:-- The specified launch template can specify only a subset of the
	   inputs of `aws.ec2.LaunchTemplate`.  There are limitations on
	   what you can specify as spot fleet does not support all the attributes that are supported by autoscaling groups. [AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#launch-templates-spot-fleet) is currently sparse, but at least `instance_initiated_shutdown_behavior` is confirmed unsupported.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
