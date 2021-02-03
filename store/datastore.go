// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mendersoftware/deviceconfig/model"
)

var (
	ErrDeviceNoExist       = errors.New("device does not exist")
	ErrDeviceAlreadyExists = errors.New("device already exists")
)

// DataStore interface for DataStore services
//nolint:lll - skip line length check for interface declaration.
//go:generate ../x/mockgen.sh
type DataStore interface {
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	DropDatabase(ctx context.Context) error

	// Migrate applies migrations to tenants in the provided context. That
	// is, if ctx has an associated identity.Identity set, it will ONLY
	// migrate the single tenant's db. If ctx does not have an identity,
	// all deviceconfig collections will be migrated to the given version.
	Migrate(ctx context.Context, version string, automigrate bool) error

	// MigrateLatest calls Migrate with the latest schema version.
	MigrateLatest(ctx context.Context) error

	// InsertDeviceConfig inserts a new device configuration
	InsertDevice(ctx context.Context, dev model.Device) error

	// UpsertDeviceConfig updates or inserts a new device configuration
	UpsertConfiguration(ctx context.Context, dev model.Device) error

	// UpsertReportedConfiguration updates or inserts a new device reported configuration
	UpsertReportedConfiguration(ctx context.Context, dev model.Device) error

	// DeleteDevice removes the device object with the given ID from the database.
	DeleteDevice(ctx context.Context, devID uuid.UUID) error

	// GetDevice returns a device
	GetDevice(ctx context.Context, devID uuid.UUID) (model.Device, error)
}
