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

// Chain method can group many errors as a single error. this is a helpful method if there are many errors
// will be occurred in one method.
//
// Example: if we have a loop in one method, and each iterator create a goroutine, each goroutine may
// an error occurred, but we have to wait all goroutine done, in this situation, we can use this method
// to collect all errors occurred in different goroutine.
//
// Notice: this method is not CONCURRENCY SAFETY. you can see somewhere call this function to know more
// details about how to use this method.
func Chain(errs ...error) error { _ = "STUB: not implemented"; return nil }

func Is(err error, target error) bool { _ = "STUB: not implemented"; return false }

func UnwrapOrUnknown(err error) error { _ = "STUB: not implemented"; return nil }

func FromError(err error) (*ErrorType, bool) { _ = "STUB: not implemented"; return nil, false }

// ConvertToGRPCError convert an internal error to an exported error defined in gRPC.
func ConvertToGRPCError(err error) error { _ = "STUB: not implemented"; return nil }
