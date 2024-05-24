package location

type RouteCalculation struct {
	// The name of the route calculator resource.
	CalculatorName string `json:"calculatorName,omitempty" yaml:"calculatorName,omitempty"`

	/*
	   Specifies the data provider of traffic and road network data.

	   The following arguments are optional:
	*/
	DataSource string `json:"dataSource,omitempty" yaml:"dataSource,omitempty"`

	// The optional description for the route calculator resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
