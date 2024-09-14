package billing

type ProjectInfo struct {
	/*
	   The ID of the billing account associated with the project, if
	   any. Set to empty string to disable billing for the project.
	   For example, `"012345-567890-ABCDEF"` or `""`.


	   - - -
	*/
	BillingAccount string `json:"billingAccount,omitempty" yaml:"billingAccount,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
