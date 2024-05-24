package types

type Sagemaker_DataQualityJobDefinitionNetworkConfig struct {
	// Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. Fields are documented below.
	VpcConfig Sagemaker_DataQualityJobDefinitionNetworkConfigVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Whether to encrypt all communications between the instances used for the monitoring jobs. Choose `true` to encrypt communications. Encryption provides greater security for distributed jobs, but the processing might take longer.
	EnableInterContainerTrafficEncryption bool `json:"enableInterContainerTrafficEncryption,omitempty" yaml:"enableInterContainerTrafficEncryption,omitempty"`

	// Whether to allow inbound and outbound network calls to and from the containers used for the monitoring job.
	EnableNetworkIsolation bool `json:"enableNetworkIsolation,omitempty" yaml:"enableNetworkIsolation,omitempty"`
}
