package types

type Apigateway_DomainNameEndpointConfiguration struct {
	// List of endpoint  This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
	Types string `json:"types,omitempty" yaml:"types,omitempty"`
}
