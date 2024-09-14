package storagegateway

type WorkingStorage struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId string `json:"diskId,omitempty" yaml:"diskId,omitempty"`

	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`
}
