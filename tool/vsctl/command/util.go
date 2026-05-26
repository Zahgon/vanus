// Copyright 2023 Linkall Inc.
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

package command

import (
	"regexp"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const (
	dns1123LabelFmt          string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
	dns1123SubdomainFmt      string = dns1123LabelFmt + "(\\." + dns1123LabelFmt + ")*"
	dns1123SubdomainErrorMsg string = "a lowercase RFC 1123 subdomain must consist of lower case alphanumeric " +
		"characters, '-' or '.', and must start and end with an alphanumeric character"
	// DNS1123SubdomainMaxLength is a subdomain's max length in DNS (RFC 1123).
	DNS1123SubdomainMaxLength int = 253
)

var dns1123SubdomainRegexp = regexp.MustCompile("^" + dns1123SubdomainFmt + "$")

func cmdFailedf(cmd *cobra.Command, format string, a ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func cmdFailedWithHelpNotice(cmd *cobra.Command, format string) { _ = "STUB: not implemented"; return }

func operatorIsDeployed(_ *cobra.Command, endpoint string) bool {
	_ = "STUB: not implemented"
	return false
}

func getOperatorEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }

func LoadConfig(filename string, config interface{}) error { _ = "STUB: not implemented"; return nil }

// IsDNS1123Subdomain tests for a string that conforms to the definition of a
// subdomain in DNS (RFC 1123).
func IsDNS1123Subdomain(value string) bool { _ = "STUB: not implemented"; return false }

func Error(err error) string { _ = "STUB: not implemented"; return "" }

func formatID(id uint64) string { _ = "STUB: not implemented"; return "" }

func getColumnConfig(header table.Row) []table.ColumnConfig { _ = "STUB: not implemented"; return nil }
