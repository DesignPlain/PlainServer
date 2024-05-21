package folder

import types "DesignSphere_Server/src/resource/gcp/types"

type IAMBinding struct {
	//
	Condition types.Folder_IAMBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	/*
	   An array of identities that will be granted the privilege in the `role`.
	   Each entry can have one of the following values:
	   - --user:{emailid}--: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	   - For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	*/
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
