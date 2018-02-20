package jsonutil

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	// jsoniter.ConfigFastest marshals the float with 6 digits precision (lossy),
	// which is significantly faster. It also does not escape HTML.
	jcf = jsoniter.ConfigFastest
)

// Marshal uses jsoniter for effieciently marshalling the given struct into a
// byte stream.
func Marshal(v interface{}) ([]byte, error) {
	// See https://github.com/json-iterator/go/blob/master/example_test.go#L47-L67
	stream := jcf.BorrowStream(nil)
	defer jcf.ReturnStream(stream)

	stream.WriteVal(v)

	if stream.Error != nil {
		return nil, errors.Wrap(stream.Error, "jsonutil: marshal using jsoniter failed")
	}

	return stream.Buffer(), nil
}

// Unmarshal uses jsoniter for efficiently unmarshalling the byte stream into
// the struct pointer.
func Unmarshal(data []byte, v interface{}) error {
	// See https://github.com/json-iterator/go/blob/master/example_test.go#L69-L88
	iter := jcf.BorrowIterator(data)
	defer jcf.ReturnIterator(iter)

	iter.ReadVal(v)
	if iter.Error != nil {
		return errors.Wrap(iter.Error, "jsonutil: unmarshal using jsoniter failed")
	}
	return nil
}
