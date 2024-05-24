package macie2

import types "DesignSphere_Server/src/resource/aws/types"

type ClassificationExportConfiguration struct {
	// Configuration block for a S3 Destination. Defined below
	S3Destination types.Macie2_ClassificationExportConfigurationS3Destination `json:"s3Destination,omitempty" yaml:"s3Destination,omitempty"`
}
