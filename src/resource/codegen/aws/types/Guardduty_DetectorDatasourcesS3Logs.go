package types

type Guardduty_DetectorDatasourcesS3Logs struct {
	/*
	   If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	   Defaults to `true`.
	*/
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
}
