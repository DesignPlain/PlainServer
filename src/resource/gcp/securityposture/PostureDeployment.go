package securityposture

type PostureDeployment struct {
	/*
	   ID of the posture deployment.


	   - - -
	*/
	PostureDeploymentId string `json:"postureDeploymentId,omitempty" yaml:"postureDeploymentId,omitempty"`

	/*
	   Relative name of the posture which needs to be deployed. It should be in the format:
	   organizations/{organization_id}/locations/{location}/postures/{posture_id}
	*/
	PostureId string `json:"postureId,omitempty" yaml:"postureId,omitempty"`

	// Revision_id the posture which needs to be deployed.
	PostureRevisionId string `json:"postureRevisionId,omitempty" yaml:"postureRevisionId,omitempty"`

	/*
	   The resource on which the posture should be deployed. This can be in one of the following formats:
	   projects/{project_number},
	   folders/{folder_number},
	   organizations/{organization_id}
	*/
	TargetResource string `json:"targetResource,omitempty" yaml:"targetResource,omitempty"`

	// Description of the posture deployment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The location of the resource, eg. global`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
