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

package cluster

import (
	"context"

	"github.com/vanus-labs/vanus/api/cluster/raw_client"
	ctrlpb "github.com/vanus-labs/vanus/api/controller"
	metapb "github.com/vanus-labs/vanus/api/meta"
)

type authService struct {
	client ctrlpb.AuthControllerClient
}

func (a *authService) GetUserRole(ctx context.Context, user string) ([]*metapb.UserRole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authService) GetUserByToken(ctx context.Context, token string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *authService) RawClient() ctrlpb.AuthControllerClient {
	_ = "STUB: not implemented"
	return *new(ctrlpb.AuthControllerClient)
}

func newAuthService(cc *raw_client.Conn) AuthService {
	_ = "STUB: not implemented"
	return *new(AuthService)
}
