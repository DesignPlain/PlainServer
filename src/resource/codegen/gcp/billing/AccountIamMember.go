package billing

import types "libds/gcp/types"

type AccountIamMember struct {
	/*
	   The billing account id.

	   For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	*/
	BillingAccountId string `json:"billingAccountId,omitempty" yaml:"billingAccountId,omitempty"`

	//
	Condition types.Billing_AccountIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

	   `gcp.billing.AccountIamPolicy` only:
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
