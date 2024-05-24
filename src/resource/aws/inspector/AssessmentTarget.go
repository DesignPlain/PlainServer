package inspector

type AssessmentTarget struct {
	// The name of the assessment target.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
	ResourceGroupArn string `json:"resourceGroupArn,omitempty" yaml:"resourceGroupArn,omitempty"`
}
