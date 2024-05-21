package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationReplaceConfigNewValue struct {
	/*
	   Represents a whole or partial calendar date.
	   Structure is documented below.
	*/
	DateValue Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationReplaceConfigNewValueDateValue `json:"dateValue,omitempty" yaml:"dateValue,omitempty"`

	/*
	   Represents a day of the week.
	   Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	DayOfWeekValue string `json:"dayOfWeekValue,omitempty" yaml:"dayOfWeekValue,omitempty"`

	// A float value.
	FloatValue float64 `json:"floatValue,omitempty" yaml:"floatValue,omitempty"`

	// An integer value (int64 format)
	IntegerValue string `json:"integerValue,omitempty" yaml:"integerValue,omitempty"`

	// A string value.
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`

	/*
	   Represents a time of day.
	   Structure is documented below.
	*/
	TimeValue Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationReplaceConfigNewValueTimeValue `json:"timeValue,omitempty" yaml:"timeValue,omitempty"`

	/*
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	   Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	TimestampValue string `json:"timestampValue,omitempty" yaml:"timestampValue,omitempty"`

	// A boolean value.
	BooleanValue bool `json:"booleanValue,omitempty" yaml:"booleanValue,omitempty"`
}
