package certificatemanager

type CertificateMapEntry struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A set of Certificates defines for the given hostname.
	   There can be defined up to fifteen certificates in each Certificate Map Entry.
	   Each certificate must match pattern projects/-/locations/-/certificates/-.
	*/
	Certificates []string `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (-.example.com)
	   for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	   selecting a proper certificate.
	*/
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   Set of labels associated with a Certificate Map Entry.
	   An object containing a list of "key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   A map entry that is inputted into the cetrificate map


	   - - -
	*/
	Map string `json:"map,omitempty" yaml:"map,omitempty"`

	// A predefined matcher for particular cases, other than SNI selection
	Matcher string `json:"matcher,omitempty" yaml:"matcher,omitempty"`

	/*
	   A user-defined name of the Certificate Map Entry. Certificate Map Entry
	   names must be unique globally and match pattern
	   'projects/-/locations/-/certificateMaps/-/certificateMapEntries/-'
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
