package types

type Osconfig_PatchDeploymentRecurringSchedule struct {
	/*
	   (Output)
	   The time the next patch job is scheduled to run.
	   A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	*/
	NextExecuteTime string `json:"nextExecuteTime,omitempty" yaml:"nextExecuteTime,omitempty"`

	/*
	   The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.
	   A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	/*
	   Time of the day to run a recurring deployment.
	   Structure is documented below.
	*/
	TimeOfDay Osconfig_PatchDeploymentRecurringScheduleTimeOfDay `json:"timeOfDay,omitempty" yaml:"timeOfDay,omitempty"`

	/*
	   Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are
	   determined by the chosen time zone.
	   Structure is documented below.
	*/
	TimeZone Osconfig_PatchDeploymentRecurringScheduleTimeZone `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	/*
	   Schedule with weekly executions.
	   Structure is documented below.
	*/
	Weekly Osconfig_PatchDeploymentRecurringScheduleWeekly `json:"weekly,omitempty" yaml:"weekly,omitempty"`

	/*
	   The end time at which a recurring patch deployment schedule is no longer active.
	   A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	/*
	   (Output)
	   The time the last patch job ran successfully.
	   A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	*/
	LastExecuteTime string `json:"lastExecuteTime,omitempty" yaml:"lastExecuteTime,omitempty"`

	/*
	   Schedule with monthly executions.
	   Structure is documented below.
	*/
	Monthly Osconfig_PatchDeploymentRecurringScheduleMonthly `json:"monthly,omitempty" yaml:"monthly,omitempty"`
}
