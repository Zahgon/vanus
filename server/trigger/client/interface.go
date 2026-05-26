// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockgen -source=interface.go  -destination=mock_interface.go -package=client
package client

import (
	"context"
	"errors"
	nethttp "net/http"

	ce "github.com/cloudevents/sdk-go/v2"
)

type Sender interface {
	Send(ctx context.Context, events ...*ce.Event) Result
}

type EventClient interface {
	Sender
}

type Result struct {
	StatusCode int
	Err        error
}

func newInternalErr(err error) Result { _ = "STUB: not implemented"; return *new(Result) }

func newUnknownErr(err error) Result { _ = "STUB: not implemented"; return *new(Result) }

func convertHTTPResponse(statusCode int, desc string, body []byte) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

var (
	Success         = Result{}
	DeliveryTimeout = Result{errDeliveryTimeout, errors.New("DeliveryTimeout")}
)

const (
	errStatusCode = nethttp.StatusBadRequest

	errUnknown         = 600
	errDeliveryTimeout = 601
)
