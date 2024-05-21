package iap

type Client struct {
	/*
	   Identifier of the brand to which this client
	   is attached to. The format is
	   `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.


	   - - -
	*/
	Brand string `json:"brand,omitempty" yaml:"brand,omitempty"`

	// Human-friendly name given to the OAuth client.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
