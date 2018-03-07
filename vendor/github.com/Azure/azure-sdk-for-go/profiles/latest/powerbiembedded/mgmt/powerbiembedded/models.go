// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package powerbiembedded

import original "github.com/Azure/azure-sdk-for-go/services/powerbiembedded/mgmt/2016-01-29/powerbiembedded"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type AccessKeyName = original.AccessKeyName

const (
	Key1 AccessKeyName = original.Key1
	Key2 AccessKeyName = original.Key2
)

type CheckNameReason = original.CheckNameReason

const (
	Invalid     CheckNameReason = original.Invalid
	Unavailable CheckNameReason = original.Unavailable
)

type AzureSku = original.AzureSku
type CheckNameRequest = original.CheckNameRequest
type CheckNameResponse = original.CheckNameResponse
type CreateWorkspaceCollectionRequest = original.CreateWorkspaceCollectionRequest
type Display = original.Display
type Error = original.Error
type ErrorDetail = original.ErrorDetail
type MigrateWorkspaceCollectionRequest = original.MigrateWorkspaceCollectionRequest
type Operation = original.Operation
type OperationList = original.OperationList
type UpdateWorkspaceCollectionRequest = original.UpdateWorkspaceCollectionRequest
type Workspace = original.Workspace
type WorkspaceCollection = original.WorkspaceCollection
type WorkspaceCollectionAccessKey = original.WorkspaceCollectionAccessKey
type WorkspaceCollectionAccessKeys = original.WorkspaceCollectionAccessKeys
type WorkspaceCollectionList = original.WorkspaceCollectionList
type WorkspaceCollectionsDeleteFuture = original.WorkspaceCollectionsDeleteFuture
type WorkspaceList = original.WorkspaceList

func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}

type WorkspaceCollectionsClient = original.WorkspaceCollectionsClient

func NewWorkspaceCollectionsClient(subscriptionID string) WorkspaceCollectionsClient {
	return original.NewWorkspaceCollectionsClient(subscriptionID)
}
func NewWorkspaceCollectionsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceCollectionsClient {
	return original.NewWorkspaceCollectionsClientWithBaseURI(baseURI, subscriptionID)
}

type WorkspacesClient = original.WorkspacesClient

func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
