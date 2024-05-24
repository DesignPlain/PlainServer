package types

type Securityhub_AutomationRuleCriteria struct {
	// The level of importance that is assigned to the resources that are associated with a finding. Documented below.
	Criticalities []Securityhub_AutomationRuleCriteriaCriticality `json:"criticalities,omitempty" yaml:"criticalities,omitempty"`

	// Provides the current state of a finding. Documented below.
	RecordStates []Securityhub_AutomationRuleCriteriaRecordState `json:"recordStates,omitempty" yaml:"recordStates,omitempty"`

	// One or more finding types in the format of namespace/category/classifier that classify a finding. Documented below.
	Types []Securityhub_AutomationRuleCriteriaType `json:"types,omitempty" yaml:"types,omitempty"`

	// A timestamp that indicates when the finding record was most recently updated. Documented below.
	UpdatedAts []Securityhub_AutomationRuleCriteriaUpdatedAt `json:"updatedAts,omitempty" yaml:"updatedAts,omitempty"`

	// Provides information about the status of the investigation into a finding. Documented below.
	WorkflowStatuses []Securityhub_AutomationRuleCriteriaWorkflowStatus `json:"workflowStatuses,omitempty" yaml:"workflowStatuses,omitempty"`

	// A finding's description. Documented below.
	Descriptions []Securityhub_AutomationRuleCriteriaDescription `json:"descriptions,omitempty" yaml:"descriptions,omitempty"`

	// A timestamp that indicates when the potential security issue captured by a finding was first observed by the security findings product. Documented below.
	FirstObservedAts []Securityhub_AutomationRuleCriteriaFirstObservedAt `json:"firstObservedAts,omitempty" yaml:"firstObservedAts,omitempty"`

	// The identifier for the solution-specific component that generated a finding. Documented below.
	GeneratorIds []Securityhub_AutomationRuleCriteriaGeneratorId `json:"generatorIds,omitempty" yaml:"generatorIds,omitempty"`

	// Custom fields and values about the resource that a finding pertains to. Documented below.
	ResourceDetailsOthers []Securityhub_AutomationRuleCriteriaResourceDetailsOther `json:"resourceDetailsOthers,omitempty" yaml:"resourceDetailsOthers,omitempty"`

	// Provides a URL that links to a page about the current finding in the finding product. Documented below.
	SourceUrls []Securityhub_AutomationRuleCriteriaSourceUrl `json:"sourceUrls,omitempty" yaml:"sourceUrls,omitempty"`

	// The type of resource that the finding pertains to. Documented below.
	ResourceTypes []Securityhub_AutomationRuleCriteriaResourceType `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	// A list of user-defined name and value string pairs added to a finding. Documented below.
	UserDefinedFields []Securityhub_AutomationRuleCriteriaUserDefinedField `json:"userDefinedFields,omitempty" yaml:"userDefinedFields,omitempty"`

	// Provides the veracity of a finding. Documented below.
	VerificationStates []Securityhub_AutomationRuleCriteriaVerificationState `json:"verificationStates,omitempty" yaml:"verificationStates,omitempty"`

	// The unique identifier of a standard in which a control is enabled. Documented below.
	ComplianceAssociatedStandardsIds []Securityhub_AutomationRuleCriteriaComplianceAssociatedStandardsId `json:"complianceAssociatedStandardsIds,omitempty" yaml:"complianceAssociatedStandardsIds,omitempty"`

	// The security control ID for which a finding was generated. Security control IDs are the same across standards. Documented below.
	ComplianceSecurityControlIds []Securityhub_AutomationRuleCriteriaComplianceSecurityControlId `json:"complianceSecurityControlIds,omitempty" yaml:"complianceSecurityControlIds,omitempty"`

	// The product-specific identifier for a finding. Documented below.
	Ids []Securityhub_AutomationRuleCriteriaId `json:"ids,omitempty" yaml:"ids,omitempty"`

	// The timestamp of when the note was updated. Documented below.
	NoteUpdatedAts []Securityhub_AutomationRuleCriteriaNoteUpdatedAt `json:"noteUpdatedAts,omitempty" yaml:"noteUpdatedAts,omitempty"`

	// The identifier for the given resource type. For AWS resources that are identified by Amazon Resource Names (ARNs), this is the ARN. For AWS resources that lack ARNs, this is the identifier as defined by the AWS service that created the resource. For non-AWS resources, this is a unique identifier that is associated with the resource. Documented below.
	ResourceIds []Securityhub_AutomationRuleCriteriaResourceId `json:"resourceIds,omitempty" yaml:"resourceIds,omitempty"`

	// The severity value of the finding. Documented below.
	SeverityLabels []Securityhub_AutomationRuleCriteriaSeverityLabel `json:"severityLabels,omitempty" yaml:"severityLabels,omitempty"`

	// The AWS account ID in which a finding was generated. Documented below.
	AwsAccountIds []Securityhub_AutomationRuleCriteriaAwsAccountId `json:"awsAccountIds,omitempty" yaml:"awsAccountIds,omitempty"`

	// The name of the AWS account in which a finding was generated. Documented below.
	AwsAccountNames []Securityhub_AutomationRuleCriteriaAwsAccountName `json:"awsAccountNames,omitempty" yaml:"awsAccountNames,omitempty"`

	// The result of a security check. This field is only used for findings generated from controls. Documented below.
	ComplianceStatuses []Securityhub_AutomationRuleCriteriaComplianceStatus `json:"complianceStatuses,omitempty" yaml:"complianceStatuses,omitempty"`

	// A timestamp that indicates when this finding record was created. Documented below.
	CreatedAts []Securityhub_AutomationRuleCriteriaCreatedAt `json:"createdAts,omitempty" yaml:"createdAts,omitempty"`

	// Provides the name of the product that generated the finding. For control-based findings, the product name is Security Hub. Documented below.
	ProductNames []Securityhub_AutomationRuleCriteriaProductName `json:"productNames,omitempty" yaml:"productNames,omitempty"`

	// The product-generated identifier for a related finding.  Documented below.
	RelatedFindingsIds []Securityhub_AutomationRuleCriteriaRelatedFindingsId `json:"relatedFindingsIds,omitempty" yaml:"relatedFindingsIds,omitempty"`

	// The partition in which the resource that the finding pertains to is located. A partition is a group of AWS Regions. Each AWS account is scoped to one partition. Documented below.
	ResourcePartitions []Securityhub_AutomationRuleCriteriaResourcePartition `json:"resourcePartitions,omitempty" yaml:"resourcePartitions,omitempty"`

	// The likelihood that a finding accurately identifies the behavior or issue that it was intended to identify. `Confidence` is scored on a 0–100 basis using a ratio scale. A value of `0` means 0 percent confidence, and a value of `100` means 100 percent confidence. Documented below.
	Confidences []Securityhub_AutomationRuleCriteriaConfidence `json:"confidences,omitempty" yaml:"confidences,omitempty"`

	// The principal that created a note. Documented below.
	NoteUpdatedBies []Securityhub_AutomationRuleCriteriaNoteUpdatedBy `json:"noteUpdatedBies,omitempty" yaml:"noteUpdatedBies,omitempty"`

	// The AWS Region where the resource that a finding pertains to is located. Documented below.
	ResourceRegions []Securityhub_AutomationRuleCriteriaResourceRegion `json:"resourceRegions,omitempty" yaml:"resourceRegions,omitempty"`

	// The name of the company for the product that generated the finding. For control-based findings, the company is AWS. Documented below.
	CompanyNames []Securityhub_AutomationRuleCriteriaCompanyName `json:"companyNames,omitempty" yaml:"companyNames,omitempty"`

	// A timestamp that indicates when the potential security issue captured by a finding was most recently observed by the security findings product. Documented below.
	LastObservedAts []Securityhub_AutomationRuleCriteriaLastObservedAt `json:"lastObservedAts,omitempty" yaml:"lastObservedAts,omitempty"`

	// The ARN for the product that generated a related finding. Documented below.
	RelatedFindingsProductArns []Securityhub_AutomationRuleCriteriaRelatedFindingsProductArn `json:"relatedFindingsProductArns,omitempty" yaml:"relatedFindingsProductArns,omitempty"`

	// The Amazon Resource Name (ARN) of the application that is related to a finding. Documented below.
	ResourceApplicationArns []Securityhub_AutomationRuleCriteriaResourceApplicationArn `json:"resourceApplicationArns,omitempty" yaml:"resourceApplicationArns,omitempty"`

	// The text of a user-defined note that's added to a finding. Documented below.
	NoteTexts []Securityhub_AutomationRuleCriteriaNoteText `json:"noteTexts,omitempty" yaml:"noteTexts,omitempty"`

	// The Amazon Resource Name (ARN) for a third-party product that generated a finding in Security Hub. Documented below.
	ProductArns []Securityhub_AutomationRuleCriteriaProductArn `json:"productArns,omitempty" yaml:"productArns,omitempty"`

	// The name of the application that is related to a finding. Documented below.
	ResourceApplicationNames []Securityhub_AutomationRuleCriteriaResourceApplicationName `json:"resourceApplicationNames,omitempty" yaml:"resourceApplicationNames,omitempty"`

	// A list of AWS tags associated with a resource at the time the finding was processed. Documented below.
	ResourceTags []Securityhub_AutomationRuleCriteriaResourceTag `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// A finding's title. Documented below.
	Titles []Securityhub_AutomationRuleCriteriaTitle `json:"titles,omitempty" yaml:"titles,omitempty"`
}
