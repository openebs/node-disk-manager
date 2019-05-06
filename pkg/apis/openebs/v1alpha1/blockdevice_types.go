package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

// DeviceSpec defines the desired state of BlockDevice
type DeviceSpec struct {
	Path            string          `json:"path"`                      //Path contain devpath (e.g. /dev/sdb)
	Capacity        DeviceCapacity  `json:"capacity"`                  //Capacity
	Details         DeviceDetails   `json:"details"`                   //Details contains static attributes (model, serial ..)
	DevLinks        []DeviceDevLink `json:"devlinks"`                  //DevLinks contains soft links of one disk
	FileSystem      FileSystemInfo  `json:"filesystem,omitempty"`      //FileSystem contains mountpoint and filesystem type
	Partitioned     string          `json:"partitioned"`               //BlockDevice has partions or not (YES/NO)
	ParentDevice    string          `json:"parentDevice,omitempty"`    //ParentDevice has the UUID of the parent device
	AggregateDevice string          `json:"aggregateDevice,omitempty"` //AggregateDevice has the UUID of the aggregate device created from this device
}

// DeviceCapacity defines the physical and logical size of the block device
type DeviceCapacity struct {
	Storage            uint64 `json:"storage"`            // blockdevice capacity in bytes
	PhysicalSectorSize uint32 `json:"physicalSectorSize"` // blockdevice physical-Sector size in bytes
	LogicalSectorSize  uint32 `json:"logicalSectorSize"`  // blockdevice logical-sector size in bytes
}

// DeviceDetails represent certain hardware/static attributes of the block device
type DeviceDetails struct {
	DeviceType       string `json:"deviceType"`       // DeviceType represents the type of drive like SSD, HDD etc.,
	Model            string `json:"model"`            // Model is model of disk
	Compliance       string `json:"compliance"`       // Implemented standards/specifications version such as SPC-1, SPC-2, etc
	Serial           string `json:"serial"`           // Serial is serial no of disk
	Vendor           string `json:"vendor"`           // Vendor is vendor of disk
	FirmwareRevision string `json:"firmwareRevision"` // disk firmware revision
}

// FileSystemInfo defines the filesystem type and mountpoint of the device if it exists
type FileSystemInfo struct {
	Type       string `json:"fsType,omitempty"`     //Type represents the FileSystem type of the block device
	Mountpoint string `json:"mountPoint,omitempty"` //MountPoint represents the mountpoint of the block device.
}

// DeviceDevLink holds the maping between type and links like by-id type or by-path type link
type DeviceDevLink struct {
	Kind  string   `json:"kind,omitempty"`  // Kind is the type of link like by-id or by-path.
	Links []string `json:"links,omitempty"` // Links are the soft links of Type type
}

// DeviceClaimState defines the observed state of BlockDevice
type DeviceClaimState struct {
	State string `json:"state"` //current claim state of the blockdevice (Claimed/Unclaimed)
}

// DeviceStatus defines the observed state of BlockDevice
type DeviceStatus struct {
	State string `json:"state"` //current state of the blockdevice (Active/Inactive)
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlockDevice is the Schema for the devices API
// +k8s:openapi-gen=true
type BlockDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec       DeviceSpec          `json:"spec,omitempty"`
	Status     DeviceStatus        `json:"status,omitempty"`
	ClaimState DeviceClaimState    `json:"claimState"`
	ClaimRef   *v1.ObjectReference `json:"claimRef,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlockDeviceList contains a list of BlockDevice
type BlockDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlockDevice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BlockDevice{}, &BlockDeviceList{})
}
