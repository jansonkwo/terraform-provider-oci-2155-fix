---
layout: "oci"
page_title: "OCI: oci_core_boot_volume_backups"
sidebar_current: "docs-oci-datasource-core-boot_volume_backups"
description: |-
  Provides a list of BootVolumeBackups
---

# Data Source: oci_core_boot_volume_backups
The `oci_core_boot_volume_backups` data source allows access to the list of OCI boot_volume_backups

Lists the boot volume backups in the specified compartment. You can filter the results by boot volume.


## Example Usage

```hcl
data "oci_core_boot_volume_backups" "test_boot_volume_backups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	boot_volume_id = "${oci_core_boot_volume.test_boot_volume.id}"
	display_name = "${var.boot_volume_backup_display_name}"
	state = "${var.boot_volume_backup_state}"
}
```

## Argument Reference

The following arguments are supported:

* `boot_volume_id` - (Optional) The OCID of the boot volume.
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `boot_volume_backups` - The list of boot_volume_backups.

### BootVolumeBackup Reference

The following attributes are exported:

* `boot_volume_id` - The OCID of the boot volume.
* `compartment_id` - The OCID of the compartment that contains the boot volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the boot volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume backup will expire and be automatically deleted. Format defined by RFC3339. This parameter will always be present for backups that were created automatically by a scheduled-backup policy. For manually created backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the boot volume backup.
* `image_id` - The image OCID used to create the boot volume the backup is taken from.
* `size_in_gbs` - The size of the boot volume, in GBs. 
* `source_type` - Specifies whether the backup was created manually, or via scheduled backup policy.
* `state` - The current state of a boot volume backup.
* `time_created` - The date and time the boot volume backup was created. This is the time the actual point-in-time image of the volume data was taken. Format defined by RFC3339. 
* `time_request_received` - The date and time the request to create the boot volume backup was received. Format defined by RFC3339. 
* `type` - The type of a volume backup. Supported values are 'FULL' or 'INCREMENTAL'.
* `unique_size_in_gbs` - The size used by the backup, in GBs. It is typically smaller than `size_in_gbs`, depending on the space consumed on the boot volume and whether the backup is full or incremental. 
