package types

type Compute_SecurityPolicyAdaptiveProtectionConfig struct {
	/*
	   Configuration for [Automatically deploy Adaptive Protection suggested rules](https://cloud.google.com/armor/docs/adaptive-protection-auto-deploy?hl=en). Structure is documented below.

	   <a name="nested_layer_7_ddos_defense_config"></a>The `layer_7_ddos_defense_config` block supports:
	*/
	AutoDeployConfig Compute_SecurityPolicyAdaptiveProtectionConfigAutoDeployConfig `json:"autoDeployConfig,omitempty" yaml:"autoDeployConfig,omitempty"`

	// Configuration for [Google Cloud Armor Adaptive Protection Layer 7 DDoS Defense](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
	Layer7DdosDefenseConfig Compute_SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig `json:"layer7DdosDefenseConfig,omitempty" yaml:"layer7DdosDefenseConfig,omitempty"`
}
