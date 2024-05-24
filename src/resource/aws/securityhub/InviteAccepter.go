package securityhub

type InviteAccepter struct {
	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterId string `json:"masterId,omitempty" yaml:"masterId,omitempty"`
}
