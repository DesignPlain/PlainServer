package detective

type OrganizationConfiguration struct {
	// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s Detective delegated administrator and Detective is enabled in that AWS Region.
	AutoEnable bool `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// ARN of the behavior graph.
	GraphArn string `json:"graphArn,omitempty" yaml:"graphArn,omitempty"`
}
