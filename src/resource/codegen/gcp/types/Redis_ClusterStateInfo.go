package types

type Redis_ClusterStateInfo struct {
	/*
	   A nested object resource
	   Structure is documented below.
	*/
	UpdateInfo Redis_ClusterStateInfoUpdateInfo `json:"updateInfo,omitempty" yaml:"updateInfo,omitempty"`
}
