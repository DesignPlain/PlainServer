package bigtable

type InstanceIamPolicy struct {
	/*
	   The project in which the instance belongs. If it
	   is not provided, a default will be supplied.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The name or relative resource id of the instance to manage IAM policies for.

	   For `gcp.bigtable.InstanceIamMember` or `gcp.bigtable.InstanceIamBinding`:

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   The policy data generated by a `gcp.organizations.getIAMPolicy` data source.

	   - - -
	*/
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`
}
