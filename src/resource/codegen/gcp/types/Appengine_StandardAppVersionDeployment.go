package types

type Appengine_StandardAppVersionDeployment struct {
	/*
	   Manifest of the files stored in Google Cloud Storage that are included as part of this version.
	   All files must be readable using the credentials supplied with this call.
	   Structure is documented below.
	*/
	Files []Appengine_StandardAppVersionDeploymentFile `json:"files,omitempty" yaml:"files,omitempty"`

	/*
	   Zip File
	   Structure is documented below.
	*/
	Zip Appengine_StandardAppVersionDeploymentZip `json:"zip,omitempty" yaml:"zip,omitempty"`
}
