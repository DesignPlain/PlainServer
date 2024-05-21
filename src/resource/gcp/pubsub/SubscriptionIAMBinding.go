package pubsub

import types "DesignSphere_Server/src/resource/gcp/types"

type SubscriptionIAMBinding struct {
	/*
	   The subscription name or id to bind to attach IAM policy to.

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	*/
	Subscription string `json:"subscription,omitempty" yaml:"subscription,omitempty"`

	//
	Condition types.Pubsub_SubscriptionIAMBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
