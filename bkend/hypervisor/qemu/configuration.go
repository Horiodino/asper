package qemu

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

type QEMU interface {
}

type QEMUClient struct {
	Resource_name string
	Resource_type string
}

func InitializeQEMU() *QEMUClient {
	return &QEMUClient{}
}

type Domain libvirtxml.Domain

type DomainGenID libvirtxml.DomainGenID
type DomainMetadata libvirtxml.DomainMetadata
type DomainMaxMemory libvirtxml.DomainMaxMemory
type DomainMemory libvirtxml.DomainMemory
type DomainCurrentMemory libvirtxml.DomainCurrentMemory
type DomainBlockIOTune libvirtxml.DomainBlockIOTune
type DomainMemoryTune libvirtxml.DomainMemoryTune
type DomainMemoryBacking libvirtxml.DomainMemoryBacking
type DomainVCPU libvirtxml.DomainVCPU
type DomainVCPUs libvirtxml.DomainVCPUs
type DomainIOThreadIDs libvirtxml.DomainIOThreadIDs
type DomainCPUTune libvirtxml.DomainCPUTune
type DomainNUMATune libvirtxml.DomainNUMATune
type DomainResource libvirtxml.DomainResource
type DomainSysInfo []libvirtxml.DomainSysInfo
type DomainOS libvirtxml.DomainOS
type DomainIDMap libvirtxml.DomainIDMap
type DomainFeatureList libvirtxml.DomainFeatureList
type DomainCPU libvirtxml.DomainCPU
type DomainClock libvirtxml.DomainClock
type DomainPM libvirtxml.DomainPM
type DomainPerf libvirtxml.DomainPerf
type DomainDeviceList libvirtxml.DomainDeviceList
type DomainSecLabel []libvirtxml.DomainSecLabel
type DomainKeyWrap libvirtxml.DomainKeyWrap
type DomainLaunchSecurity libvirtxml.DomainLaunchSecurity
	
	
type NetworkMetadata libvirtxml.NetworkMetadata    
type NetworkForward libvirtxml.NetworkForward     
type NetworkBridge libvirtxml.NetworkBridge      
type NetworkMTU  libvirtxml.NetworkMTU         
type NetworkMAC libvirtxml.NetworkMAC         
type NetworkDomain libvirtxml.NetworkDomain      
type NetworkDNS libvirtxml.NetworkDNS         
type NetworkVLAN libvirtxml.NetworkVLAN     
type NetworkBandwidth libvirtxml.NetworkBandwidth   
type NetworkPortOptions  libvirtxml.NetworkPortOptions 
type NetworkIP  []libvirtxml.NetworkIP         
type NetworkRoute    []libvirtxml.NetworkRoute      
type NetworkVirtualPort libvirtxml.NetworkVirtualPort 
type NetworkPortGroup []libvirtxml.NetworkPortGroup  


// ////////////////////////////////////////////////////////////////////////////////////
type DomainAddress libvirtxml.DomainAddress

// the domain controller is for the USB device passthrough
type DomainController libvirtxml.DomainController

// for disk creation
type StoragePool libvirtxml.StoragePool

// for network creation
type Network libvirtxml.Network

// for volume creation
type StorageVolume libvirtxml.StorageVolume

// for snapshot creation
type DomainSnapshot libvirtxml.DomainSnapshot

// list of domains
type DomainUpdate libvirtxml.DomainDeviceList

// list of storage pools

func (c *QEMUClient) ReturnDomain() *Domain {
	return &Domain{}
}
