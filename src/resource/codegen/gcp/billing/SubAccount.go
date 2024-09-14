package billing

type SubAccount struct {
	/*
	   If set to "RENAME_ON_DESTROY" the billing account display_name
	   will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
	   Default is "".
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// The display name of the billing account.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The name of the master billing account that the subaccount
	   will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	*/
	MasterBillingAccount string `json:"masterBillingAccount,omitempty" yaml:"masterBillingAccount,omitempty"`
}
