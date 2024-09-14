package auditmanager

type AssessmentReport struct {
	/*
	   Unique identifier of the assessment to create the report from.

	   The following arguments are optional:
	*/
	AssessmentId string `json:"assessmentId,omitempty" yaml:"assessmentId,omitempty"`

	// Description of the assessment report.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the assessment report.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
