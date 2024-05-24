package types

type Ssoadmin_ApplicationPortalOptionsSignInOptions struct {
	// URL that accepts authentication requests for an application.
	ApplicationUrl string `json:"applicationUrl,omitempty" yaml:"applicationUrl,omitempty"`

	/*
	   Determines how IAM Identity Center navigates the user to the target application.
	   Valid values are `APPLICATION` and `IDENTITY_CENTER`.
	   If `APPLICATION` is set, IAM Identity Center redirects the customer to the configured `application_url`.
	   If `IDENTITY_CENTER` is set, IAM Identity Center uses SAML identity-provider initiated authentication to sign the customer directly into a SAML-based application.
	*/
	Origin string `json:"origin,omitempty" yaml:"origin,omitempty"`
}
