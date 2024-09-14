package xray

type SamplingRule struct {
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Matches attributes derived from the request.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate float64 `json:"fixedRate,omitempty" yaml:"fixedRate,omitempty"`

	// Matches the HTTP method of a request.
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	// The priority of the sampling rule.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The name of the sampling rule.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	// Matches the hostname from a request URL.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize int `json:"reservoirSize,omitempty" yaml:"reservoirSize,omitempty"`

	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType string `json:"serviceType,omitempty" yaml:"serviceType,omitempty"`

	// Matches the path from a request URL.
	UrlPath string `json:"urlPath,omitempty" yaml:"urlPath,omitempty"`

	// The version of the sampling rule format (`1` )
	Version int `json:"version,omitempty" yaml:"version,omitempty"`
}
