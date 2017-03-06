// Copyright 2016 Mender Software AS
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

// generate with:
// mockery -name=DataStore  -inpkg -print

package main

import log "github.com/mendersoftware/go-lib-micro/log"
import migrate "github.com/mendersoftware/go-lib-micro/mongo/migrate"
import mock "github.com/stretchr/testify/mock"

// MockDataStore is an autogenerated mock type for the DataStore type
type MockDataStore struct {
	mock.Mock
}

// AddAuthSet provides a mock function with given fields: set
func (_m *MockDataStore) AddAuthSet(set AuthSet) error {
	ret := _m.Called(set)

	var r0 error
	if rf, ok := ret.Get(0).(func(AuthSet) error); ok {
		r0 = rf(set)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddDevice provides a mock function with given fields: d
func (_m *MockDataStore) AddDevice(d Device) error {
	ret := _m.Called(d)

	var r0 error
	if rf, ok := ret.Get(0).(func(Device) error); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddToken provides a mock function with given fields: t
func (_m *MockDataStore) AddToken(t Token) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(Token) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteToken provides a mock function with given fields: jti
func (_m *MockDataStore) DeleteToken(jti string) error {
	ret := _m.Called(jti)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(jti)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTokenByDevId provides a mock function with given fields: dev_id
func (_m *MockDataStore) DeleteTokenByDevId(dev_id string) error {
	ret := _m.Called(dev_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dev_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAuthSetByDataKey provides a mock function with given fields: data, key
func (_m *MockDataStore) GetAuthSetByDataKey(data string, key string) (*AuthSet, error) {
	ret := _m.Called(data, key)

	var r0 *AuthSet
	if rf, ok := ret.Get(0).(func(string, string) *AuthSet); ok {
		r0 = rf(data, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AuthSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(data, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthSetById provides a mock function with given fields: id
func (_m *MockDataStore) GetAuthSetById(id string) (*AuthSet, error) {
	ret := _m.Called(id)

	var r0 *AuthSet
	if rf, ok := ret.Get(0).(func(string) *AuthSet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AuthSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceById provides a mock function with given fields: id
func (_m *MockDataStore) GetDeviceById(id string) (*Device, error) {
	ret := _m.Called(id)

	var r0 *Device
	if rf, ok := ret.Get(0).(func(string) *Device); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceByIdentityData provides a mock function with given fields: idata
func (_m *MockDataStore) GetDeviceByIdentityData(idata string) (*Device, error) {
	ret := _m.Called(idata)

	var r0 *Device
	if rf, ok := ret.Get(0).(func(string) *Device); ok {
		r0 = rf(idata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevices provides a mock function with given fields: skip, limit
func (_m *MockDataStore) GetDevices(skip uint, limit uint) ([]Device, error) {
	ret := _m.Called(skip, limit)

	var r0 []Device
	if rf, ok := ret.Get(0).(func(uint, uint) []Device); ok {
		r0 = rf(skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(skip, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetToken provides a mock function with given fields: jti
func (_m *MockDataStore) GetToken(jti string) (*Token, error) {
	ret := _m.Called(jti)

	var r0 *Token
	if rf, ok := ret.Get(0).(func(string) *Token); ok {
		r0 = rf(jti)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jti)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields: version, migrations
func (_m *MockDataStore) Migrate(version string, migrations []migrate.Migration) error {
	ret := _m.Called(version, migrations)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []migrate.Migration) error); ok {
		r0 = rf(version, migrations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAuthSet provides a mock function with given fields: orig, mod
func (_m *MockDataStore) UpdateAuthSet(orig AuthSet, mod AuthSetUpdate) error {
	ret := _m.Called(orig, mod)

	var r0 error
	if rf, ok := ret.Get(0).(func(AuthSet, AuthSetUpdate) error); ok {
		r0 = rf(orig, mod)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: d
func (_m *MockDataStore) UpdateDevice(d *Device) error {
	ret := _m.Called(d)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Device) error); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseLog provides a mock function with given fields: l
func (_m *MockDataStore) UseLog(l *log.Logger) {
	_m.Called(l)
}

var _ DataStore = (*MockDataStore)(nil)
