package inspector2

type DelegatedAdminAccount struct {
	// Account to enable as delegated admin account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
