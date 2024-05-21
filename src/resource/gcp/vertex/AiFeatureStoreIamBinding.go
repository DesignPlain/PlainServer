package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiFeatureStoreIamBinding struct {
	//
	Condition types.Vertex_AiFeatureStoreIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// Used to find the parent resource to bind the IAM policy to
	Featurestore string `json:"featurestore,omitempty" yaml:"featurestore,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	   - --projectOwner:projectid--: Owners of the given project. For example, "projectOwner:my-example-project"
	   - --projectEditor:projectid--: Editors of the given project. For example, "projectEditor:my-example-project"
	   - --projectViewer:projectid--: Viewers of the given project. For example, "projectViewer:my-example-project"
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
	   the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	   region is specified, it is taken from the provider configuration.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
