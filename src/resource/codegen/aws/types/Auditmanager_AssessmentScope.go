package types

type Auditmanager_AssessmentScope struct {
	// Amazon Web Services accounts that are in scope for the assessment. See `aws_accounts` below.
	AwsAccounts []Auditmanager_AssessmentScopeAwsAccount `json:"awsAccounts,omitempty" yaml:"awsAccounts,omitempty"`

	// Amazon Web Services services that are included in the scope of the assessment. See `aws_services` below.
	AwsServices []Auditmanager_AssessmentScopeAwsService `json:"awsServices,omitempty" yaml:"awsServices,omitempty"`
}
