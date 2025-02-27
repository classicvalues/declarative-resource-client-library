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
package beta

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type LogExclusion struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Filter      *string `json:"filter"`
	Disabled    *bool   `json:"disabled"`
	CreateTime  *string `json:"createTime"`
	UpdateTime  *string `json:"updateTime"`
	Parent      *string `json:"parent"`
}

func (r *LogExclusion) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *LogExclusion) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "logging",
		Type:    "LogExclusion",
		Version: "beta",
	}
}

func (r *LogExclusion) ID() (string, error) {
	if err := extractLogExclusionFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"filter":      dcl.ValueOrEmptyString(nr.Filter),
		"disabled":    dcl.ValueOrEmptyString(nr.Disabled),
		"createTime":  dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":  dcl.ValueOrEmptyString(nr.UpdateTime),
		"parent":      dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.Nprintf("{{parent}}/exclusions/{{name}}", params), nil
}

const LogExclusionMaxPage = -1

type LogExclusionList struct {
	Items []*LogExclusion

	nextToken string

	pageSize int32

	resource *LogExclusion
}

func (l *LogExclusionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *LogExclusionList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listLogExclusion(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListLogExclusion(ctx context.Context, r *LogExclusion) (*LogExclusionList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListLogExclusionWithMaxResults(ctx, r, LogExclusionMaxPage)

}

func (c *Client) ListLogExclusionWithMaxResults(ctx context.Context, r *LogExclusion, pageSize int32) (*LogExclusionList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listLogExclusion(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &LogExclusionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetLogExclusion(ctx context.Context, r *LogExclusion) (*LogExclusion, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getLogExclusionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalLogExclusion(b, c)
	if err != nil {
		return nil, err
	}
	result.Parent = r.Parent
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeLogExclusionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteLogExclusion(ctx context.Context, r *LogExclusion) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("LogExclusion resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting LogExclusion...")
	deleteOp := deleteLogExclusionOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllLogExclusion deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllLogExclusion(ctx context.Context, parent string, filter func(*LogExclusion) bool) error {
	r := &LogExclusion{
		Parent: &parent,
	}
	listObj, err := c.ListLogExclusion(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllLogExclusion(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllLogExclusion(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyLogExclusion(ctx context.Context, rawDesired *LogExclusion, opts ...dcl.ApplyOption) (*LogExclusion, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *LogExclusion
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyLogExclusionHelper(c, ctx, rawDesired, opts...)
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

func applyLogExclusionHelper(c *Client, ctx context.Context, rawDesired *LogExclusion, opts ...dcl.ApplyOption) (*LogExclusion, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyLogExclusion...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractLogExclusionFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.logExclusionDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToLogExclusionDiffs(c.Config, fieldDiffs, opts)
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
	var ops []logExclusionApiOperation
	if create {
		ops = append(ops, &createLogExclusionOperation{})
	} else if recreate {
		ops = append(ops, &deleteLogExclusionOperation{})
		ops = append(ops, &createLogExclusionOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeLogExclusionDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetLogExclusion(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createLogExclusionOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapLogExclusion(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeLogExclusionNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeLogExclusionNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeLogExclusionDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffLogExclusion(c, newDesired, newState)
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
