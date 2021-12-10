// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v53/common"
)

// JmesPathDimensionValue Evaluated type of dimension value.
type JmesPathDimensionValue struct {

	// The location to use for deriving the dimension value (evaluated).
	// The path must start with `logContent` in an acceptable notation style with supported JMESPath selectors (https://jmespath.org/specification.html): expression with dot and index operator (`.`, and `MetricDataDetails.
	// The returned value depends on the results of evaluation.
	// If the evaluated value is valid, then the evaluated value is returned without double quotes. (Any front or trailing double quotes are trimmed before returning the value. For example, the evaluated value `"compartmentId"` is returned as `compartmentId`.)
	// If the evaluated value is invalid, then the returned value is `SCH_EVAL_INVALID_VALUE`.
	// If the evaluated value is empty, then the returned value is `SCH_EVAL_VALUE_EMPTY`.
	Path *string `mandatory:"true" json:"path"`
}

func (m JmesPathDimensionValue) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m JmesPathDimensionValue) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJmesPathDimensionValue JmesPathDimensionValue
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeJmesPathDimensionValue
	}{
		"jmesPath",
		(MarshalTypeJmesPathDimensionValue)(m),
	}

	return json.Marshal(&s)
}
