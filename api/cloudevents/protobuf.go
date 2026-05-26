/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package cloudevents

import (
	// standard libraries.

	stdtime "time"

	// third-party libraries.
	"github.com/cloudevents/sdk-go/v2/event"
)

const (
	// ContentTypeProtobuf indicates that the data attribute is a protobuf message.
	ContentTypeProtobuf = "application/protobuf"
)

const (
	datacontenttype = "datacontenttype"
	dataschema      = "dataschema"
	subject         = "subject"
	time            = "time"
)

var zeroTime = stdtime.Time{}

// ToProto convert an SDK event to a protobuf variant of the event that can be marshaled.
func ToProto(e *event.Event) (*CloudEvent, error) { _ = "STUB: not implemented"; return nil, nil }

func attributeFor(v interface{}) (*CloudEvent_CloudEventAttributeValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func valueFrom(attr *CloudEvent_CloudEventAttributeValue) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromProto Convert from a protobuf variant into the generic, SDK event.
func FromProto(container *CloudEvent) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: There are some issues around missing data content type values that
// are still unresolved. It is an optional field and if unset then it is
// implied that the encoding used for the envelope was also used for the
// data. However, there is no mapping that exists between data content types
// and the envelope content types. For example, how would this system know
// that receiving an envelope in application/cloudevents+protobuf know that
// the implied data content type if missing is application/protobuf.
//
// It is also not clear what should happen if the data content type is unset
// but it is known that the data content type is _not_ the same as the
// envelope. For example, a JSON encoded data value would be stored within
// the BinaryData attribute of the protobuf formatted envelope. Protobuf
// data values, however, are _always_ stored as a protobuf encoded Any type
// within the ProtoData field. Any use of the BinaryData or TextData fields
// means the value is _not_ protobuf. If content type is not set then have
// no way of knowing what the data encoding actually is. Currently, this
// code does not address this and only loads explicitly set data content
// type values.

// NOTE: If we use SetData then the current implementation always sets
// the Base64 bit to true. Direct assignment appears to be the only way
// to set non-base64 encoded binary data.
// if err := e.SetData(contentType, dt.BinaryData); err != nil {
// 	return nil, fmt.Errorf("failed to convert binary type (%s) data: %s", contentType, err)
// }
