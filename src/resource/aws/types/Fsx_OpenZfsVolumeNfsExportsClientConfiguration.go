package types

type Fsx_OpenZfsVolumeNfsExportsClientConfiguration struct {
	// A value that specifies who can mount the file system. You can provide a wildcard character (-), an IP address (0.0.0.0), or a CIDR address (192.0.2.0/24. By default, Amazon FSx uses the wildcard character when specifying the client.
	Clients string `json:"clients,omitempty" yaml:"clients,omitempty"`

	// The options to use when mounting the file system. Maximum of 20 items. See the [Linix NFS exports man page](https://linux.die.net/man/5/exports) for more information. `crossmount` and `sync` are used by default.
	Options []string `json:"options,omitempty" yaml:"options,omitempty"`
}
