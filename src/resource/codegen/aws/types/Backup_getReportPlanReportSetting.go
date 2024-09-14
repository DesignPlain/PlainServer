package types

type Backup_getReportPlanReportSetting struct {
	// Identifies the report template for the report. Reports are built using a report template.
	ReportTemplate string `json:"reportTemplate,omitempty" yaml:"reportTemplate,omitempty"`

	// (Optional) Specifies the list of accounts a report covers.
	Accounts []string `json:"accounts,omitempty" yaml:"accounts,omitempty"`

	// ARNs of the frameworks a report covers.
	FrameworkArns []string `json:"frameworkArns,omitempty" yaml:"frameworkArns,omitempty"`

	// Specifies the number of frameworks a report covers.
	NumberOfFrameworks int `json:"numberOfFrameworks,omitempty" yaml:"numberOfFrameworks,omitempty"`

	// (Optional) Specifies the list of Organizational Units a report covers.
	OrganizationUnits []string `json:"organizationUnits,omitempty" yaml:"organizationUnits,omitempty"`

	// (Optional) Specifies the list of regions a report covers.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`
}
