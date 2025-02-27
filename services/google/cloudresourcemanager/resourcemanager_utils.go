// Copyright 2021 Google LLC. All Rights Reserved.
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
// Package cloudresourcemanager contains support code for the CRM service.
package cloudresourcemanager

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// The project is already effectively deleted if it's in DELETE_REQUESTED state.
func projectDeletePrecondition(r *Project) bool {
	return *r.LifecycleState == *ProjectLifecycleStateEnumRef("DELETE_REQUESTED")
}

func (r *Folder) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("folders?parent={{parent}}", "https://cloudresourcemanager.googleapis.com/v2", userBasePath, params), nil

}

func (r *Folder) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("folders/{{name}}", "https://cloudresourcemanager.googleapis.com/v2", userBasePath, params), nil
}

func (r *Folder) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("folders?parent={{parent}}", "https://cloudresourcemanager.googleapis.com/v2", userBasePath, params), nil
}

func (r *Folder) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "MoveFolder" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("folders/{{name}}:move", "https://cloudresourcemanager.googleapis.com/v2", userBasePath, fields), nil

	}
	if updateName == "UpdateFolder" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("folders/{{name}}?updateMask=displayName", "https://cloudresourcemanager.googleapis.com/v2", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

func (r *Folder) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("folders/{{name}}", "https://cloudresourcemanager.googleapis.com/v2", userBasePath, params), nil
}

func (op *updateFolderMoveFolderOperation) do(ctx context.Context, r *Folder, c *Client) error {
	_, err := c.GetFolder(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "MoveFolder")
	if err != nil {
		return err
	}

	req, err := newUpdateFolderMoveFolderRequest(ctx, r, c)
	if err != nil {
		return err
	}

	if p, ok := req["parent"]; ok {
		req["destinationParent"] = p
		delete(req, "parent")
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFolderMoveFolderRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://cloudresourcemanager.googleapis.com/v1", "GET")

	if err != nil {
		return err
	}

	return nil
}
