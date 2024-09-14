package apprunner

import types "libds/aws/types"

type Service struct {
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
	EncryptionConfiguration types.Apprunner_ServiceEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
	HealthCheckConfiguration types.Apprunner_ServiceHealthCheckConfiguration `json:"healthCheckConfiguration,omitempty" yaml:"healthCheckConfiguration,omitempty"`

	// The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
	InstanceConfiguration types.Apprunner_ServiceInstanceConfiguration `json:"instanceConfiguration,omitempty" yaml:"instanceConfiguration,omitempty"`

	// Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
	NetworkConfiguration types.Apprunner_ServiceNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// The observability configuration of your service. See Observability Configuration below for more details.
	ObservabilityConfiguration types.Apprunner_ServiceObservabilityConfiguration `json:"observabilityConfiguration,omitempty" yaml:"observabilityConfiguration,omitempty"`

	// ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	AutoScalingConfigurationArn string `json:"autoScalingConfigurationArn,omitempty" yaml:"autoScalingConfigurationArn,omitempty"`

	// Name of the service.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	/*
	   The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.

	   The following arguments are optional:
	*/
	SourceConfiguration types.Apprunner_ServiceSourceConfiguration `json:"sourceConfiguration,omitempty" yaml:"sourceConfiguration,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
