package networksecurity

type UrlList struct {
	/*
	   Short name of the UrlList resource to be created.
	   This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. 'urlList'.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// FQDNs and URLs.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The location of the url lists.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
