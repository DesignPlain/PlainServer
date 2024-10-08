package storagegateway

type UploadBuffer struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId string `json:"diskId,omitempty" yaml:"diskId,omitempty"`

	// Local disk path. For example, `/dev/nvme1n1`.
	DiskPath string `json:"diskPath,omitempty" yaml:"diskPath,omitempty"`

	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`
}
