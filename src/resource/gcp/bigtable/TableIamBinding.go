package bigtable

import types "DesignSphere_Server/src/resource/gcp/types"

type TableIamBinding struct {
	//
	Condition types.Bigtable_TableIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// The name or relative resource id of the instance that owns the table.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The project in which the table belongs. If it
	   is not provided, this provider will use the provider default.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

	   `gcp.bigtable.TableIamPolicy` only:
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	/*
	   The name or relative resource id of the table to manage IAM policies for.

	   For `gcp.bigtable.TableIamMember` or `gcp.bigtable.TableIamBinding`:

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	*/
	Table string `json:"table,omitempty" yaml:"table,omitempty"`
}
