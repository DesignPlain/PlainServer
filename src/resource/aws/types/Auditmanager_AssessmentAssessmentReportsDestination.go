package types

type Auditmanager_AssessmentAssessmentReportsDestination struct {
	// Destination of the assessment report. This value be in the form `s3://{bucket_name}`.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Destination type. Currently, `S3` is the only valid value.
	DestinationType string `json:"destinationType,omitempty" yaml:"destinationType,omitempty"`
}
