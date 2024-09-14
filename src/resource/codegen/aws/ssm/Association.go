package ssm

import types "libds/aws/types"

type Association struct {
	// Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls. This should be set to the SSM document `parameter` that will define how your automation will branch out.
	AutomationTargetParameterName string `json:"automationTargetParameterName,omitempty" yaml:"automationTargetParameterName,omitempty"`

	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%!.(MISSING)
	MaxConcurrency string `json:"maxConcurrency,omitempty" yaml:"maxConcurrency,omitempty"`

	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets []types.Ssm_AssociationTarget `json:"targets,omitempty" yaml:"targets,omitempty"`

	/*
	   The number of seconds to wait for the association status to be `Success`. If `Success` status is not reached within the given time, create opration will fail.

	   Output Location (`output_location`) is an S3 bucket where you want to store the results of this association:
	*/
	WaitForSuccessTimeoutSeconds int `json:"waitForSuccessTimeoutSeconds,omitempty" yaml:"waitForSuccessTimeoutSeconds,omitempty"`

	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion string `json:"documentVersion,omitempty" yaml:"documentVersion,omitempty"`

	// The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above. Use the `targets` attribute instead.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The name of the SSM document to apply.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html) that specifies when the association runs.
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`

	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%! (MISSING)If you specify a threshold of 3, the stop command is sent when the fourth error is returned. If you specify a threshold of 10%!f(MISSING)or 50 associations, the stop command is sent when the sixth error is returned.
	MaxErrors string `json:"maxErrors,omitempty" yaml:"maxErrors,omitempty"`

	// The mode for generating association compliance. You can specify `AUTO` or `MANUAL`.
	SyncCompliance string `json:"syncCompliance,omitempty" yaml:"syncCompliance,omitempty"`

	// By default, when you create a new or update associations, the system runs it immediately and then according to the schedule you specified. Enable this option if you do not want an association to run immediately after you create or update it. This parameter is not supported for rate expressions. Default: `false`.
	ApplyOnlyAtCronInterval bool `json:"applyOnlyAtCronInterval,omitempty" yaml:"applyOnlyAtCronInterval,omitempty"`

	// The descriptive name for the association.
	AssociationName string `json:"associationName,omitempty" yaml:"associationName,omitempty"`

	// The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
	ComplianceSeverity string `json:"complianceSeverity,omitempty" yaml:"complianceSeverity,omitempty"`

	// An output location block. Output Location is documented below.
	OutputLocation types.Ssm_AssociationOutputLocation `json:"outputLocation,omitempty" yaml:"outputLocation,omitempty"`

	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// A map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
