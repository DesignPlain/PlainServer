package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationBucketingConfigBucketReplacementValueTimeValue struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Hours int `json:"hours,omitempty" yaml:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos int `json:"nanos,omitempty" yaml:"nanos,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	Seconds int `json:"seconds,omitempty" yaml:"seconds,omitempty"`
}
