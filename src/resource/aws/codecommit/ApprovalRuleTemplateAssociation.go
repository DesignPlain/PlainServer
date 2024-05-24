package codecommit

type ApprovalRuleTemplateAssociation struct {
	// The name for the approval rule template.
	ApprovalRuleTemplateName string `json:"approvalRuleTemplateName,omitempty" yaml:"approvalRuleTemplateName,omitempty"`

	// The name of the repository that you want to associate with the template.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
