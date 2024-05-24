package guardduty

type InviteAccepter struct {
	// The detector ID of the member GuardDuty account.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// AWS account ID for primary account.
	MasterAccountId string `json:"masterAccountId,omitempty" yaml:"masterAccountId,omitempty"`
}
