package folder

import types "DesignSphere_Server/src/resource/gcp/types"

type IAMMember struct {
	/*
	   An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	   Structure is documented below.
	*/
	Condition types.Folder_IAMMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	   `organizations/{{org_id}}/roles/{{role_id}}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
