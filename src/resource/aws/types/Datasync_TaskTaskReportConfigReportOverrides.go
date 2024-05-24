package types

type Datasync_TaskTaskReportConfigReportOverrides struct {
	// Specifies the level of reporting for the files, objects, and directories that DataSync attempted to delete in your destination location. This only applies if you configure your task to delete data in the destination that isn't in the source. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
	DeletedOverride string `json:"deletedOverride,omitempty" yaml:"deletedOverride,omitempty"`

	// Specifies the level of reporting for the files, objects, and directories that DataSync attempted to skip during your transfer. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
	SkippedOverride string `json:"skippedOverride,omitempty" yaml:"skippedOverride,omitempty"`

	// Specifies the level of reporting for the files, objects, and directories that DataSync attempted to transfer. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
	TransferredOverride string `json:"transferredOverride,omitempty" yaml:"transferredOverride,omitempty"`

	/*
	   Specifies the level of reporting for the files, objects, and directories that DataSync attempted to verify at the end of your transfer. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.

	   > --NOTE:-- If any `report_overrides` are set to the same value as `task_report_config.report_level`, they will always be flagged as changed. Only set overrides to a value that differs from `task_report_config.report_level`.
	*/
	VerifiedOverride string `json:"verifiedOverride,omitempty" yaml:"verifiedOverride,omitempty"`
}
