package workstations

import types "DesignSphere_Server/src/resource/gcp/types"

type WorkstationConfig struct {
	/*
	   Client-specified annotations. This is distinct from labels.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// Disables support for plain TCP connections in the workstation. By default the service supports TCP connections via a websocket relay. Setting this option to true disables that relay, which prevents the usage of services that require plain tcp connections, such as ssh. When enabled, all communication must occur over https or wss.
	DisableTcpConnections bool `json:"disableTcpConnections,omitempty" yaml:"disableTcpConnections,omitempty"`

	/*
	   How long to wait before automatically stopping an instance that hasn't recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	IdleTimeout string `json:"idleTimeout,omitempty" yaml:"idleTimeout,omitempty"`

	/*
	   Readiness checks to be performed on a workstation.
	   Structure is documented below.
	*/
	ReadinessChecks []types.Workstations_WorkstationConfigReadinessCheck `json:"readinessChecks,omitempty" yaml:"readinessChecks,omitempty"`

	/*
	   How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryption_key` is set. Defaults to 12 hours.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	RunningTimeout string `json:"runningTimeout,omitempty" yaml:"runningTimeout,omitempty"`

	// The ID of the parent workstation cluster.
	WorkstationClusterId string `json:"workstationClusterId,omitempty" yaml:"workstationClusterId,omitempty"`

	// Human-readable name for this resource.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Encrypts resources of this workstation configuration using a customer-managed encryption key.
	   If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
	   If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
	   If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
	   Structure is documented below.
	*/
	EncryptionKey types.Workstations_WorkstationConfigEncryptionKey `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`

	/*
	   Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Directories to persist across workstation sessions.
	   Structure is documented below.
	*/
	PersistentDirectories []types.Workstations_WorkstationConfigPersistentDirectory `json:"persistentDirectories,omitempty" yaml:"persistentDirectories,omitempty"`

	// The ID to be assigned to the workstation cluster config.
	WorkstationConfigId string `json:"workstationConfigId,omitempty" yaml:"workstationConfigId,omitempty"`

	// Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
	EnableAuditAgent bool `json:"enableAuditAgent,omitempty" yaml:"enableAuditAgent,omitempty"`

	/*
	   Runtime host for a workstation.
	   Structure is documented below.
	*/
	Host types.Workstations_WorkstationConfigHost `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   The location where the workstation cluster config should reside.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster's region must be specified—for example, `['us-central1-a', 'us-central1-f']`.
	   If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
	*/
	ReplicaZones []string `json:"replicaZones,omitempty" yaml:"replicaZones,omitempty"`

	/*
	   Container that will be run for each workstation using this configuration when that workstation is started.
	   Structure is documented below.
	*/
	Container types.Workstations_WorkstationConfigContainer `json:"container,omitempty" yaml:"container,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
