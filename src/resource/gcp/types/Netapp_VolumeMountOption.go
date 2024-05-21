package types

type Netapp_VolumeMountOption struct {
	/*
	   (Output)
	   Full export path of the volume.
	   Format for NFS volumes: `<export_ip>:/<shareName>`
	   Format for SMB volumes: `\\\\netbios_prefix-four_random_hex_letters.domain_name\\shareName`
	*/
	ExportFull string `json:"exportFull,omitempty" yaml:"exportFull,omitempty"`

	/*
	   (Output)
	   Human-readable mount instructions.
	*/
	Instructions string `json:"instructions,omitempty" yaml:"instructions,omitempty"`

	/*
	   (Output)
	   Protocol to mount with.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   (Output)
	   Export path of the volume.
	*/
	Export string `json:"export,omitempty" yaml:"export,omitempty"`
}
