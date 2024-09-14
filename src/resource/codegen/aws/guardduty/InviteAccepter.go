package guardduty

type InviteAccepter struct {
	// AWS account ID for primary account.
	MasterAccountId string `json:"masterAccountId,omitempty" yaml:"masterAccountId,omitempty"`

	// The detector ID of the member GuardDuty account.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`
}
