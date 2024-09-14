package logging

type LogView struct {
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The resource name of the view. For example: \`projects/my-project/locations/global/buckets/my-bucket/views/my-view\`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parent of the resource.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The bucket of the resource


	   - - -
	*/
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Describes this view.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
