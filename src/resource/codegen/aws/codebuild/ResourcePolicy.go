package codebuild

type ResourcePolicy struct {
	// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
