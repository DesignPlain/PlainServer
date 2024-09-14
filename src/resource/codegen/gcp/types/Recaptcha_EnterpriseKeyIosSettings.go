package types

type Recaptcha_EnterpriseKeyIosSettings struct {
	// If set to true, it means allowed_bundle_ids will not be enforced.
	AllowAllBundleIds bool `json:"allowAllBundleIds,omitempty" yaml:"allowAllBundleIds,omitempty"`

	// iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname'
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty" yaml:"allowedBundleIds,omitempty"`
}
