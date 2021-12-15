// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// CreateWaitStageDetails Specifies the Wait stage. You can specify variable wait times or an absolute duration.
type CreateWaitStageDetails struct {

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"true" json:"buildPipelineStagePredecessorCollection"`

	WaitCriteria CreateWaitCriteriaDetails `mandatory:"true" json:"waitCriteria"`

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the stage.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateWaitStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m CreateWaitStageDetails) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineId returns BuildPipelineId
func (m CreateWaitStageDetails) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m CreateWaitStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateWaitStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateWaitStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateWaitStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateWaitStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateWaitStageDetails CreateWaitStageDetails
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeCreateWaitStageDetails
	}{
		"WAIT",
		(MarshalTypeCreateWaitStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateWaitStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                             *string                                  `json:"displayName"`
		Description                             *string                                  `json:"description"`
		FreeformTags                            map[string]string                        `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}        `json:"definedTags"`
		BuildPipelineId                         *string                                  `json:"buildPipelineId"`
		BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `json:"buildPipelineStagePredecessorCollection"`
		WaitCriteria                            createwaitcriteriadetails                `json:"waitCriteria"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.BuildPipelineId = model.BuildPipelineId

	m.BuildPipelineStagePredecessorCollection = model.BuildPipelineStagePredecessorCollection

	nn, e = model.WaitCriteria.UnmarshalPolymorphicJSON(model.WaitCriteria.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WaitCriteria = nn.(CreateWaitCriteriaDetails)
	} else {
		m.WaitCriteria = nil
	}

	return
}