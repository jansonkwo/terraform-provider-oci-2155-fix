// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateExternalDbSystemConnectorDetails The details required to update an external DB system connector.
type UpdateExternalDbSystemConnectorDetails interface {
}

type updateexternaldbsystemconnectordetails struct {
	JsonData      []byte
	ConnectorType string `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *updateexternaldbsystemconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateexternaldbsystemconnectordetails updateexternaldbsystemconnectordetails
	s := struct {
		Model Unmarshalerupdateexternaldbsystemconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateexternaldbsystemconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := UpdateExternalDbSystemMacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updateexternaldbsystemconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateexternaldbsystemconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum
const (
	UpdateExternalDbSystemConnectorDetailsConnectorTypeMacs UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum = "MACS"
)

var mappingUpdateExternalDbSystemConnectorDetailsConnectorTypeEnum = map[string]UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum{
	"MACS": UpdateExternalDbSystemConnectorDetailsConnectorTypeMacs,
}

var mappingUpdateExternalDbSystemConnectorDetailsConnectorTypeEnumLowerCase = map[string]UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum{
	"macs": UpdateExternalDbSystemConnectorDetailsConnectorTypeMacs,
}

// GetUpdateExternalDbSystemConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum
func GetUpdateExternalDbSystemConnectorDetailsConnectorTypeEnumValues() []UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum {
	values := make([]UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingUpdateExternalDbSystemConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExternalDbSystemConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum
func GetUpdateExternalDbSystemConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingUpdateExternalDbSystemConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExternalDbSystemConnectorDetailsConnectorTypeEnum(val string) (UpdateExternalDbSystemConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingUpdateExternalDbSystemConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}