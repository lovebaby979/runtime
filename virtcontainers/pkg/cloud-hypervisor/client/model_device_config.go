/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DeviceConfig struct for DeviceConfig
type DeviceConfig struct {
	Path string `json:"path"`
	Iommu bool `json:"iommu,omitempty"`
	Id string `json:"id,omitempty"`
}
