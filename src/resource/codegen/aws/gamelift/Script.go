package gamelift

import types "libds/aws/types"

type Script struct {
	// Name of the script
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Information indicating where your game script files are stored. See below.
	StorageLocation types.Gamelift_ScriptStorageLocation `json:"storageLocation,omitempty" yaml:"storageLocation,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Version that is associated with this script.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile string `json:"zipFile,omitempty" yaml:"zipFile,omitempty"`
}
