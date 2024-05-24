package types

type Datasync_TaskExcludes struct {
	// The type of filter rule to apply. Valid values: `SIMPLE_PATTERN`.
	FilterType string `json:"filterType,omitempty" yaml:"filterType,omitempty"`

	// A single filter string that consists of the patterns to exclude. The patterns are delimited by "|" (that is, a pipe), for example: `/folder1|/folder2`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
