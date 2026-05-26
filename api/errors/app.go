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

package errors

func New(desc string) *ErrorType { _ = "STUB: not implemented"; return nil }

func Convert(str string) (*ErrorType, bool) { _ = "STUB: not implemented"; return nil, false }

type ErrorType struct {
	Description    string    `json:"description"`
	Message        string    `json:"message"`
	Code           ErrorCode `json:"code"`
	underlayErrors []error
}

func (e *ErrorType) WithGRPCCode(c ErrorCode) *ErrorType { _ = "STUB: not implemented"; return nil }

// WithMessage add additional message to explain what try to do cause this error.
// the explanation was used to improve understandability of the error in order to make
// people know what they should do
func (e *ErrorType) WithMessage(str string) *ErrorType { _ = "STUB: not implemented"; return nil }

// Wrap the other error as the underlay errors of this error. sometimes we return an error because
// of another error(named underlay error). So, we should add the underlay error to this error's context.
// By this, the people can understand why this error they received
func (e *ErrorType) Wrap(err error) *ErrorType { _ = "STUB: not implemented"; return nil }

func (e *ErrorType) JSON() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorType) copy() *ErrorType { _ = "STUB: not implemented"; return nil }

// Error return readable error message by JSON format
func (e ErrorType) Error() string { _ = "STUB: not implemented"; return "" }
