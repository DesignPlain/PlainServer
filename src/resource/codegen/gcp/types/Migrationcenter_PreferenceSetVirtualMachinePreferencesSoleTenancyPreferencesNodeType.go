package types

type Migrationcenter_PreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesNodeType struct {
	// Name of the Sole Tenant node. Consult https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes
	NodeName string `json:"nodeName,omitempty" yaml:"nodeName,omitempty"`
}
