package types

type Ecs_ClusterServiceConnectDefaults struct {
	// ARN of the `aws.servicediscovery.HttpNamespace` that's used when you create a service and don't specify a Service Connect configuration.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
