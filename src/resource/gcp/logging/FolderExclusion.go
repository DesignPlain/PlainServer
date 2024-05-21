package logging

type FolderExclusion struct {
	// A human-readable description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether this exclusion rule should be disabled or not. This defaults to
	   false.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	   See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	   write a filter.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	   accepted.
	*/
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	// The name of the logging exclusion.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
