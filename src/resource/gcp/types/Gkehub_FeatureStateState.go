package types

type Gkehub_FeatureStateState struct {
	/*
	   (Output)
	   The high-level, machine-readable status of this Feature.
	*/
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	/*
	   (Output)
	   A human-readable description of the current status.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   (Output)
	   The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
