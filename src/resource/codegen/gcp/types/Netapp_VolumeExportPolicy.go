package types

type Netapp_VolumeExportPolicy struct {
	/*
	   Export rules (up to 5) control NFS volume access.
	   Structure is documented below.
	*/
	Rules []Netapp_VolumeExportPolicyRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
