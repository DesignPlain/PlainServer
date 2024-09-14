package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionConditionExpressionsConditionsConditionValueDateValue struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year int `json:"year,omitempty" yaml:"year,omitempty"`

	/*
	   Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.

	   - - -
	*/
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month int `json:"month,omitempty" yaml:"month,omitempty"`
}
