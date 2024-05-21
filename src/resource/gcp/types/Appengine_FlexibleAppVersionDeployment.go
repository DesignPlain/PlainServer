package types

type Appengine_FlexibleAppVersionDeployment struct {
	/*
	   Options for the build operations performed as a part of the version deployment. Only applicable when creating a version using source code directly.
	   Structure is documented below.
	*/
	CloudBuildOptions Appengine_FlexibleAppVersionDeploymentCloudBuildOptions `json:"cloudBuildOptions,omitempty" yaml:"cloudBuildOptions,omitempty"`

	/*
	   The Docker image for the container that runs the version.
	   Structure is documented below.
	*/
	Container Appengine_FlexibleAppVersionDeploymentContainer `json:"container,omitempty" yaml:"container,omitempty"`

	/*
	   Manifest of the files stored in Google Cloud Storage that are included as part of this version.
	   All files must be readable using the credentials supplied with this call.
	   Structure is documented below.
	*/
	Files []Appengine_FlexibleAppVersionDeploymentFile `json:"files,omitempty" yaml:"files,omitempty"`

	/*
	   Zip File
	   Structure is documented below.
	*/
	Zip Appengine_FlexibleAppVersionDeploymentZip `json:"zip,omitempty" yaml:"zip,omitempty"`
}
