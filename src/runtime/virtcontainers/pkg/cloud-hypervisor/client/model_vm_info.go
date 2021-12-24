/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VmInfo Virtual Machine information
type VmInfo struct {
	Config           VmConfig               `json:"config"`
	State            string                 `json:"state"`
	MemoryActualSize *int64                 `json:"memory_actual_size,omitempty"`
	DeviceTree       *map[string]DeviceNode `json:"device_tree,omitempty"`
}

// NewVmInfo instantiates a new VmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmInfo(config VmConfig, state string) *VmInfo {
	this := VmInfo{}
	this.Config = config
	this.State = state
	return &this
}

// NewVmInfoWithDefaults instantiates a new VmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmInfoWithDefaults() *VmInfo {
	this := VmInfo{}
	return &this
}

// GetConfig returns the Config field value
func (o *VmInfo) GetConfig() VmConfig {
	if o == nil {
		var ret VmConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *VmInfo) GetConfigOk() (*VmConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *VmInfo) SetConfig(v VmConfig) {
	o.Config = v
}

// GetState returns the State field value
func (o *VmInfo) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VmInfo) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VmInfo) SetState(v string) {
	o.State = v
}

// GetMemoryActualSize returns the MemoryActualSize field value if set, zero value otherwise.
func (o *VmInfo) GetMemoryActualSize() int64 {
	if o == nil || o.MemoryActualSize == nil {
		var ret int64
		return ret
	}
	return *o.MemoryActualSize
}

// GetMemoryActualSizeOk returns a tuple with the MemoryActualSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmInfo) GetMemoryActualSizeOk() (*int64, bool) {
	if o == nil || o.MemoryActualSize == nil {
		return nil, false
	}
	return o.MemoryActualSize, true
}

// HasMemoryActualSize returns a boolean if a field has been set.
func (o *VmInfo) HasMemoryActualSize() bool {
	if o != nil && o.MemoryActualSize != nil {
		return true
	}

	return false
}

// SetMemoryActualSize gets a reference to the given int64 and assigns it to the MemoryActualSize field.
func (o *VmInfo) SetMemoryActualSize(v int64) {
	o.MemoryActualSize = &v
}

// GetDeviceTree returns the DeviceTree field value if set, zero value otherwise.
func (o *VmInfo) GetDeviceTree() map[string]DeviceNode {
	if o == nil || o.DeviceTree == nil {
		var ret map[string]DeviceNode
		return ret
	}
	return *o.DeviceTree
}

// GetDeviceTreeOk returns a tuple with the DeviceTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmInfo) GetDeviceTreeOk() (*map[string]DeviceNode, bool) {
	if o == nil || o.DeviceTree == nil {
		return nil, false
	}
	return o.DeviceTree, true
}

// HasDeviceTree returns a boolean if a field has been set.
func (o *VmInfo) HasDeviceTree() bool {
	if o != nil && o.DeviceTree != nil {
		return true
	}

	return false
}

// SetDeviceTree gets a reference to the given map[string]DeviceNode and assigns it to the DeviceTree field.
func (o *VmInfo) SetDeviceTree(v map[string]DeviceNode) {
	o.DeviceTree = &v
}

func (o VmInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.MemoryActualSize != nil {
		toSerialize["memory_actual_size"] = o.MemoryActualSize
	}
	if o.DeviceTree != nil {
		toSerialize["device_tree"] = o.DeviceTree
	}
	return json.Marshal(toSerialize)
}

type NullableVmInfo struct {
	value *VmInfo
	isSet bool
}

func (v NullableVmInfo) Get() *VmInfo {
	return v.value
}

func (v *NullableVmInfo) Set(val *VmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmInfo(val *VmInfo) *NullableVmInfo {
	return &NullableVmInfo{value: val, isSet: true}
}

func (v NullableVmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
