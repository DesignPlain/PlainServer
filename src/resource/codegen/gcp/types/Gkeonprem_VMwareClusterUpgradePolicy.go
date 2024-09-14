package types

type Gkeonprem_VMwareClusterUpgradePolicy struct {
	// Controls whether the upgrade applies to the control plane only.
	ControlPlaneOnly bool `json:"controlPlaneOnly,omitempty" yaml:"controlPlaneOnly,omitempty"`
}
