package types

type Backup_FrameworkControlScope struct {
	// The ID of the only AWS resource that you want your control scope to contain. Minimum number of 1 item. Maximum number of 100 items.
	ComplianceResourceIds []string `json:"complianceResourceIds,omitempty" yaml:"complianceResourceIds,omitempty"`

	// Describes whether the control scope includes one or more types of resources, such as EFS or RDS.
	ComplianceResourceTypes []string `json:"complianceResourceTypes,omitempty" yaml:"complianceResourceTypes,omitempty"`

	// The tag key-value pair applied to those AWS resources that you want to trigger an evaluation for a rule. A maximum of one key-value pair can be provided.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
