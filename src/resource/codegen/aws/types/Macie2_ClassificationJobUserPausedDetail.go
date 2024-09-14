package types

type Macie2_ClassificationJobUserPausedDetail struct {
	//
	JobExpiresAt string `json:"jobExpiresAt,omitempty" yaml:"jobExpiresAt,omitempty"`

	//
	JobImminentExpirationHealthEventArn string `json:"jobImminentExpirationHealthEventArn,omitempty" yaml:"jobImminentExpirationHealthEventArn,omitempty"`

	//
	JobPausedAt string `json:"jobPausedAt,omitempty" yaml:"jobPausedAt,omitempty"`
}
