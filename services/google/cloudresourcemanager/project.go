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
package cloudresourcemanager

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Project struct {
	Labels         map[string]string          `json:"labels"`
	LifecycleState *ProjectLifecycleStateEnum `json:"lifecycleState"`
	DisplayName    *string                    `json:"displayname"`
	Parent         *ProjectParent             `json:"parent"`
	Name           *string                    `json:"name"`
	ProjectNumber  *int64                     `json:"projectNumber"`
}

func (r *Project) String() string {
	return dcl.SprintResource(r)
}

// The enum ProjectLifecycleStateEnum.
type ProjectLifecycleStateEnum string

// ProjectLifecycleStateEnumRef returns a *ProjectLifecycleStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ProjectLifecycleStateEnumRef(s string) *ProjectLifecycleStateEnum {
	if s == "" {
		return nil
	}

	v := ProjectLifecycleStateEnum(s)
	return &v
}

func (v ProjectLifecycleStateEnum) Validate() error {
	for _, s := range []string{"LIFECYCLE_STATE_UNSPECIFIED", "ACTIVE", "DELETE_REQUESTED", "DELETE_IN_PROGRESS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ProjectLifecycleStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ProjectParent struct {
	empty bool    `json:"-"`
	Type  *string `json:"type"`
	Id    *string `json:"id"`
}

type jsonProjectParent ProjectParent

func (r *ProjectParent) UnmarshalJSON(data []byte) error {
	var res jsonProjectParent
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyProjectParent
	} else {

		r.Type = res.Type

		r.Id = res.Id

	}
	return nil
}

// This object is used to assert a desired state where this ProjectParent is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyProjectParent *ProjectParent = &ProjectParent{empty: true}

func (r *ProjectParent) Empty() bool {
	return r.empty
}

func (r *ProjectParent) String() string {
	return dcl.SprintResource(r)
}

func (r *ProjectParent) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Project) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloud_resource_manager",
		Type:    "Project",
		Version: "cloudresourcemanager",
	}
}

func (r *Project) ID() (string, error) {
	if err := extractProjectFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"labels":         dcl.ValueOrEmptyString(nr.Labels),
		"lifecycleState": dcl.ValueOrEmptyString(nr.LifecycleState),
		"displayName":    dcl.ValueOrEmptyString(nr.DisplayName),
		"parent":         dcl.ValueOrEmptyString(nr.Parent),
		"name":           dcl.ValueOrEmptyString(nr.Name),
		"projectNumber":  dcl.ValueOrEmptyString(nr.ProjectNumber),
	}
	return dcl.Nprintf("v1/projects/{{name}}", params), nil
}

const ProjectMaxPage = -1

type ProjectList struct {
	Items []*Project

	nextToken string

	pageSize int32

	resource *Project
}

func (l *ProjectList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ProjectList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listProject(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListProject(ctx context.Context, r *Project) (*ProjectList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListProjectWithMaxResults(ctx, r, ProjectMaxPage)

}

func (c *Client) ListProjectWithMaxResults(ctx context.Context, r *Project, pageSize int32) (*ProjectList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listProject(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ProjectList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetProject(ctx context.Context, r *Project) (*Project, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getProjectRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalProject(b, c)
	if err != nil {
		return nil, err
	}
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeProjectNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteProject(ctx context.Context, r *Project) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Project resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Project...")
	deleteOp := deleteProjectOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllProject deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllProject(ctx context.Context, filter func(*Project) bool) error {
	r := &Project{}
	listObj, err := c.ListProject(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllProject(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllProject(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyProject(ctx context.Context, rawDesired *Project, opts ...dcl.ApplyOption) (*Project, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Project
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyProjectHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyProjectHelper(c *Client, ctx context.Context, rawDesired *Project, opts ...dcl.ApplyOption) (*Project, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyProject...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractProjectFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.projectDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToProjectDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []projectApiOperation
	if create {
		ops = append(ops, &createProjectOperation{})
	} else if recreate {
		ops = append(ops, &deleteProjectOperation{})
		ops = append(ops, &createProjectOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeProjectDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetProject(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createProjectOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapProject(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeProjectNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeProjectNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeProjectDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffProject(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
func (r *Project) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	body.WriteString(fmt.Sprintf(`{"options":{"requestedPolicyVersion": %d}}`, r.IAMPolicyVersion()))
	return u, "POST", body, nil
}
