/*
Copyright 2018 OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package probe

import (
	"errors"
	"sync"
	"testing"

	apis "github.com/openebs/node-disk-manager/api/v1alpha1"
	"github.com/openebs/node-disk-manager/blockdevice"
	"github.com/openebs/node-disk-manager/cmd/ndm_daemonset/controller"
	libudevwrapper "github.com/openebs/node-disk-manager/pkg/udev"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type alwaysTrueFilter struct{}

func (nf *alwaysTrueFilter) Start() {}

func (nf *alwaysTrueFilter) Include(fakeDiskInfo *blockdevice.BlockDevice) bool {
	return true
}

func (nf *alwaysTrueFilter) Exclude(fakeDiskInfo *blockdevice.BlockDevice) bool {
	return true
}

func mockOsDiskToAPI() (apis.BlockDevice, error) {
	mockOsDiskDetails, err := libudevwrapper.MockDiskDetails()
	if err != nil {
		return apis.BlockDevice{}, err
	}
	fakeDetails := apis.DeviceDetails{
		Model:  mockOsDiskDetails.Model,
		Serial: mockOsDiskDetails.Serial,
		Vendor: mockOsDiskDetails.Vendor,
	}
	fakeObj := apis.DeviceSpec{
		Path:        mockOsDiskDetails.DevNode,
		Details:     fakeDetails,
		Partitioned: controller.NDMNotPartitioned,
	}

	devLinks := make([]apis.DeviceDevLink, 0)
	if len(mockOsDiskDetails.ByIdDevLinks) != 0 {
		byIdLinks := apis.DeviceDevLink{
			Kind:  "by-id",
			Links: mockOsDiskDetails.ByIdDevLinks,
		}
		devLinks = append(devLinks, byIdLinks)
	}
	if len(mockOsDiskDetails.ByPathDevLinks) != 0 {
		byPathLinks := apis.DeviceDevLink{
			Kind:  "by-path",
			Links: mockOsDiskDetails.ByPathDevLinks,
		}
		devLinks = append(devLinks, byPathLinks)
	}
	fakeObj.DevLinks = devLinks

	fakeTypeMeta := metav1.TypeMeta{
		Kind:       controller.NDMBlockDeviceKind,
		APIVersion: controller.NDMVersion,
	}
	fakeObjectMeta := metav1.ObjectMeta{
		Labels: make(map[string]string),
		Name:   mockOsDiskDetails.Uid,
	}
	fakeDiskStatus := apis.DeviceStatus{
		State:      controller.NDMActive,
		ClaimState: apis.BlockDeviceUnclaimed,
	}
	fakeDr := apis.BlockDevice{
		TypeMeta:   fakeTypeMeta,
		ObjectMeta: fakeObjectMeta,
		Spec:       fakeObj,
		Status:     fakeDiskStatus,
	}
	return fakeDr, nil
}

func TestFillDiskDetails(t *testing.T) {
	mockOsDiskDetails, err := libudevwrapper.MockDiskDetails()
	if err != nil {
		t.Fatal(err)
	}
	uProbe := udevProbe{}
	actualDiskInfo := &blockdevice.BlockDevice{}
	actualDiskInfo.SysPath = mockOsDiskDetails.SysPath
	uProbe.FillBlockDeviceDetails(actualDiskInfo)
	expectedDiskInfo := &blockdevice.BlockDevice{}
	expectedDiskInfo.SysPath = mockOsDiskDetails.SysPath
	expectedDiskInfo.DevPath = mockOsDiskDetails.DevNode
	expectedDiskInfo.DeviceAttributes.Model = mockOsDiskDetails.Model
	expectedDiskInfo.DeviceAttributes.Serial = mockOsDiskDetails.Serial
	expectedDiskInfo.DeviceAttributes.Vendor = mockOsDiskDetails.Vendor
	expectedDiskInfo.DeviceAttributes.WWN = mockOsDiskDetails.Wwn
	expectedDiskInfo.PartitionInfo.PartitionTableType = mockOsDiskDetails.PartTableType
	expectedDiskInfo.DeviceAttributes.IDType = mockOsDiskDetails.IdType
	if len(mockOsDiskDetails.ByIdDevLinks) > 0 {
		expectedDiskInfo.DevLinks = append(expectedDiskInfo.DevLinks, blockdevice.DevLink{
			Kind:  libudevwrapper.BY_ID_LINK,
			Links: mockOsDiskDetails.ByIdDevLinks,
		})
	}
	if len(mockOsDiskDetails.ByPathDevLinks) > 0 {
		expectedDiskInfo.DevLinks = append(expectedDiskInfo.DevLinks, blockdevice.DevLink{
			Kind:  libudevwrapper.BY_PATH_LINK,
			Links: mockOsDiskDetails.ByPathDevLinks,
		})
	}
	if len(mockOsDiskDetails.SymLinks) > 0 {
		expectedDiskInfo.DevLinks = append(expectedDiskInfo.DevLinks, blockdevice.DevLink{
			Kind:  libudevwrapper.SYMLINK,
			Links: mockOsDiskDetails.SymLinks,
		})
	}

	// The devlinks are compared separately as the ordering of devlinks can be different in some systems
	// eg: ubuntu 20.04 in github actions
	assert.True(t, compareDevLinks(expectedDiskInfo.DevLinks, actualDiskInfo.DevLinks))

	// The devlinks are made nil since they are already compared
	expectedDiskInfo.DevLinks = nil
	actualDiskInfo.DevLinks = nil

	assert.Equal(t, expectedDiskInfo, actualDiskInfo)
}

func TestUdevProbe(t *testing.T) {
	mockOsDiskDetails, err := libudevwrapper.MockDiskDetails()
	if err != nil {
		t.Fatal(err)
	}
	fakeHostName := "node-name"
	fakeNdmClient := CreateFakeClient(t)
	probes := make([]*controller.Probe, 0)
	filters := make([]*controller.Filter, 0)
	nodeAttributes := make(map[string]string)
	nodeAttributes[controller.HostNameKey] = fakeHostName
	mutex := &sync.Mutex{}
	fakeController := &controller.Controller{
		Clientset:      fakeNdmClient,
		Mutex:          mutex,
		Probes:         probes,
		Filters:        filters,
		NodeAttributes: nodeAttributes,
		BDHierarchy:    make(blockdevice.Hierarchy),
	}
	udevprobe := newUdevProbe(fakeController)
	var pi controller.ProbeInterface = udevprobe
	newRegisterProbe := &registerProbe{
		priority:   1,
		name:       "udev probe",
		state:      true,
		pi:         pi,
		controller: fakeController,
	}

	newRegisterProbe.register()

	// Add one filter
	filter := &alwaysTrueFilter{}
	filter1 := &controller.Filter{
		Name:      "filter1",
		State:     true,
		Interface: filter,
	}
	fakeController.AddNewFilter(filter1)
	probeEvent := &ProbeEvent{
		Controller: fakeController,
	}
	eventmsg := make([]*blockdevice.BlockDevice, 0)
	deviceDetails := &blockdevice.BlockDevice{}
	deviceDetails.SysPath = mockOsDiskDetails.SysPath
	eventmsg = append(eventmsg, deviceDetails)
	eventDetails := controller.EventMessage{
		Action:  libudevwrapper.UDEV_ACTION_ADD,
		Devices: eventmsg,
	}
	probeEvent.addBlockDeviceEvent(eventDetails)
	// Retrieve disk resource
	uuid, ok := generateUUID(*deviceDetails)
	cdr1, err1 := fakeController.GetBlockDevice(uuid)
	fakeDr, err := mockOsDiskToAPI()
	if err != nil {
		t.Fatal(err)
	}
	fakeDr.Name = uuid
	fakeDr.ObjectMeta.Labels[controller.KubernetesHostNameLabel] = fakeController.NodeAttributes[controller.HostNameKey]
	fakeDr.ObjectMeta.Labels[controller.NDMDeviceTypeKey] = "blockdevice"
	fakeDr.ObjectMeta.Labels[controller.NDMManagedKey] = controller.TrueString
	tests := map[string]struct {
		actualDisk    *apis.BlockDevice
		expectedDisk  apis.BlockDevice
		actualError   error
		expectedError error
	}{
		"add event for resource with 'fake-disk-uid' uuid": {actualDisk: cdr1, expectedDisk: fakeDr, actualError: err1, expectedError: nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if !ok {
				assert.Nil(t, cdr1)
			} else {
				compareBlockDevice(t, test.expectedDisk, *test.actualDisk)
				assert.Equal(t, test.expectedError, test.actualError)
			}
		})
	}
}

func TestNewUdevProbeForFillDiskDetails(t *testing.T) {
	// Creating the actual udev probe struct
	mockDisk, err := libudevwrapper.MockDiskDetails()
	if err != nil {
		t.Fatal(err)
	}
	sysPath := mockDisk.SysPath
	udev, err := libudevwrapper.NewUdev()
	if err != nil {
		t.Fatal(err)
	}
	actualUdevProbe := &udevProbe{
		udev: udev,
	}
	actualUdevProbe.udevDevice, err = actualUdevProbe.udev.NewDeviceFromSysPath(sysPath)
	if err != nil {
		t.Fatal(err)
	}
	udevProbeError := errors.New("unable to create Udevice object for null struct struct_udev_device")

	// expected cases
	expectedUdevProbe1, expectedError1 := newUdevProbeForFillDiskDetails(sysPath)
	expectedUdevProbe2, expectedError2 := newUdevProbeForFillDiskDetails("")
	tests := map[string]struct {
		actualUdevProbe   *udevProbe
		expectedUdevProbe *udevProbe
		actualError       error
		expectedError     error
	}{
		"udev probe with correct syspath": {actualUdevProbe: actualUdevProbe, expectedUdevProbe: expectedUdevProbe1, actualError: nil, expectedError: expectedError1},
		"udev probe with empty syspath":   {actualUdevProbe: nil, expectedUdevProbe: expectedUdevProbe2, actualError: udevProbeError, expectedError: expectedError2},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedUdevProbe, test.actualUdevProbe)
			assert.Equal(t, test.expectedError, test.actualError)
		})
	}
}

func compareDevLinks(devLink1, devLink2 []blockdevice.DevLink) bool {
	if len(devLink1) != len(devLink2) {
		return false
	}
	cmp := true
	for i := 0; i < len(devLink1); i++ {
		cmp = cmp && unorderedEqual(devLink1[0].Links, devLink2[0].Links)
	}
	return cmp
}

func unorderedEqual(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	exists := make(map[string]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}
