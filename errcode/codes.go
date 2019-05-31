package errcode

import "errors"

// List of all the error Codes used
var (
	ErrProviderResourceNotRead       = errors.New("the resource did not return an ID")
	ErrProviderResourceDoNotMatchTag = errors.New("the resource does not match the required tags")
	ErrProviderResourceAutogenerated = errors.New("the resource is autogenerated and should not be imported")

	ErrCacheKeyNotFound        = errors.New("the key used to search was not found")
	ErrCacheKeyAlreadyExisting = errors.New("the key already exists on the cache")
)
