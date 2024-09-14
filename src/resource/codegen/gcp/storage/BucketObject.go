package storage

import types "libds/gcp/types"

type BucketObject struct {
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage string `json:"contentLanguage,omitempty" yaml:"contentLanguage,omitempty"`

	// The name of the object. If you're interpolating the name of this object, see `output_name` instead.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the containing bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
	Retention types.Storage_BucketObjectRetention `json:"retention,omitempty" yaml:"retention,omitempty"`

	/*
	   The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	   Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	   storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	*/
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold bool `json:"temporaryHold,omitempty" yaml:"temporaryHold,omitempty"`

	/*
	   [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	   directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	*/
	CacheControl string `json:"cacheControl,omitempty" yaml:"cacheControl,omitempty"`

	// Data as `string` to be uploaded. Must be defined if `source` is not. --Note--: The `content` field is marked as sensitive.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition string `json:"contentDisposition,omitempty" yaml:"contentDisposition,omitempty"`

	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding string `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`

	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	/*
	   Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	   Structure is documented below.
	*/
	CustomerEncryption types.Storage_BucketObjectCustomerEncryption `json:"customerEncryption,omitempty" yaml:"customerEncryption,omitempty"`

	// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	//
	DetectMd5hash string `json:"detectMd5hash,omitempty" yaml:"detectMd5hash,omitempty"`

	// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold bool `json:"eventBasedHold,omitempty" yaml:"eventBasedHold,omitempty"`

	/*
	   User-provided metadata, in key/value pairs.

	   One of the following is required:
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   A path to the data you want to upload. Must be defined
	   if `content` is not.

	   - - -
	*/
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
