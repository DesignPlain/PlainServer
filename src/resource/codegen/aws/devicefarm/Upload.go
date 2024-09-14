package devicefarm

type Upload struct {
	// The upload's content type (for example, application/octet-stream).
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of the project for the upload.
	ProjectArn string `json:"projectArn,omitempty" yaml:"projectArn,omitempty"`

	// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
