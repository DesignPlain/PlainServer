package types

type Sagemaker_DataQualityJobDefinitionJobResourcesClusterConfig struct {
	// The number of ML compute instances to use in the model monitoring job. For distributed processing jobs, specify a value greater than 1.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// The ML compute instance type for the processing job.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	VolumeKmsKeyId string `json:"volumeKmsKeyId,omitempty" yaml:"volumeKmsKeyId,omitempty"`

	// The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML storage for your scenario.
	VolumeSizeInGb int `json:"volumeSizeInGb,omitempty" yaml:"volumeSizeInGb,omitempty"`
}
