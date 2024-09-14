package certificatemanager

type DnsAuthorization struct {
	/*
	   Set of label tags associated with the DNS Authorization resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The Certificate Manager location. If not specified, "global" is used.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is created.
	   The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]- which means the first character must be a letter,
	   and all following characters must be a dash, underscore, letter or digit.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A domain which is being authorized. A DnsAuthorization resource covers a
	   single domain and its wildcard, e.g. authorization for "example.com" can
	   be used to issue certificates for "example.com" and "-.example.com".
	*/
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
