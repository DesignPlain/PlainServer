package types

type Dataplex_DatascanExecutionStatus struct {
	/*
	   (Output)
	   The time when the latest DataScanJob started.
	*/
	LatestJobEndTime string `json:"latestJobEndTime,omitempty" yaml:"latestJobEndTime,omitempty"`

	/*
	   (Output)
	   The time when the latest DataScanJob ended.
	*/
	LatestJobStartTime string `json:"latestJobStartTime,omitempty" yaml:"latestJobStartTime,omitempty"`
}
