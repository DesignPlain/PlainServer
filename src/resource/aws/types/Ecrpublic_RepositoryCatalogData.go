package types

type Ecrpublic_RepositoryCatalogData struct {
	// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
	AboutText string `json:"aboutText,omitempty" yaml:"aboutText,omitempty"`

	// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `ARM`, `ARM 64`, `x86`, `x86-64`
	Architectures []string `json:"architectures,omitempty" yaml:"architectures,omitempty"`

	// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
	LogoImageBlob string `json:"logoImageBlob,omitempty" yaml:"logoImageBlob,omitempty"`

	// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: `Linux`, `Windows`
	OperatingSystems []string `json:"operatingSystems,omitempty" yaml:"operatingSystems,omitempty"`

	// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
	UsageText string `json:"usageText,omitempty" yaml:"usageText,omitempty"`
}
