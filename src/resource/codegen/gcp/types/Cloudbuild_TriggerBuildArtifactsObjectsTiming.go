package types

type Cloudbuild_TriggerBuildArtifactsObjectsTiming struct {
	/*
	   End of time span.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	   nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	/*
	   Start of time span.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	   nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
