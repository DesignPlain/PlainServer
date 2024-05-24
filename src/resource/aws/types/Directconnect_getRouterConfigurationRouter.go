package types

type Directconnect_getRouterConfigurationRouter struct {
	// Router platform
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`

	/*
	   ID of the Router Type. For example: `CiscoSystemsInc-2900SeriesRouters-IOS124`

	   There is currently no AWS API to retrieve the full list of `router_type_identifier` values. Here is a list of known `RouterType` objects that can be used:
	*/
	RouterTypeIdentifier string `json:"routerTypeIdentifier,omitempty" yaml:"routerTypeIdentifier,omitempty"`

	// Router operating system
	Software string `json:"software,omitempty" yaml:"software,omitempty"`

	// Router vendor
	Vendor string `json:"vendor,omitempty" yaml:"vendor,omitempty"`

	// Router XSLT Template Name
	XsltTemplateName string `json:"xsltTemplateName,omitempty" yaml:"xsltTemplateName,omitempty"`

	//
	XsltTemplateNameForMacSec string `json:"xsltTemplateNameForMacSec,omitempty" yaml:"xsltTemplateNameForMacSec,omitempty"`
}
