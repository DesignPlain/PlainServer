package discoveryengine

type DataStore struct {
	/*
	   The content config of the data store.
	   Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
	*/
	ContentConfig string `json:"contentConfig,omitempty" yaml:"contentConfig,omitempty"`

	/*
	   If true, an advanced data store for site search will be created. If the
	   data store is not configured as site search (GENERIC vertical and
	   PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	*/
	CreateAdvancedSiteSearch bool `json:"createAdvancedSiteSearch,omitempty" yaml:"createAdvancedSiteSearch,omitempty"`

	/*
	   The unique id of the data store.


	   - - -
	*/
	DataStoreId string `json:"dataStoreId,omitempty" yaml:"dataStoreId,omitempty"`

	/*
	   The display name of the data store. This field must be a UTF-8 encoded
	   string with a length limit of 128 characters.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The industry vertical that the data store registers.
	   Possible values are: `GENERIC`, `MEDIA`.
	*/
	IndustryVertical string `json:"industryVertical,omitempty" yaml:"industryVertical,omitempty"`

	/*
	   The geographic location where the data store should reside. The value can
	   only be one of "global", "us" and "eu".
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The solutions that the data store enrolls.
	   Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
	*/
	SolutionTypes []string `json:"solutionTypes,omitempty" yaml:"solutionTypes,omitempty"`
}
