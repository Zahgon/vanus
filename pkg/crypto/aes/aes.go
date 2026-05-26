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

package aes

const nonceSize = 12

func Encrypt(value, key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Get the AES block cipher

// Get the GCM cipher mode

func Decrypt(value, key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Get the AES block cipher

// Get the GCM cipher mode

// paddingKey lt size append 0, gt size will discard.
func paddingKey(key string, size int) []byte { _ = "STUB: not implemented"; return nil }
