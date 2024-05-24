package auditmanager

import types "DesignSphere_Server/src/resource/aws/types"

type Assessment struct {
	// Unique identifier of the framework the assessment will be created from.
	FrameworkId string `json:"frameworkId,omitempty" yaml:"frameworkId,omitempty"`

	// Name of the assessment.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of roles for the assessment. See `roles` below.
	Roles []types.Auditmanager_AssessmentRole `json:"roles,omitempty" yaml:"roles,omitempty"`

	/*
	   Amazon Web Services accounts and services that are in scope for the assessment. See `scope` below.

	   The following arguments are optional:
	*/
	Scope types.Auditmanager_AssessmentScope `json:"scope,omitempty" yaml:"scope,omitempty"`

	// A map of tags to assign to the assessment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Assessment report storage destination configuration. See `assessment_reports_destination` below.
	AssessmentReportsDestination types.Auditmanager_AssessmentAssessmentReportsDestination `json:"assessmentReportsDestination,omitempty" yaml:"assessmentReportsDestination,omitempty"`

	// Description of the assessment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
