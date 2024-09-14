package devopsguru

import types "libds/aws/types"

type ServiceIntegration struct {
	// Information about whether DevOps Guru is configured to encrypt server-side data using KMS. See `kms_server_side_encryption` below.
	KmsServerSideEncryption types.Devopsguru_ServiceIntegrationKmsServerSideEncryption `json:"kmsServerSideEncryption,omitempty" yaml:"kmsServerSideEncryption,omitempty"`

	// Information about whether DevOps Guru is configured to perform log anomaly detection on Amazon CloudWatch log groups. See `logs_anomaly_detection` below.
	LogsAnomalyDetection types.Devopsguru_ServiceIntegrationLogsAnomalyDetection `json:"logsAnomalyDetection,omitempty" yaml:"logsAnomalyDetection,omitempty"`

	// Information about whether DevOps Guru is configured to create an OpsItem in AWS Systems Manager OpsCenter for each created insight. See `ops_center` below.
	OpsCenter types.Devopsguru_ServiceIntegrationOpsCenter `json:"opsCenter,omitempty" yaml:"opsCenter,omitempty"`
}
