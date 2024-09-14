package quicksight

type FolderMembership struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Identifier for the folder.
	FolderId string `json:"folderId,omitempty" yaml:"folderId,omitempty"`

	// ID of the asset (the dashboard, analysis, or dataset).
	MemberId string `json:"memberId,omitempty" yaml:"memberId,omitempty"`

	/*
	   Type of the member. Valid values are `ANALYSIS`, `DASHBOARD`, and `DATASET`.

	   The following arguments are optional:
	*/
	MemberType string `json:"memberType,omitempty" yaml:"memberType,omitempty"`
}
