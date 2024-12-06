package oss

import "errors"

var (
	ErrBucketEmpty       = errors.New("Bucket must be not empty")
	ErrEndpointEmpty     = errors.New("Endpoint must be not empty")
	ErrRegionEmpty       = errors.New("Region must be not empty")
	ErrEmptyOriginalName = errors.New("original name must not be empty")
	ErrEmptySizeName     = errors.New("size must greater than 0")
	ErrEmptyDelUrl       = errors.New("url must be not empty")
	ErrNotSameBaseUrl    = errors.New("only can del same base file")
)
