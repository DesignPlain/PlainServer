package opensearch

import types "libds/aws/types"

type Package struct {
	// Description of the package.
	PackageDescription string `json:"packageDescription,omitempty" yaml:"packageDescription,omitempty"`

	// Unique name for the package.
	PackageName string `json:"packageName,omitempty" yaml:"packageName,omitempty"`

	// Configuration block for the package source options.
	PackageSource types.Opensearch_PackagePackageSource `json:"packageSource,omitempty" yaml:"packageSource,omitempty"`

	// The type of package.
	PackageType string `json:"packageType,omitempty" yaml:"packageType,omitempty"`
}
