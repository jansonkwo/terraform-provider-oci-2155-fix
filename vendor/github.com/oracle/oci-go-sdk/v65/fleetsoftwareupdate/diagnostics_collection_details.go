// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiagnosticsCollectionDetails Details to configure diagnostics collection for targets affected by this Exadata Fleet Update Maintenance Cycle.
type DiagnosticsCollectionDetails struct {

	// Enable incident logs and trace collection.
	// Allow Oracle to collect incident logs and traces to enable fault diagnosis and issue resolution according to the selected mode.
	LogCollectionMode DataCollectionModesEnum `mandatory:"false" json:"logCollectionMode,omitempty"`
}

func (m DiagnosticsCollectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiagnosticsCollectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataCollectionModesEnum(string(m.LogCollectionMode)); !ok && m.LogCollectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogCollectionMode: %s. Supported values are: %s.", m.LogCollectionMode, strings.Join(GetDataCollectionModesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}