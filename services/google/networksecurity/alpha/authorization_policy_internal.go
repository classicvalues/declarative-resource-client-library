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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *AuthorizationPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "action"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *AuthorizationPolicyRules) validate() error {
	return nil
}
func (r *AuthorizationPolicyRulesSources) validate() error {
	return nil
}
func (r *AuthorizationPolicyRulesDestinations) validate() error {
	if err := dcl.Required(r, "hosts"); err != nil {
		return err
	}
	if err := dcl.Required(r, "ports"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.HttpHeaderMatch) {
		if err := r.HttpHeaderMatch.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) validate() error {
	if err := dcl.Required(r, "headerName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "regexMatch"); err != nil {
		return err
	}
	return nil
}
func (r *AuthorizationPolicy) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networksecurity.googleapis.com/v1alpha1/", params)
}

func (r *AuthorizationPolicy) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AuthorizationPolicy) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies", nr.basePath(), userBasePath, params), nil

}

func (r *AuthorizationPolicy) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies?authorizationPolicyId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *AuthorizationPolicy) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AuthorizationPolicy) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *AuthorizationPolicy) SetPolicyVerb() string {
	return "POST"
}

func (r *AuthorizationPolicy) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *AuthorizationPolicy) IAMPolicyVersion() int {
	return 3
}

// authorizationPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type authorizationPolicyApiOperation interface {
	do(context.Context, *AuthorizationPolicy, *Client) error
}

// newUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest creates a request for an
// AuthorizationPolicy resource's UpdateAuthorizationPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(ctx context.Context, f *AuthorizationPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		req["action"] = v
	}
	if v, err := expandAuthorizationPolicyRulesSlice(c, f.Rules); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["rules"] = v
	}
	return req, nil
}

// marshalUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAuthorizationPolicyUpdateAuthorizationPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAuthorizationPolicyUpdateAuthorizationPolicyOperation) do(ctx context.Context, r *AuthorizationPolicy, c *Client) error {
	_, err := c.GetAuthorizationPolicy(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAuthorizationPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAuthorizationPolicyRaw(ctx context.Context, r *AuthorizationPolicy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AuthorizationPolicyMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listAuthorizationPolicyOperation struct {
	AuthorizationPolicies []map[string]interface{} `json:"authorizationPolicies"`
	Token                 string                   `json:"nextPageToken"`
}

func (c *Client) listAuthorizationPolicy(ctx context.Context, r *AuthorizationPolicy, pageToken string, pageSize int32) ([]*AuthorizationPolicy, string, error) {
	b, err := c.listAuthorizationPolicyRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAuthorizationPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AuthorizationPolicy
	for _, v := range m.AuthorizationPolicies {
		res, err := unmarshalMapAuthorizationPolicy(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAuthorizationPolicy(ctx context.Context, f func(*AuthorizationPolicy) bool, resources []*AuthorizationPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAuthorizationPolicy(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteAuthorizationPolicyOperation struct{}

func (op *deleteAuthorizationPolicyOperation) do(ctx context.Context, r *AuthorizationPolicy, c *Client) error {
	r, err := c.GetAuthorizationPolicy(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "AuthorizationPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetAuthorizationPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAuthorizationPolicy(ctx, r)
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createAuthorizationPolicyOperation struct {
	response map[string]interface{}
}

func (op *createAuthorizationPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAuthorizationPolicyOperation) do(ctx context.Context, r *AuthorizationPolicy, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetAuthorizationPolicy(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAuthorizationPolicyRaw(ctx context.Context, r *AuthorizationPolicy) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) authorizationPolicyDiffsForRawDesired(ctx context.Context, rawDesired *AuthorizationPolicy, opts ...dcl.ApplyOption) (initial, desired *AuthorizationPolicy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AuthorizationPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AuthorizationPolicy); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected AuthorizationPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAuthorizationPolicy(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a AuthorizationPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AuthorizationPolicy resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that AuthorizationPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAuthorizationPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for AuthorizationPolicy: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for AuthorizationPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAuthorizationPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for AuthorizationPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAuthorizationPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for AuthorizationPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAuthorizationPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAuthorizationPolicyInitialState(rawInitial, rawDesired *AuthorizationPolicy) (*AuthorizationPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAuthorizationPolicyDesiredState(rawDesired, rawInitial *AuthorizationPolicy, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &AuthorizationPolicy{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.Action) {
		canonicalDesired.Action = rawInitial.Action
	} else {
		canonicalDesired.Action = rawDesired.Action
	}
	if dcl.IsZeroValue(rawDesired.Rules) {
		canonicalDesired.Rules = rawInitial.Rules
	} else {
		canonicalDesired.Rules = rawDesired.Rules
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeAuthorizationPolicyNewState(c *Client, rawNew, rawDesired *AuthorizationPolicy) (*AuthorizationPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Action) && dcl.IsEmptyValueIndirect(rawDesired.Action) {
		rawNew.Action = rawDesired.Action
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Rules) && dcl.IsEmptyValueIndirect(rawDesired.Rules) {
		rawNew.Rules = rawDesired.Rules
	} else {
		rawNew.Rules = canonicalizeNewAuthorizationPolicyRulesSlice(c, rawDesired.Rules, rawNew.Rules)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeAuthorizationPolicyRules(des, initial *AuthorizationPolicyRules, opts ...dcl.ApplyOption) *AuthorizationPolicyRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AuthorizationPolicyRules{}

	if dcl.IsZeroValue(des.Sources) {
		des.Sources = initial.Sources
	} else {
		cDes.Sources = des.Sources
	}
	if dcl.IsZeroValue(des.Destinations) {
		des.Destinations = initial.Destinations
	} else {
		cDes.Destinations = des.Destinations
	}

	return cDes
}

func canonicalizeNewAuthorizationPolicyRules(c *Client, des, nw *AuthorizationPolicyRules) *AuthorizationPolicyRules {
	if des == nil || nw == nil {
		return nw
	}

	nw.Sources = canonicalizeNewAuthorizationPolicyRulesSourcesSlice(c, des.Sources, nw.Sources)
	nw.Destinations = canonicalizeNewAuthorizationPolicyRulesDestinationsSlice(c, des.Destinations, nw.Destinations)

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesSet(c *Client, des, nw []AuthorizationPolicyRules) []AuthorizationPolicyRules {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRules
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAuthorizationPolicyRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewAuthorizationPolicyRulesSlice(c *Client, des, nw []AuthorizationPolicyRules) []AuthorizationPolicyRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AuthorizationPolicyRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRules(c, &d, &n))
	}

	return items
}

func canonicalizeAuthorizationPolicyRulesSources(des, initial *AuthorizationPolicyRulesSources, opts ...dcl.ApplyOption) *AuthorizationPolicyRulesSources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AuthorizationPolicyRulesSources{}

	if dcl.IsZeroValue(des.Principals) {
		des.Principals = initial.Principals
	} else {
		cDes.Principals = des.Principals
	}
	if dcl.IsZeroValue(des.IPBlocks) {
		des.IPBlocks = initial.IPBlocks
	} else {
		cDes.IPBlocks = des.IPBlocks
	}

	return cDes
}

func canonicalizeNewAuthorizationPolicyRulesSources(c *Client, des, nw *AuthorizationPolicyRulesSources) *AuthorizationPolicyRulesSources {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesSourcesSet(c *Client, des, nw []AuthorizationPolicyRulesSources) []AuthorizationPolicyRulesSources {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRulesSources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAuthorizationPolicyRulesSourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewAuthorizationPolicyRulesSourcesSlice(c *Client, des, nw []AuthorizationPolicyRulesSources) []AuthorizationPolicyRulesSources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AuthorizationPolicyRulesSources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRulesSources(c, &d, &n))
	}

	return items
}

func canonicalizeAuthorizationPolicyRulesDestinations(des, initial *AuthorizationPolicyRulesDestinations, opts ...dcl.ApplyOption) *AuthorizationPolicyRulesDestinations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AuthorizationPolicyRulesDestinations{}

	if dcl.IsZeroValue(des.Hosts) {
		des.Hosts = initial.Hosts
	} else {
		cDes.Hosts = des.Hosts
	}
	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	} else {
		cDes.Ports = des.Ports
	}
	if dcl.IsZeroValue(des.Methods) {
		des.Methods = initial.Methods
	} else {
		cDes.Methods = des.Methods
	}
	cDes.HttpHeaderMatch = canonicalizeAuthorizationPolicyRulesDestinationsHttpHeaderMatch(des.HttpHeaderMatch, initial.HttpHeaderMatch, opts...)

	return cDes
}

func canonicalizeNewAuthorizationPolicyRulesDestinations(c *Client, des, nw *AuthorizationPolicyRulesDestinations) *AuthorizationPolicyRulesDestinations {
	if des == nil || nw == nil {
		return nw
	}

	nw.HttpHeaderMatch = canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, des.HttpHeaderMatch, nw.HttpHeaderMatch)

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsSet(c *Client, des, nw []AuthorizationPolicyRulesDestinations) []AuthorizationPolicyRulesDestinations {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRulesDestinations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAuthorizationPolicyRulesDestinationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsSlice(c *Client, des, nw []AuthorizationPolicyRulesDestinations) []AuthorizationPolicyRulesDestinations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AuthorizationPolicyRulesDestinations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRulesDestinations(c, &d, &n))
	}

	return items
}

func canonicalizeAuthorizationPolicyRulesDestinationsHttpHeaderMatch(des, initial *AuthorizationPolicyRulesDestinationsHttpHeaderMatch, opts ...dcl.ApplyOption) *AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}

	if dcl.StringCanonicalize(des.HeaderName, initial.HeaderName) || dcl.IsZeroValue(des.HeaderName) {
		cDes.HeaderName = initial.HeaderName
	} else {
		cDes.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}

	return cDes
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, des, nw *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) *AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HeaderName, nw.HeaderName) {
		nw.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatchSet(c *Client, des, nw []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) []AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRulesDestinationsHttpHeaderMatch
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAuthorizationPolicyRulesDestinationsHttpHeaderMatchNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, des, nw []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) []AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AuthorizationPolicyRulesDestinationsHttpHeaderMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffAuthorizationPolicy(c *Client, desired, actual *AuthorizationPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.Info{ObjectFunction: compareAuthorizationPolicyRulesNewStyle, EmptyObject: EmptyAuthorizationPolicyRules, OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareAuthorizationPolicyRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AuthorizationPolicyRules)
	if !ok {
		desiredNotPointer, ok := d.(AuthorizationPolicyRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRules or *AuthorizationPolicyRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AuthorizationPolicyRules)
	if !ok {
		actualNotPointer, ok := a.(AuthorizationPolicyRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Sources, actual.Sources, dcl.Info{ObjectFunction: compareAuthorizationPolicyRulesSourcesNewStyle, EmptyObject: EmptyAuthorizationPolicyRulesSources, OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Sources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Destinations, actual.Destinations, dcl.Info{ObjectFunction: compareAuthorizationPolicyRulesDestinationsNewStyle, EmptyObject: EmptyAuthorizationPolicyRulesDestinations, OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Destinations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAuthorizationPolicyRulesSourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AuthorizationPolicyRulesSources)
	if !ok {
		desiredNotPointer, ok := d.(AuthorizationPolicyRulesSources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRulesSources or *AuthorizationPolicyRulesSources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AuthorizationPolicyRulesSources)
	if !ok {
		actualNotPointer, ok := a.(AuthorizationPolicyRulesSources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRulesSources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Principals, actual.Principals, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Principals")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPBlocks, actual.IPBlocks, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("IpBlocks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAuthorizationPolicyRulesDestinationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AuthorizationPolicyRulesDestinations)
	if !ok {
		desiredNotPointer, ok := d.(AuthorizationPolicyRulesDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRulesDestinations or *AuthorizationPolicyRulesDestinations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AuthorizationPolicyRulesDestinations)
	if !ok {
		actualNotPointer, ok := a.(AuthorizationPolicyRulesDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRulesDestinations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Hosts, actual.Hosts, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Hosts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Methods, actual.Methods, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("Methods")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaderMatch, actual.HttpHeaderMatch, dcl.Info{ObjectFunction: compareAuthorizationPolicyRulesDestinationsHttpHeaderMatchNewStyle, EmptyObject: EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch, OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("HttpHeaderMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAuthorizationPolicyRulesDestinationsHttpHeaderMatchNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AuthorizationPolicyRulesDestinationsHttpHeaderMatch)
	if !ok {
		desiredNotPointer, ok := d.(AuthorizationPolicyRulesDestinationsHttpHeaderMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRulesDestinationsHttpHeaderMatch or *AuthorizationPolicyRulesDestinationsHttpHeaderMatch", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AuthorizationPolicyRulesDestinationsHttpHeaderMatch)
	if !ok {
		actualNotPointer, ok := a.(AuthorizationPolicyRulesDestinationsHttpHeaderMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AuthorizationPolicyRulesDestinationsHttpHeaderMatch", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HeaderName, actual.HeaderName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("HeaderName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAuthorizationPolicyUpdateAuthorizationPolicyOperation")}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AuthorizationPolicy) urlNormalized() *AuthorizationPolicy {
	normalized := dcl.Copy(*r).(AuthorizationPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *AuthorizationPolicy) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateAuthorizationPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AuthorizationPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AuthorizationPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandAuthorizationPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AuthorizationPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAuthorizationPolicy decodes JSON responses into the AuthorizationPolicy resource schema.
func unmarshalAuthorizationPolicy(b []byte, c *Client) (*AuthorizationPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAuthorizationPolicy(m, c)
}

func unmarshalMapAuthorizationPolicy(m map[string]interface{}, c *Client) (*AuthorizationPolicy, error) {

	return flattenAuthorizationPolicy(c, m), nil
}

// expandAuthorizationPolicy expands AuthorizationPolicy into a JSON request object.
func expandAuthorizationPolicy(c *Client, f *AuthorizationPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/*/locations/%s/authorizationPolicies/%s", f.Name, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}
	if v, err := expandAuthorizationPolicyRulesSlice(c, f.Rules); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else {
		m["rules"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicy flattens AuthorizationPolicy from a JSON request object into the
// AuthorizationPolicy type.
func flattenAuthorizationPolicy(c *Client, i interface{}) *AuthorizationPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &AuthorizationPolicy{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Action = flattenAuthorizationPolicyActionEnum(m["action"])
	res.Rules = flattenAuthorizationPolicyRulesSlice(c, m["rules"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandAuthorizationPolicyRulesMap expands the contents of AuthorizationPolicyRules into a JSON
// request object.
func expandAuthorizationPolicyRulesMap(c *Client, f map[string]AuthorizationPolicyRules) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRules(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesSlice expands the contents of AuthorizationPolicyRules into a JSON
// request object.
func expandAuthorizationPolicyRulesSlice(c *Client, f []AuthorizationPolicyRules) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRules(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesMap flattens the contents of AuthorizationPolicyRules from a JSON
// response object.
func flattenAuthorizationPolicyRulesMap(c *Client, i interface{}) map[string]AuthorizationPolicyRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRules{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRules{}
	}

	items := make(map[string]AuthorizationPolicyRules)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRules(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesSlice flattens the contents of AuthorizationPolicyRules from a JSON
// response object.
func flattenAuthorizationPolicyRulesSlice(c *Client, i interface{}) []AuthorizationPolicyRules {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRules{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRules{}
	}

	items := make([]AuthorizationPolicyRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRules(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRules expands an instance of AuthorizationPolicyRules into a JSON
// request object.
func expandAuthorizationPolicyRules(c *Client, f *AuthorizationPolicyRules) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAuthorizationPolicyRulesSourcesSlice(c, f.Sources); err != nil {
		return nil, fmt.Errorf("error expanding Sources into sources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sources"] = v
	}
	if v, err := expandAuthorizationPolicyRulesDestinationsSlice(c, f.Destinations); err != nil {
		return nil, fmt.Errorf("error expanding Destinations into destinations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["destinations"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRules flattens an instance of AuthorizationPolicyRules from a JSON
// response object.
func flattenAuthorizationPolicyRules(c *Client, i interface{}) *AuthorizationPolicyRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAuthorizationPolicyRules
	}
	r.Sources = flattenAuthorizationPolicyRulesSourcesSlice(c, m["sources"])
	r.Destinations = flattenAuthorizationPolicyRulesDestinationsSlice(c, m["destinations"])

	return r
}

// expandAuthorizationPolicyRulesSourcesMap expands the contents of AuthorizationPolicyRulesSources into a JSON
// request object.
func expandAuthorizationPolicyRulesSourcesMap(c *Client, f map[string]AuthorizationPolicyRulesSources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRulesSources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesSourcesSlice expands the contents of AuthorizationPolicyRulesSources into a JSON
// request object.
func expandAuthorizationPolicyRulesSourcesSlice(c *Client, f []AuthorizationPolicyRulesSources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRulesSources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesSourcesMap flattens the contents of AuthorizationPolicyRulesSources from a JSON
// response object.
func flattenAuthorizationPolicyRulesSourcesMap(c *Client, i interface{}) map[string]AuthorizationPolicyRulesSources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRulesSources{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRulesSources{}
	}

	items := make(map[string]AuthorizationPolicyRulesSources)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRulesSources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesSourcesSlice flattens the contents of AuthorizationPolicyRulesSources from a JSON
// response object.
func flattenAuthorizationPolicyRulesSourcesSlice(c *Client, i interface{}) []AuthorizationPolicyRulesSources {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRulesSources{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRulesSources{}
	}

	items := make([]AuthorizationPolicyRulesSources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRulesSources(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRulesSources expands an instance of AuthorizationPolicyRulesSources into a JSON
// request object.
func expandAuthorizationPolicyRulesSources(c *Client, f *AuthorizationPolicyRulesSources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Principals; v != nil {
		m["principals"] = v
	}
	if v := f.IPBlocks; v != nil {
		m["ipBlocks"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRulesSources flattens an instance of AuthorizationPolicyRulesSources from a JSON
// response object.
func flattenAuthorizationPolicyRulesSources(c *Client, i interface{}) *AuthorizationPolicyRulesSources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRulesSources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAuthorizationPolicyRulesSources
	}
	r.Principals = dcl.FlattenStringSlice(m["principals"])
	r.IPBlocks = dcl.FlattenStringSlice(m["ipBlocks"])

	return r
}

// expandAuthorizationPolicyRulesDestinationsMap expands the contents of AuthorizationPolicyRulesDestinations into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsMap(c *Client, f map[string]AuthorizationPolicyRulesDestinations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesDestinationsSlice expands the contents of AuthorizationPolicyRulesDestinations into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsSlice(c *Client, f []AuthorizationPolicyRulesDestinations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesDestinationsMap flattens the contents of AuthorizationPolicyRulesDestinations from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsMap(c *Client, i interface{}) map[string]AuthorizationPolicyRulesDestinations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRulesDestinations{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRulesDestinations{}
	}

	items := make(map[string]AuthorizationPolicyRulesDestinations)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRulesDestinations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesDestinationsSlice flattens the contents of AuthorizationPolicyRulesDestinations from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsSlice(c *Client, i interface{}) []AuthorizationPolicyRulesDestinations {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRulesDestinations{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRulesDestinations{}
	}

	items := make([]AuthorizationPolicyRulesDestinations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRulesDestinations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRulesDestinations expands an instance of AuthorizationPolicyRulesDestinations into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinations(c *Client, f *AuthorizationPolicyRulesDestinations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Hosts; v != nil {
		m["hosts"] = v
	}
	if v := f.Ports; v != nil {
		m["ports"] = v
	}
	if v := f.Methods; v != nil {
		m["methods"] = v
	}
	if v, err := expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, f.HttpHeaderMatch); err != nil {
		return nil, fmt.Errorf("error expanding HttpHeaderMatch into httpHeaderMatch: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpHeaderMatch"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRulesDestinations flattens an instance of AuthorizationPolicyRulesDestinations from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinations(c *Client, i interface{}) *AuthorizationPolicyRulesDestinations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRulesDestinations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAuthorizationPolicyRulesDestinations
	}
	r.Hosts = dcl.FlattenStringSlice(m["hosts"])
	r.Ports = dcl.FlattenIntSlice(m["ports"])
	r.Methods = dcl.FlattenStringSlice(m["methods"])
	r.HttpHeaderMatch = flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, m["httpHeaderMatch"])

	return r
}

// expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap expands the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap(c *Client, f map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice expands the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, f []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap flattens the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap(c *Client, i interface{}) map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	items := make(map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice flattens the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, i interface{}) []AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	items := make([]AuthorizationPolicyRulesDestinationsHttpHeaderMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch expands an instance of AuthorizationPolicyRulesDestinationsHttpHeaderMatch into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, f *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["headerName"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch flattens an instance of AuthorizationPolicyRulesDestinationsHttpHeaderMatch from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, i interface{}) *AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch
	}
	r.HeaderName = dcl.FlattenString(m["headerName"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])

	return r
}

// flattenAuthorizationPolicyActionEnumMap flattens the contents of AuthorizationPolicyActionEnum from a JSON
// response object.
func flattenAuthorizationPolicyActionEnumMap(c *Client, i interface{}) map[string]AuthorizationPolicyActionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyActionEnum{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyActionEnum{}
	}

	items := make(map[string]AuthorizationPolicyActionEnum)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyActionEnum(item.(interface{}))
	}

	return items
}

// flattenAuthorizationPolicyActionEnumSlice flattens the contents of AuthorizationPolicyActionEnum from a JSON
// response object.
func flattenAuthorizationPolicyActionEnumSlice(c *Client, i interface{}) []AuthorizationPolicyActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyActionEnum{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyActionEnum{}
	}

	items := make([]AuthorizationPolicyActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyActionEnum(item.(interface{})))
	}

	return items
}

// flattenAuthorizationPolicyActionEnum asserts that an interface is a string, and returns a
// pointer to a *AuthorizationPolicyActionEnum with the same value as that string.
func flattenAuthorizationPolicyActionEnum(i interface{}) *AuthorizationPolicyActionEnum {
	s, ok := i.(string)
	if !ok {
		return AuthorizationPolicyActionEnumRef("")
	}

	return AuthorizationPolicyActionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AuthorizationPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAuthorizationPolicy(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type authorizationPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         authorizationPolicyApiOperation
}

func convertFieldDiffsToAuthorizationPolicyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]authorizationPolicyDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []authorizationPolicyDiff
	// For each operation name, create a authorizationPolicyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := authorizationPolicyDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAuthorizationPolicyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAuthorizationPolicyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (authorizationPolicyApiOperation, error) {
	switch opName {

	case "updateAuthorizationPolicyUpdateAuthorizationPolicyOperation":
		return &updateAuthorizationPolicyUpdateAuthorizationPolicyOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractAuthorizationPolicyFields(r *AuthorizationPolicy) error {
	return nil
}
