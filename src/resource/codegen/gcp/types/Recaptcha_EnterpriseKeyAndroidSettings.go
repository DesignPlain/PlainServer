package types

type Recaptcha_EnterpriseKeyAndroidSettings struct {
	// If set to true, it means allowed_package_names will not be enforced.
	AllowAllPackageNames bool `json:"allowAllPackageNames,omitempty" yaml:"allowAllPackageNames,omitempty"`

	// Android package names of apps allowed to use the key. Example: 'com.companyname.appname'
	AllowedPackageNames []string `json:"allowedPackageNames,omitempty" yaml:"allowedPackageNames,omitempty"`
}
