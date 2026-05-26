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

package pkg

type CredentialType string

const (
	Plain  CredentialType = "plain"
	AWS    CredentialType = "aws"
	GCloud CredentialType = "gcloud"

	SecretsMask = "******"
)

type SinkCredential interface {
	GetType() CredentialType
}

func FillSinkCredential(dst, src SinkCredential) { _ = "STUB: not implemented"; return }

type PlainSinkCredential struct {
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
}

func NewPlainSinkCredential(identifier, secret string) SinkCredential {
	_ = "STUB: not implemented"
	return *new(SinkCredential)
}

func (c *PlainSinkCredential) GetType() CredentialType {
	_ = "STUB: not implemented"
	return *new(CredentialType)
}

type AkSkSinkCredential struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

func NewAkSkSinkCredential(accessKeyID, secretAccessKey string) SinkCredential {
	_ = "STUB: not implemented"
	return *new(SinkCredential)
}

func (c *AkSkSinkCredential) GetType() CredentialType {
	_ = "STUB: not implemented"
	return *new(CredentialType)
}

type GCloudSinkCredential struct {
	CredentialJSON string `json:"credential_json"`
}

func NewGCloudSinkCredential(credentialJSON string) SinkCredential {
	_ = "STUB: not implemented"
	return *new(SinkCredential)
}

func (c *GCloudSinkCredential) GetType() CredentialType {
	_ = "STUB: not implemented"
	return *new(CredentialType)
}
