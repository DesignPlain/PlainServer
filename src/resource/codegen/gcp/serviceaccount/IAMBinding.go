package serviceaccount

import types "libds/gcp/types"

type IAMBinding struct {
	/*
	   An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	   Structure is documented below.
	*/
	Condition types.Serviceaccount_IAMBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	/*
	   The fully-qualified name of the service account to apply policy to.

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	*/
	ServiceAccountId string `json:"serviceAccountId,omitempty" yaml:"serviceAccountId,omitempty"`
}
