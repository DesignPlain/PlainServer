package types

type Route53_getQueryLogConfigFilter struct {
	// The name of the query logging configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
