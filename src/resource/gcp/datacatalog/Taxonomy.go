package datacatalog

type Taxonomy struct {
	/*
	   A list of policy types that are activated for this taxonomy. If not set,
	   defaults to an empty list.
	   Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
	*/
	ActivatedPolicyTypes []string `json:"activatedPolicyTypes,omitempty" yaml:"activatedPolicyTypes,omitempty"`

	/*
	   Description of this taxonomy. It must: contain only unicode characters,
	   tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
	   long when encoded in UTF-8. If not set, defaults to an empty description.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   User defined name of this taxonomy.
	   The taxonomy display name must be unique within an organization.
	   It must: contain only unicode letters, numbers, underscores, dashes
	   and spaces; not start or end with spaces; and be at most 200 bytes
	   long when encoded in UTF-8.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Taxonomy location region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
