package types

type Ecs_TaskSetServiceRegistries struct {
	// The container name value, already specified in the task definition, to be used for your service discovery service.
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`

	// The port value, already specified in the task definition, to be used for your service discovery service.
	ContainerPort int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`

	// The port value used if your Service Discovery service specified an SRV record.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The ARN of the Service Registry. The currently supported service registry is Amazon Route 53 Auto Naming Service(`aws.servicediscovery.Service` resource). For more information, see [Service](https://docs.aws.amazon.com/Route53/latest/APIReference/API_autonaming_Service.html).
	RegistryArn string `json:"registryArn,omitempty" yaml:"registryArn,omitempty"`
}
