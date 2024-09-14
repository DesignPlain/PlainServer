package types

type Costexplorer_getTagsTimePeriod struct {
	// Beginning of the time period.
	End string `json:"end,omitempty" yaml:"end,omitempty"`

	// End of the time period.
	Start string `json:"start,omitempty" yaml:"start,omitempty"`
}
