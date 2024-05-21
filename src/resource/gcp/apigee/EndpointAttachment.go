package apigee

type EndpointAttachment struct {
	/*
	   ID of the endpoint attachment.


	   - - -
	*/
	EndpointAttachmentId string `json:"endpointAttachmentId,omitempty" yaml:"endpointAttachmentId,omitempty"`

	// Location of the endpoint attachment.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The Apigee Organization associated with the Apigee instance,
	   in the format `organizations/{{org_name}}`.
	*/
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Format: projects/-/regions/-/serviceAttachments/-
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`
}
