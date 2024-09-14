package types

type Lambda_CodePathOptions struct {
	// Packages to explicitly exclude from the Assets for a serialized closure. This can be used when clients want to trim down the size of a closure, and they know that some package won't ever actually be needed at runtime, but is still a dependency of some package that is being used at runtime.
	ExtraExcludePackages []string `json:"extraExcludePackages,omitempty" yaml:"extraExcludePackages,omitempty"`

	// Extra packages to include when producing the Assets for a serialized closure. This can be useful if the packages are acquired in a way that the serialization code does not understand. For example, if there was some sort of module that was pulled in based off of a computed string.
	ExtraIncludePackages []string `json:"extraIncludePackages,omitempty" yaml:"extraIncludePackages,omitempty"`

	// Local file/directory paths that should be included when producing the Assets for a serialized closure.
	ExtraIncludePaths []string `json:"extraIncludePaths,omitempty" yaml:"extraIncludePaths,omitempty"`
}
