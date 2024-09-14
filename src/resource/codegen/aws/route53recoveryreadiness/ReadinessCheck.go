package route53recoveryreadiness

type ReadinessCheck struct {
	// Unique name describing the readiness check.
	ReadinessCheckName string `json:"readinessCheckName,omitempty" yaml:"readinessCheckName,omitempty"`

	/*
	   Name describing the resource set that will be monitored for readiness.

	   The following arguments are optional:
	*/
	ResourceSetName string `json:"resourceSetName,omitempty" yaml:"resourceSetName,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
