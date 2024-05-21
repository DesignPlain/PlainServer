package healthcare

type Dataset struct {
	/*
	   The location for the Dataset.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The resource name for the Dataset.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	   "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	   (e.g., HL7 messages) where no explicit timezone is specified.
	*/
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}
