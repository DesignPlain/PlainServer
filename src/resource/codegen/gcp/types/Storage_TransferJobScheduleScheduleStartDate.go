package types

type Storage_TransferJobScheduleScheduleStartDate struct {
	/*
	   Day of month. Must be from 1 to 31 and valid for the year and month.

	   <a name="nested_start_time_of_day"></a>The `start_time_of_day` blocks support:
	*/
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// Month of year. Must be from 1 to 12.
	Month int `json:"month,omitempty" yaml:"month,omitempty"`

	// Year of date. Must be from 1 to 9999.
	Year int `json:"year,omitempty" yaml:"year,omitempty"`
}
