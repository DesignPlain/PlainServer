package types

type Secretsmanager_getSecretsFilter struct {
	// Name of the filter field. Valid values can be found in the [Secrets Manager ListSecrets API Reference](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
