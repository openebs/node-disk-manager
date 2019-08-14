/*
Copyright 2019 The OpenEBS Authors

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	Path             string         `json:"path"`                       //Path contain devpath (e.g. /dev/sdb)
	Capacity         DiskCapacity   `json:"capacity"`                   //Capacity
	Details          DiskDetails    `json:"details"`                    //Details contains static attributes (model, serial ..)
	DevLinks         []DiskDevLink  `json:"devlinks"`                   //DevLinks contains soft links of one disk
	FileSystem       FileSystemInfo `json:"fileSystem,omitempty"`       //Contains the data about filesystem on the disk
	PartitionDetails []Partition    `json:"partitionDetails,omitempty"` //Details of partitions in the disk (filesystem, partition type)
}

type DiskCapacity struct {
	Storage            uint64 `json:"storage"`            // disk size in bytes
	PhysicalSectorSize uint32 `json:"physicalSectorSize"` // disk physical-Sector size in bytes
	LogicalSectorSize  uint32 `json:"logicalSectorSize"`  // disk logical-sector size in bytes
}

type DiskDetails struct {
	RotationRate     uint16 `json:"rotationRate"`     // Disk rotation speed if disk is not SSD
	DriveType        string `json:"driveType"`        // DriveType represents the type of drive like SSD, HDD etc.,
	Model            string `json:"model"`            // Model is model of disk
	Compliance       string `json:"compliance"`       // Implemented standards/specifications version such as SPC-1, SPC-2, etc
	Serial           string `json:"serial"`           // Serial is serial no of disk
	Vendor           string `json:"vendor"`           // Vendor is vendor of disk
	FirmwareRevision string `json:"firmwareRevision"` // disk firmware revision
}

// DiskDevlink holds the maping between type and links like by-id type or by-path type link
type DiskDevLink struct {
	Kind  string   `json:"kind"`  // Kind is the type of link like by-id or by-path.
	Links []string `json:"links"` // Links are the soft links of Type type
}

// Partition represents the partition information of the disk
type Partition struct {
	PartitionType string         `json:"partitionType"`
	FileSystem    FileSystemInfo `json:"fileSystem,omitempty"`
}

// DiskStatus defines the observed state of Disk
type DiskStatus struct {
	State DiskState `json:"state"` //current state of the disk (Active/Inactive)
}

// DiskState defines the observed state of the disk
type DiskState string

const (
	// DiskActive is the state for a physical disk that is connected to the node
	DiskActive DiskState = "Active"
	// DiskInactive is the state for a physical disk that is disconnected from a node
	DiskInactive DiskState = "Inactive"
	// DiskUnknown is the state for a physical disk whose state (attached/detached) cannot
	// be determined at this time.
	DiskUnknown DiskState = "Unknown"
)

type Temperature struct {
	CurrentTemperature int16 `json:"currentTemperature"`
	HighestTemperature int16 `json:"highestTemperature"`
	LowestTemperature  int16 `json:"lowestTemperature"`
}

type DiskStat struct {
	TempInfo              Temperature `json:"diskTemperature"`
	TotalBytesRead        uint64      `json:"totalBytesRead"`
	TotalBytesWritten     uint64      `json:"totalBytesWritten"`
	DeviceUtilizationRate float64     `json:"deviceUtilizationRate"`
	PercentEnduranceUsed  float64     `json:"percentEnduranceUsed"`
}

type DeviceInfo struct {
	DeviceUID string `json:"blockDeviceUID"` //Cross reference to BlockDevice CR backed by this disk
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Disk is the Schema for the disks API
// +k8s:openapi-gen=true
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiskSpec   `json:"spec,omitempty"`
	Status DiskStatus `json:"status,omitempty"`
	Stats  DiskStat   `json:"stats,omitempty"`
	Device DeviceInfo `json:"deviceInfo"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiskList contains a list of Disk
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
