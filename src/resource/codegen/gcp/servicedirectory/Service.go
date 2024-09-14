package servicedirectory

type Service struct {
	// The resource name of the namespace this service will belong to.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	/*
	   The Resource ID must be 1-63 characters long, including digits,
	   lowercase letters or the hyphen character.


	   - - -
	*/
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`

	/*
	   Metadata for the service. This data can be consumed
	   by service clients. The entire metadata dictionary may contain
	   up to 2000 characters, spread across all key-value pairs.
	   Metadata that goes beyond any these limits will be rejected.
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}
