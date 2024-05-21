package types

type Healthcare_Hl7StoreParserConfig struct {
	/*
	   Byte(s) to be used as the segment terminator. If this is unset, '\r' will be used as segment terminator.
	   A base64-encoded string.
	*/
	SegmentTerminator string `json:"segmentTerminator,omitempty" yaml:"segmentTerminator,omitempty"`

	/*
	   The version of the unschematized parser to be used when a custom `schema` is not set.
	   Default value is `V1`.
	   Possible values are: `V1`, `V2`, `V3`.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Determines whether messages with no header are allowed.
	AllowNullHeader bool `json:"allowNullHeader,omitempty" yaml:"allowNullHeader,omitempty"`

	/*
	   JSON encoded string for schemas used to parse messages in this
	   store if schematized parsing is desired.
	*/
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
