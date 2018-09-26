// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package networkserver

import (
	"context"

	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
)

type DeviceRegistry interface {
	GetByEUI(ctx context.Context, joinEUI types.EUI64, devEUI types.EUI64) (*ttnpb.EndDevice, error)
	GetByID(ctx context.Context, appID ttnpb.ApplicationIdentifiers, devID string) (*ttnpb.EndDevice, error)
	RangeByAddr(devAddr types.DevAddr, f func(*ttnpb.EndDevice) bool) error
	SetByID(ctx context.Context, appID ttnpb.ApplicationIdentifiers, devID string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, error)) (*ttnpb.EndDevice, error)
}

func DeleteDevice(ctx context.Context, r DeviceRegistry, appID ttnpb.ApplicationIdentifiers, devID string) error {
	_, err := r.SetByID(ctx, appID, devID, func(*ttnpb.EndDevice) (*ttnpb.EndDevice, error) { return nil, nil })
	return err
}

func CreateDevice(ctx context.Context, r DeviceRegistry, dev *ttnpb.EndDevice) (*ttnpb.EndDevice, error) {
	dev, err := r.SetByID(ctx, dev.EndDeviceIdentifiers.ApplicationIdentifiers, dev.EndDeviceIdentifiers.DeviceID, func(stored *ttnpb.EndDevice) (*ttnpb.EndDevice, error) {
		if stored != nil {
			return nil, errDuplicateIdentifiers
		}
		return dev, nil
	})
	if err != nil {
		return nil, err
	}
	return dev, nil
}