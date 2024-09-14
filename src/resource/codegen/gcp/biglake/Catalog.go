package biglake

type Catalog struct {
	// The geographic location where the Catalog should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The name of the Catalog. Format:
	   projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
