package types

type Projects_ApiKeyRestrictionsApiTarget struct {
	// Optional. List of one or more methods that can be called. If empty, all methods for the service are allowed. A wildcard (-) can be used as the last symbol. Valid examples: `google.cloud.translate.v2.TranslateService.GetSupportedLanguage` `TranslateText` `Get-` `translate.googleapis.com.Get-`
	Methods []string `json:"methods,omitempty" yaml:"methods,omitempty"`

	// The service for this restriction. It should be the canonical service name, for example: `translate.googleapis.com`. You can use `gcloud services list` to get a list of services that are enabled in the project.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
