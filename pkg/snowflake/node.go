// SPDX-FileCopyrightText: 2023 Linkall Inc.
//
// SPDX-License-Identifier: Apache-2.0

package snowflake

const (
	// NodeName space: [0, 65535], DON'T CHANGE THEM!!!
	controllerNodeIDStart           = uint16(16)
	reservedControlPanelNodeIDStart = uint16(32)
	storeNodeIDStart                = uint16(1024)
	reservedNodeIDStart             = uint16(8192)
)

type Service int

const (
	ControllerService Service = iota
	StoreService
	UnknownService
)

func (s Service) Name() string { _ = "STUB: not implemented"; return "" }

type node struct {
	start uint16
	end   uint16
	id    uint16
	svc   Service
}

func NewNode(svc Service, id uint16) *node {
	_ = "STUB: not implemented" //nolint: revive // it's ok
	return nil
}

func (n *node) logicID() uint16 { _ = "STUB: not implemented"; return 0 }

func (n *node) valid() bool { _ = "STUB: not implemented"; return false }
