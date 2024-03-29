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
	"github.com/openebs/node-disk-manager/blockdevice"
	"github.com/openebs/node-disk-manager/cmd/ndm_daemonset/controller"
	"github.com/openebs/node-disk-manager/pkg/seachest"
	"github.com/openebs/node-disk-manager/pkg/util"
	"k8s.io/klog/v2"
)

// seachest contains required variables for populating diskInfo
type seachestProbe struct {
	// Every new probe needs a controller object to register itself.
	// Here Controller consists of Clientset, kubeClientset, probes, etc which is used to
	// create, update, delete, deactivate the disk resources or list the probes already registered.
	Controller         *controller.Controller
	SeachestIdentifier *seachest.Identifier
}

const (
	seachestConfigKey     = "seachest-probe"
	seachestProbePriority = 6
)

var (
	seachestProbeName  = "seachest probe"
	seachestProbeState = defaultEnabled
)

// init is used to get a controller object and then register itself
var seachestProbeRegister = func() {
	// Get a controller object
	ctrl := <-controller.ControllerBroadcastChannel
	if ctrl == nil {
		klog.Error("unable to configure", seachestProbeName)
		return
	}
	if ctrl.NDMConfig != nil {
		for _, probeConfig := range ctrl.NDMConfig.ProbeConfigs {
			if probeConfig.Key == seachestConfigKey {
				seachestProbeName = probeConfig.Name
				seachestProbeState = util.CheckTruthy(probeConfig.State)
				break
			}
		}
	}
	newRegisterProbe := &registerProbe{
		priority:   seachestProbePriority,
		name:       seachestProbeName,
		state:      seachestProbeState,
		pi:         &seachestProbe{Controller: ctrl},
		controller: ctrl,
	}
	// Here we register the probe (seachest probe in this case)
	newRegisterProbe.register()
}

// newSeachestProbe returns seachestProbe struct which helps populate diskInfo struct
// with the basic disk details such as logical size, firmware revision, etc
func newSeachestProbe(devPath string) *seachestProbe {
	seachestIdentifier := &seachest.Identifier{
		DevPath: devPath,
	}
	seachestProbe := &seachestProbe{
		SeachestIdentifier: seachestIdentifier,
	}
	return seachestProbe
}

// Start is mainly used for one time activities such as monitoring.
// It is a part of probe interface but here we does not require to perform
// such activities, hence empty implementation
func (scp *seachestProbe) Start() {}

// fillDiskDetails fills details in diskInfo struct using information it gets from probe
func (scp *seachestProbe) FillBlockDeviceDetails(blockDevice *blockdevice.BlockDevice) {
	if blockDevice.DevPath == "" {
		klog.Error("seachestIdentifier is found empty, seachest probe will not fill disk details.")
		return
	}

	seachestProbe := newSeachestProbe(blockDevice.DevPath)
	driveInfo, err := seachestProbe.SeachestIdentifier.SeachestBasicDiskInfo()
	if err != 0 {
		klog.Error(err)
		return
	}

	if blockDevice.DeviceAttributes.FirmwareRevision == "" {
		blockDevice.DeviceAttributes.FirmwareRevision = seachestProbe.SeachestIdentifier.GetFirmwareRevision(driveInfo)
		klog.V(4).Infof("Disk: %s FirmwareRevision:%s filled by seachest.", blockDevice.DevPath, blockDevice.DeviceAttributes.FirmwareRevision)
	}

	if blockDevice.DeviceAttributes.LogicalBlockSize == 0 {
		blockDevice.DeviceAttributes.LogicalBlockSize = seachestProbe.SeachestIdentifier.GetLogicalSectorSize(driveInfo)
		klog.V(4).Infof("Disk: %s LogicalBlockSize:%d filled by seachest.", blockDevice.DevPath, blockDevice.DeviceAttributes.LogicalBlockSize)
	}

	if blockDevice.DeviceAttributes.PhysicalBlockSize == 0 {
		blockDevice.DeviceAttributes.PhysicalBlockSize = seachestProbe.SeachestIdentifier.GetPhysicalSectorSize(driveInfo)
		klog.V(4).Infof("Disk: %s PhysicalBlockSize:%d filled by seachest.", blockDevice.DevPath, blockDevice.DeviceAttributes.PhysicalBlockSize)
	}

	if blockDevice.DeviceAttributes.DriveType == "" ||
		blockDevice.DeviceAttributes.DriveType == blockdevice.DriveTypeUnknown {
		blockDevice.DeviceAttributes.DriveType = seachestProbe.SeachestIdentifier.DriveType(driveInfo)
		klog.V(4).Infof("Disk: %s DriveType:%s filled by seachest.", blockDevice.DevPath, blockDevice.DeviceAttributes.DriveType)
	}

	// All the below mentioned fields will be filled in only after BlockDevice struct
	// starts supporting them.
	if blockDevice.SMARTInfo.RotationRate == 0 {
		blockDevice.SMARTInfo.RotationRate = seachestProbe.SeachestIdentifier.GetRotationRate(driveInfo)
		klog.V(4).Infof("Disk: %s RotationRate:%d filled by seachest.", blockDevice.DevPath, blockDevice.SMARTInfo.RotationRate)
	}

	if blockDevice.SMARTInfo.TotalBytesRead == 0 {
		blockDevice.SMARTInfo.TotalBytesRead = seachestProbe.SeachestIdentifier.GetTotalBytesRead(driveInfo)
		klog.V(4).Infof("Disk: %s TotalBytesRead:%d filled by seachest.", blockDevice.DevPath, blockDevice.SMARTInfo.TotalBytesRead)
	}

	if blockDevice.SMARTInfo.TotalBytesWritten == 0 {
		blockDevice.SMARTInfo.TotalBytesWritten = seachestProbe.SeachestIdentifier.GetTotalBytesWritten(driveInfo)
		klog.V(4).Infof("Disk: %s TotalBytesWritten:%d filled by seachest.", blockDevice.DevPath, blockDevice.SMARTInfo.TotalBytesWritten)
	}

	if blockDevice.SMARTInfo.UtilizationRate == 0 {
		blockDevice.SMARTInfo.UtilizationRate = seachestProbe.SeachestIdentifier.GetDeviceUtilizationRate(driveInfo)
		klog.V(4).Infof("Disk: %s UtilizationRate:%f filled by seachest.", blockDevice.DevPath, blockDevice.SMARTInfo.UtilizationRate)
	}

	if blockDevice.SMARTInfo.PercentEnduranceUsed == 0 {
		blockDevice.SMARTInfo.PercentEnduranceUsed = seachestProbe.SeachestIdentifier.GetPercentEnduranceUsed(driveInfo)
		klog.V(4).Infof("Disk: %s PercentEnduranceUsed:%f filled by seachest.", blockDevice.DevPath, blockDevice.SMARTInfo.PercentEnduranceUsed)
	}

	blockDevice.SMARTInfo.TemperatureInfo.CurrentTemperatureDataValid = seachestProbe.
		SeachestIdentifier.GetTemperatureDataValidStatus(driveInfo)

	klog.V(4).Infof("Disk: %s TemperatureDataValid:%t filled by seachest.",
		blockDevice.DevPath, blockDevice.SMARTInfo.TemperatureInfo.CurrentTemperatureDataValid)

	if blockDevice.SMARTInfo.TemperatureInfo.CurrentTemperatureDataValid == true {
		blockDevice.SMARTInfo.TemperatureInfo.CurrentTemperature = seachestProbe.
			SeachestIdentifier.GetCurrentTemperature(driveInfo)

		klog.V(4).Infof("Disk: %s CurrentTemperature:%d filled by seachest.",
			blockDevice.DevPath, blockDevice.SMARTInfo.TemperatureInfo.CurrentTemperature)

	}

	blockDevice.SMARTInfo.TemperatureInfo.HighestTemperatureDataValid = seachestProbe.
		SeachestIdentifier.GetHighestValid(driveInfo)

	klog.V(4).Infof("Disk: %s HighestTemperatureDataValid:%t filled by seachest.",
		blockDevice.DevPath, blockDevice.SMARTInfo.TemperatureInfo.HighestTemperatureDataValid)

	if blockDevice.SMARTInfo.TemperatureInfo.HighestTemperatureDataValid == true {

		blockDevice.SMARTInfo.TemperatureInfo.HighestTemperature = seachestProbe.
			SeachestIdentifier.GetHighestTemperature(driveInfo)

		klog.V(4).Infof("Disk: %s HighestTemperature:%d filled by seachest.",
			blockDevice.DevPath, blockDevice.SMARTInfo.TemperatureInfo.HighestTemperature)
	}
	blockDevice.SMARTInfo.TemperatureInfo.LowestTemperatureDataValid = seachestProbe.
		SeachestIdentifier.GetLowestValid(driveInfo)

	klog.V(4).Infof("Disk: %s LowestValid:%t filled by seachest.",
		blockDevice.DevPath, blockDevice.SMARTInfo.TemperatureInfo.LowestTemperatureDataValid)

	if blockDevice.SMARTInfo.TemperatureInfo.LowestTemperatureDataValid == true {
		blockDevice.SMARTInfo.TemperatureInfo.LowestTemperature = seachestProbe.
			SeachestIdentifier.GetLowestTemperature(driveInfo)

		klog.V(4).Infof("Disk: %s LowestTemperature:%d filled by seachest.",
			blockDevice.DevPath, blockDevice.SMARTInfo.TemperatureInfo.LowestTemperature)
	}
}
