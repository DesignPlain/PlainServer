package types

type Bigquery_ConnectionAzure struct {
	/*
	   (Output)
	   The name of the Azure Active Directory Application.
	*/
	Application string `json:"application,omitempty" yaml:"application,omitempty"`

	/*
	   (Output)
	   The client id of the Azure Active Directory Application.
	*/
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The id of customer's directory that host the data.
	CustomerTenantId string `json:"customerTenantId,omitempty" yaml:"customerTenantId,omitempty"`

	// The Azure Application (client) ID where the federated credentials will be hosted.
	FederatedApplicationClientId string `json:"federatedApplicationClientId,omitempty" yaml:"federatedApplicationClientId,omitempty"`

	/*
	   (Output)
	   A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's Azure Active Directory Application.
	*/
	Identity string `json:"identity,omitempty" yaml:"identity,omitempty"`

	/*
	   (Output)
	   The object id of the Azure Active Directory Application.
	*/
	ObjectId string `json:"objectId,omitempty" yaml:"objectId,omitempty"`

	/*
	   (Output)
	   The URL user will be redirected to after granting consent during connection setup.
	*/
	RedirectUri string `json:"redirectUri,omitempty" yaml:"redirectUri,omitempty"`
}
