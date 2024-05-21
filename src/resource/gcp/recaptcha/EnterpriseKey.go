package recaptcha

import types "DesignSphere_Server/src/resource/gcp/types"

type EnterpriseKey struct {
	// Settings for keys that can be used by iOS apps.
	IosSettings types.Recaptcha_EnterpriseKeyIosSettings `json:"iosSettings,omitempty" yaml:"iosSettings,omitempty"`

	/*
	   See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Options for user acceptance testing.
	TestingOptions types.Recaptcha_EnterpriseKeyTestingOptions `json:"testingOptions,omitempty" yaml:"testingOptions,omitempty"`

	// Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings types.Recaptcha_EnterpriseKeyWafSettings `json:"wafSettings,omitempty" yaml:"wafSettings,omitempty"`

	// Settings for keys that can be used by websites.
	WebSettings types.Recaptcha_EnterpriseKeyWebSettings `json:"webSettings,omitempty" yaml:"webSettings,omitempty"`

	// Settings for keys that can be used by Android apps.
	AndroidSettings types.Recaptcha_EnterpriseKeyAndroidSettings `json:"androidSettings,omitempty" yaml:"androidSettings,omitempty"`

	/*
	   Human-readable display name of this key. Modifiable by user.



	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
