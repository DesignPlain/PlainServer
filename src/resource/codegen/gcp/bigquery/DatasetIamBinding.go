package bigquery

import types "libds/gcp/types"

type DatasetIamBinding struct {
	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Bigquery_DatasetIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	/*
	   The dataset ID.

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --iamMember:{principal}--: Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. This is used for example for workload/workforce federated identities (principal, principalSet).
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	*/
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`
}
