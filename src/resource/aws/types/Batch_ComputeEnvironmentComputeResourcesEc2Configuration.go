package types

type Batch_ComputeEnvironmentComputeResourcesEc2Configuration struct {
	// The AMI ID used for instances launched in the compute environment that match the image type. This setting overrides the `image_id` argument in the `compute_resources` block.
	ImageIdOverride string `json:"imageIdOverride,omitempty" yaml:"imageIdOverride,omitempty"`

	// The image type to match with the instance type to select an AMI. If the `image_id_override` parameter isn't specified, then a recent [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) (`ECS_AL2`) is used.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`
}
