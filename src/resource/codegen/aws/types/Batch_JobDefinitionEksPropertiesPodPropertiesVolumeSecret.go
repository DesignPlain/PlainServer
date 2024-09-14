package types

type Batch_JobDefinitionEksPropertiesPodPropertiesVolumeSecret struct {
	// Whether the secret or the secret's keys must be defined.
	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`

	// Name of the secret. The name must be allowed as a DNS subdomain name.
	SecretName string `json:"secretName,omitempty" yaml:"secretName,omitempty"`
}
