package types

type Sagemaker_ModelPrimaryContainer struct {
	// The location of model data to deploy. Use this for uncompressed model deployment. For information about how to deploy an uncompressed model, see [Deploying uncompressed models](https://docs.aws.amazon.com/sagemaker/latest/dg/large-model-inference-uncompressed.html) in the _AWS SageMaker Developer Guide_.
	ModelDataSource Sagemaker_ModelPrimaryContainerModelDataSource `json:"modelDataSource,omitempty" yaml:"modelDataSource,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	ModelDataUrl string `json:"modelDataUrl,omitempty" yaml:"modelDataUrl,omitempty"`

	// The Amazon Resource Name (ARN) of the model package to use to create the model.
	ModelPackageName string `json:"modelPackageName,omitempty" yaml:"modelPackageName,omitempty"`

	// The DNS host name for the container.
	ContainerHostname string `json:"containerHostname,omitempty" yaml:"containerHostname,omitempty"`

	/*
	   Environment variables for the Docker container.
	   A list of key value pairs.
	*/
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see [Using a Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html). see Image Config.
	ImageConfig Sagemaker_ModelPrimaryContainerImageConfig `json:"imageConfig,omitempty" yaml:"imageConfig,omitempty"`

	// The container hosts value `SingleModel/MultiModel`. The default value is `SingleModel`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
