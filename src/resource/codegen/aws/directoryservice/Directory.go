package directoryservice

import types "libds/aws/types"

type Directory struct {
	// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
	DesiredNumberOfDomainControllers int `json:"desiredNumberOfDomainControllers,omitempty" yaml:"desiredNumberOfDomainControllers,omitempty"`

	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso bool `json:"enableSso,omitempty" yaml:"enableSso,omitempty"`

	// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
	Size string `json:"size,omitempty" yaml:"size,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// VPC related information about the directory. Fields documented below.
	VpcSettings types.Directoryservice_DirectoryVpcSettings `json:"vpcSettings,omitempty" yaml:"vpcSettings,omitempty"`

	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Connector related information about the directory. Fields documented below.
	ConnectSettings types.Directoryservice_DirectoryConnectSettings `json:"connectSettings,omitempty" yaml:"connectSettings,omitempty"`

	// The fully qualified name for the directory, such as `corp.example.com`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The password for the directory administrator or connector user.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The short name of the directory, such as `CORP`.
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`

	// A textual description for the directory.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`
}
