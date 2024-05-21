package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet struct {
	/*
	   The regex-filtered set of files to scan.
	   Structure is documented below.
	*/
	RegexFileSet Dataloss_PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet `json:"regexFileSet,omitempty" yaml:"regexFileSet,omitempty"`

	/*
	   The Cloud Storage url of the file(s) to scan, in the format `gs://<bucket>/<path>`. Trailing wildcard
	   in the path is allowed.
	   If the url ends in a trailing slash, the bucket or directory represented by the url will be scanned
	   non-recursively (content in sub-directories will not be scanned). This means that `gs://mybucket/` is
	   equivalent to `gs://mybucket/-`, and `gs://mybucket/directory/` is equivalent to `gs://mybucket/directory/-`.
	*/
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
