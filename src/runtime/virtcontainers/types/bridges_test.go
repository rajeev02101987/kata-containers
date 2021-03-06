// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAddRemoveDevice(t *testing.T, b *Bridge) {
	assert := assert.New(t)

	// add device
	devID := "abc123"

	addr, err := b.AddDevice(devID)
	assert.NoError(err)
	if addr < 1 {
		assert.Fail("address cannot be less than 1")
	}

	// remove device
	err = b.RemoveDevice("")
	assert.Error(err)

	err = b.RemoveDevice(devID)
	assert.NoError(err)

	// add device when the bridge is full
	b.Devices = make(map[uint32]string)
	for i := uint32(1); i <= b.MaxCapacity; i++ {
		b.Devices[i] = fmt.Sprintf("%d", i)
	}
	addr, err = b.AddDevice(devID)
	assert.Error(err)
	if addr != 0 {
		assert.Fail("address should be 0")
	}
}

func TestAddRemoveDevicePCI(t *testing.T) {

	// create a pci bridge
	bridges := []*Bridge{{make(map[uint32]string), "rgb123", 5, PCI, PCIBridgeMaxCapacity}}

	testAddRemoveDevice(t, bridges[0])
}

func TestAddRemoveDeviceCCW(t *testing.T) {

	// create a CCW bridge
	bridges := []*Bridge{{make(map[uint32]string), "rgb123", 5, CCW, CCWBridgeMaxCapacity}}

	testAddRemoveDevice(t, bridges[0])
}
