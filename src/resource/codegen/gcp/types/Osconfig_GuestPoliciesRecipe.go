package types

type Osconfig_GuestPoliciesRecipe struct {
	/*
	   Resources available to be used in the steps in the recipe.
	   Structure is documented below.
	*/
	Artifacts []Osconfig_GuestPoliciesRecipeArtifact `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`

	/*
	   Default is INSTALLED. The desired state the agent should maintain for this recipe.
	   INSTALLED: The software recipe is installed on the instance but won't be updated to new versions.
	   INSTALLED_KEEP_UPDATED: The software recipe is installed on the instance. The recipe is updated to a higher version,
	   if a higher version of the recipe is assigned to this instance.
	   REMOVE: Remove is unsupported for software recipes and attempts to create or update a recipe to the REMOVE state is rejected.
	   Default value is `INSTALLED`.
	   Possible values are: `INSTALLED`, `UPDATED`, `REMOVED`.
	*/
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	/*
	   Actions to be taken for installing this recipe. On failure it stops executing steps and does not attempt another installation.
	   Any steps taken (including partially completed steps) are not rolled back.
	   Structure is documented below.
	*/
	InstallSteps []Osconfig_GuestPoliciesRecipeInstallStep `json:"installSteps,omitempty" yaml:"installSteps,omitempty"`

	/*
	   Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
	   Names are also used to identify resources which helps to determine whether guest policies have conflicts.
	   This means that requests to create multiple recipes with the same name and version are rejected since they
	   could potentially have conflicting assignments.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Actions to be taken for updating this recipe. On failure it stops executing steps and does not attempt another update for this recipe.
	   Any steps taken (including partially completed steps) are not rolled back.
	   Structure is documented below.
	*/
	UpdateSteps []Osconfig_GuestPoliciesRecipeUpdateStep `json:"updateSteps,omitempty" yaml:"updateSteps,omitempty"`

	// The version of this software recipe. Version can be up to 4 period separated numbers (e.g. 12.34.56.78).
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
