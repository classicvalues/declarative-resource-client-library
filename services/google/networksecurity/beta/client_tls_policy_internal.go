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

func (r *ClientTlsPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ClientCertificate) {
		if err := r.ClientCertificate.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClientTlsPolicyClientCertificate) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"GrpcEndpoint", "CertificateProviderInstance"}, r.GrpcEndpoint, r.CertificateProviderInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.GrpcEndpoint) {
		if err := r.GrpcEndpoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CertificateProviderInstance) {
		if err := r.CertificateProviderInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) validate() error {
	if err := dcl.Required(r, "targetUri"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) validate() error {
	if err := dcl.Required(r, "pluginInstance"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyServerValidationCa) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"GrpcEndpoint", "CertificateProviderInstance"}, r.GrpcEndpoint, r.CertificateProviderInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.GrpcEndpoint) {
		if err := r.GrpcEndpoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CertificateProviderInstance) {
		if err := r.CertificateProviderInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) validate() error {
	if err := dcl.Required(r, "targetUri"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) validate() error {
	if err := dcl.Required(r, "pluginInstance"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicy) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networksecurity.googleapis.com/v1beta1/", params)
}

func (r *ClientTlsPolicy) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *ClientTlsPolicy) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies", nr.basePath(), userBasePath, params), nil

}

func (r *ClientTlsPolicy) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies?clientTlsPolicyId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *ClientTlsPolicy) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *ClientTlsPolicy) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *ClientTlsPolicy) SetPolicyVerb() string {
	return "POST"
}

func (r *ClientTlsPolicy) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *ClientTlsPolicy) IAMPolicyVersion() int {
	return 3
}

// clientTlsPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type clientTlsPolicyApiOperation interface {
	do(context.Context, *ClientTlsPolicy, *Client) error
}

// newUpdateClientTlsPolicyUpdateClientTlsPolicyRequest creates a request for an
// ClientTlsPolicy resource's UpdateClientTlsPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(ctx context.Context, f *ClientTlsPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Sni; !dcl.IsEmptyValueIndirect(v) {
		req["sni"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificate(c, f.ClientCertificate); err != nil {
		return nil, fmt.Errorf("error expanding ClientCertificate into clientCertificate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["clientCertificate"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaSlice(c, f.ServerValidationCa); err != nil {
		return nil, fmt.Errorf("error expanding ServerValidationCa into serverValidationCa: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["serverValidationCa"] = v
	}
	return req, nil
}

// marshalUpdateClientTlsPolicyUpdateClientTlsPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateClientTlsPolicyUpdateClientTlsPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClientTlsPolicyUpdateClientTlsPolicyOperation) do(ctx context.Context, r *ClientTlsPolicy, c *Client) error {
	_, err := c.GetClientTlsPolicy(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateClientTlsPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(c, req)
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

func (c *Client) listClientTlsPolicyRaw(ctx context.Context, r *ClientTlsPolicy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ClientTlsPolicyMaxPage {
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

type listClientTlsPolicyOperation struct {
	ClientTlsPolicies []map[string]interface{} `json:"clientTlsPolicies"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listClientTlsPolicy(ctx context.Context, r *ClientTlsPolicy, pageToken string, pageSize int32) ([]*ClientTlsPolicy, string, error) {
	b, err := c.listClientTlsPolicyRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listClientTlsPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ClientTlsPolicy
	for _, v := range m.ClientTlsPolicies {
		res, err := unmarshalMapClientTlsPolicy(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllClientTlsPolicy(ctx context.Context, f func(*ClientTlsPolicy) bool, resources []*ClientTlsPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteClientTlsPolicy(ctx, res)
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

type deleteClientTlsPolicyOperation struct{}

func (op *deleteClientTlsPolicyOperation) do(ctx context.Context, r *ClientTlsPolicy, c *Client) error {
	r, err := c.GetClientTlsPolicy(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "ClientTlsPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetClientTlsPolicy checking for existence. error: %v", err)
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
		_, err = c.GetClientTlsPolicy(ctx, r)
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
type createClientTlsPolicyOperation struct {
	response map[string]interface{}
}

func (op *createClientTlsPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createClientTlsPolicyOperation) do(ctx context.Context, r *ClientTlsPolicy, c *Client) error {
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

	if _, err := c.GetClientTlsPolicy(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getClientTlsPolicyRaw(ctx context.Context, r *ClientTlsPolicy) ([]byte, error) {

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

func (c *Client) clientTlsPolicyDiffsForRawDesired(ctx context.Context, rawDesired *ClientTlsPolicy, opts ...dcl.ApplyOption) (initial, desired *ClientTlsPolicy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ClientTlsPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ClientTlsPolicy); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected ClientTlsPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetClientTlsPolicy(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a ClientTlsPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ClientTlsPolicy resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that ClientTlsPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClientTlsPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for ClientTlsPolicy: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for ClientTlsPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClientTlsPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for ClientTlsPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClientTlsPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for ClientTlsPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffClientTlsPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeClientTlsPolicyInitialState(rawInitial, rawDesired *ClientTlsPolicy) (*ClientTlsPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeClientTlsPolicyDesiredState(rawDesired, rawInitial *ClientTlsPolicy, opts ...dcl.ApplyOption) (*ClientTlsPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ClientCertificate = canonicalizeClientTlsPolicyClientCertificate(rawDesired.ClientCertificate, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &ClientTlsPolicy{}
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
	if dcl.StringCanonicalize(rawDesired.Sni, rawInitial.Sni) {
		canonicalDesired.Sni = rawInitial.Sni
	} else {
		canonicalDesired.Sni = rawDesired.Sni
	}
	canonicalDesired.ClientCertificate = canonicalizeClientTlsPolicyClientCertificate(rawDesired.ClientCertificate, rawInitial.ClientCertificate, opts...)
	if dcl.IsZeroValue(rawDesired.ServerValidationCa) {
		canonicalDesired.ServerValidationCa = rawInitial.ServerValidationCa
	} else {
		canonicalDesired.ServerValidationCa = rawDesired.ServerValidationCa
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

func canonicalizeClientTlsPolicyNewState(c *Client, rawNew, rawDesired *ClientTlsPolicy) (*ClientTlsPolicy, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Sni) && dcl.IsEmptyValueIndirect(rawDesired.Sni) {
		rawNew.Sni = rawDesired.Sni
	} else {
		if dcl.StringCanonicalize(rawDesired.Sni, rawNew.Sni) {
			rawNew.Sni = rawDesired.Sni
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientCertificate) && dcl.IsEmptyValueIndirect(rawDesired.ClientCertificate) {
		rawNew.ClientCertificate = rawDesired.ClientCertificate
	} else {
		rawNew.ClientCertificate = canonicalizeNewClientTlsPolicyClientCertificate(c, rawDesired.ClientCertificate, rawNew.ClientCertificate)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerValidationCa) && dcl.IsEmptyValueIndirect(rawDesired.ServerValidationCa) {
		rawNew.ServerValidationCa = rawDesired.ServerValidationCa
	} else {
		rawNew.ServerValidationCa = canonicalizeNewClientTlsPolicyServerValidationCaSlice(c, rawDesired.ServerValidationCa, rawNew.ServerValidationCa)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeClientTlsPolicyClientCertificate(des, initial *ClientTlsPolicyClientCertificate, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.GrpcEndpoint != nil || (initial != nil && initial.GrpcEndpoint != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CertificateProviderInstance) {
			des.GrpcEndpoint = nil
			if initial != nil {
				initial.GrpcEndpoint = nil
			}
		}
	}

	if des.CertificateProviderInstance != nil || (initial != nil && initial.CertificateProviderInstance != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.GrpcEndpoint) {
			des.CertificateProviderInstance = nil
			if initial != nil {
				initial.CertificateProviderInstance = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ClientTlsPolicyClientCertificate{}

	cDes.GrpcEndpoint = canonicalizeClientTlsPolicyClientCertificateGrpcEndpoint(des.GrpcEndpoint, initial.GrpcEndpoint, opts...)
	cDes.CertificateProviderInstance = canonicalizeClientTlsPolicyClientCertificateCertificateProviderInstance(des.CertificateProviderInstance, initial.CertificateProviderInstance, opts...)

	return cDes
}

func canonicalizeNewClientTlsPolicyClientCertificate(c *Client, des, nw *ClientTlsPolicyClientCertificate) *ClientTlsPolicyClientCertificate {
	if des == nil || nw == nil {
		return nw
	}

	nw.GrpcEndpoint = canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpoint(c, des.GrpcEndpoint, nw.GrpcEndpoint)
	nw.CertificateProviderInstance = canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstance(c, des.CertificateProviderInstance, nw.CertificateProviderInstance)

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateSet(c *Client, des, nw []ClientTlsPolicyClientCertificate) []ClientTlsPolicyClientCertificate {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClientTlsPolicyClientCertificateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClientTlsPolicyClientCertificateSlice(c *Client, des, nw []ClientTlsPolicyClientCertificate) []ClientTlsPolicyClientCertificate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClientTlsPolicyClientCertificate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificate(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyClientCertificateGrpcEndpoint(des, initial *ClientTlsPolicyClientCertificateGrpcEndpoint, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClientTlsPolicyClientCertificateGrpcEndpoint{}

	if dcl.StringCanonicalize(des.TargetUri, initial.TargetUri) || dcl.IsZeroValue(des.TargetUri) {
		cDes.TargetUri = initial.TargetUri
	} else {
		cDes.TargetUri = des.TargetUri
	}

	return cDes
}

func canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, des, nw *ClientTlsPolicyClientCertificateGrpcEndpoint) *ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TargetUri, nw.TargetUri) {
		nw.TargetUri = des.TargetUri
	}

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpointSet(c *Client, des, nw []ClientTlsPolicyClientCertificateGrpcEndpoint) []ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificateGrpcEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClientTlsPolicyClientCertificateGrpcEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, des, nw []ClientTlsPolicyClientCertificateGrpcEndpoint) []ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClientTlsPolicyClientCertificateGrpcEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyClientCertificateCertificateProviderInstance(des, initial *ClientTlsPolicyClientCertificateCertificateProviderInstance, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClientTlsPolicyClientCertificateCertificateProviderInstance{}

	if dcl.StringCanonicalize(des.PluginInstance, initial.PluginInstance) || dcl.IsZeroValue(des.PluginInstance) {
		cDes.PluginInstance = initial.PluginInstance
	} else {
		cDes.PluginInstance = des.PluginInstance
	}

	return cDes
}

func canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, des, nw *ClientTlsPolicyClientCertificateCertificateProviderInstance) *ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PluginInstance, nw.PluginInstance) {
		nw.PluginInstance = des.PluginInstance
	}

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstanceSet(c *Client, des, nw []ClientTlsPolicyClientCertificateCertificateProviderInstance) []ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificateCertificateProviderInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClientTlsPolicyClientCertificateCertificateProviderInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, des, nw []ClientTlsPolicyClientCertificateCertificateProviderInstance) []ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClientTlsPolicyClientCertificateCertificateProviderInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstance(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyServerValidationCa(des, initial *ClientTlsPolicyServerValidationCa, opts ...dcl.ApplyOption) *ClientTlsPolicyServerValidationCa {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.GrpcEndpoint != nil || (initial != nil && initial.GrpcEndpoint != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CertificateProviderInstance) {
			des.GrpcEndpoint = nil
			if initial != nil {
				initial.GrpcEndpoint = nil
			}
		}
	}

	if des.CertificateProviderInstance != nil || (initial != nil && initial.CertificateProviderInstance != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.GrpcEndpoint) {
			des.CertificateProviderInstance = nil
			if initial != nil {
				initial.CertificateProviderInstance = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ClientTlsPolicyServerValidationCa{}

	cDes.GrpcEndpoint = canonicalizeClientTlsPolicyServerValidationCaGrpcEndpoint(des.GrpcEndpoint, initial.GrpcEndpoint, opts...)
	cDes.CertificateProviderInstance = canonicalizeClientTlsPolicyServerValidationCaCertificateProviderInstance(des.CertificateProviderInstance, initial.CertificateProviderInstance, opts...)

	return cDes
}

func canonicalizeNewClientTlsPolicyServerValidationCa(c *Client, des, nw *ClientTlsPolicyServerValidationCa) *ClientTlsPolicyServerValidationCa {
	if des == nil || nw == nil {
		return nw
	}

	nw.GrpcEndpoint = canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpoint(c, des.GrpcEndpoint, nw.GrpcEndpoint)
	nw.CertificateProviderInstance = canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstance(c, des.CertificateProviderInstance, nw.CertificateProviderInstance)

	return nw
}

func canonicalizeNewClientTlsPolicyServerValidationCaSet(c *Client, des, nw []ClientTlsPolicyServerValidationCa) []ClientTlsPolicyServerValidationCa {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyServerValidationCa
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClientTlsPolicyServerValidationCaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClientTlsPolicyServerValidationCaSlice(c *Client, des, nw []ClientTlsPolicyServerValidationCa) []ClientTlsPolicyServerValidationCa {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClientTlsPolicyServerValidationCa
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyServerValidationCa(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyServerValidationCaGrpcEndpoint(des, initial *ClientTlsPolicyServerValidationCaGrpcEndpoint, opts ...dcl.ApplyOption) *ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClientTlsPolicyServerValidationCaGrpcEndpoint{}

	if dcl.StringCanonicalize(des.TargetUri, initial.TargetUri) || dcl.IsZeroValue(des.TargetUri) {
		cDes.TargetUri = initial.TargetUri
	} else {
		cDes.TargetUri = des.TargetUri
	}

	return cDes
}

func canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, des, nw *ClientTlsPolicyServerValidationCaGrpcEndpoint) *ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TargetUri, nw.TargetUri) {
		nw.TargetUri = des.TargetUri
	}

	return nw
}

func canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpointSet(c *Client, des, nw []ClientTlsPolicyServerValidationCaGrpcEndpoint) []ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyServerValidationCaGrpcEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClientTlsPolicyServerValidationCaGrpcEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, des, nw []ClientTlsPolicyServerValidationCaGrpcEndpoint) []ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClientTlsPolicyServerValidationCaGrpcEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyServerValidationCaCertificateProviderInstance(des, initial *ClientTlsPolicyServerValidationCaCertificateProviderInstance, opts ...dcl.ApplyOption) *ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClientTlsPolicyServerValidationCaCertificateProviderInstance{}

	if dcl.StringCanonicalize(des.PluginInstance, initial.PluginInstance) || dcl.IsZeroValue(des.PluginInstance) {
		cDes.PluginInstance = initial.PluginInstance
	} else {
		cDes.PluginInstance = des.PluginInstance
	}

	return cDes
}

func canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, des, nw *ClientTlsPolicyServerValidationCaCertificateProviderInstance) *ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PluginInstance, nw.PluginInstance) {
		nw.PluginInstance = des.PluginInstance
	}

	return nw
}

func canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstanceSet(c *Client, des, nw []ClientTlsPolicyServerValidationCaCertificateProviderInstance) []ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyServerValidationCaCertificateProviderInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClientTlsPolicyServerValidationCaCertificateProviderInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, des, nw []ClientTlsPolicyServerValidationCaCertificateProviderInstance) []ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClientTlsPolicyServerValidationCaCertificateProviderInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &d, &n))
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
func diffClientTlsPolicy(c *Client, desired, actual *ClientTlsPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateClientTlsPolicyUpdateClientTlsPolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateClientTlsPolicyUpdateClientTlsPolicyOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sni, actual.Sni, dcl.Info{OperationSelector: dcl.TriggersOperation("updateClientTlsPolicyUpdateClientTlsPolicyOperation")}, fn.AddNest("Sni")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientCertificate, actual.ClientCertificate, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateNewStyle, EmptyObject: EmptyClientTlsPolicyClientCertificate, OperationSelector: dcl.TriggersOperation("updateClientTlsPolicyUpdateClientTlsPolicyOperation")}, fn.AddNest("ClientCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerValidationCa, actual.ServerValidationCa, dcl.Info{ObjectFunction: compareClientTlsPolicyServerValidationCaNewStyle, EmptyObject: EmptyClientTlsPolicyServerValidationCa, OperationSelector: dcl.TriggersOperation("updateClientTlsPolicyUpdateClientTlsPolicyOperation")}, fn.AddNest("ServerValidationCa")); len(ds) != 0 || err != nil {
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
func compareClientTlsPolicyClientCertificateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificate)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificate or *ClientTlsPolicyClientCertificate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificate)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GrpcEndpoint, actual.GrpcEndpoint, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateGrpcEndpointNewStyle, EmptyObject: EmptyClientTlsPolicyClientCertificateGrpcEndpoint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrpcEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificateProviderInstance, actual.CertificateProviderInstance, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateCertificateProviderInstanceNewStyle, EmptyObject: EmptyClientTlsPolicyClientCertificateCertificateProviderInstance, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificateProviderInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyClientCertificateGrpcEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificateGrpcEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificateGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateGrpcEndpoint or *ClientTlsPolicyClientCertificateGrpcEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificateGrpcEndpoint)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificateGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateGrpcEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetUri, actual.TargetUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyClientCertificateCertificateProviderInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificateCertificateProviderInstance)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificateCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateCertificateProviderInstance or *ClientTlsPolicyClientCertificateCertificateProviderInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificateCertificateProviderInstance)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificateCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateCertificateProviderInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PluginInstance, actual.PluginInstance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PluginInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyServerValidationCaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyServerValidationCa)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyServerValidationCa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCa or *ClientTlsPolicyServerValidationCa", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyServerValidationCa)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyServerValidationCa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCa", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GrpcEndpoint, actual.GrpcEndpoint, dcl.Info{ObjectFunction: compareClientTlsPolicyServerValidationCaGrpcEndpointNewStyle, EmptyObject: EmptyClientTlsPolicyServerValidationCaGrpcEndpoint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrpcEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificateProviderInstance, actual.CertificateProviderInstance, dcl.Info{ObjectFunction: compareClientTlsPolicyServerValidationCaCertificateProviderInstanceNewStyle, EmptyObject: EmptyClientTlsPolicyServerValidationCaCertificateProviderInstance, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificateProviderInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyServerValidationCaGrpcEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyServerValidationCaGrpcEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyServerValidationCaGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaGrpcEndpoint or *ClientTlsPolicyServerValidationCaGrpcEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyServerValidationCaGrpcEndpoint)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyServerValidationCaGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaGrpcEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetUri, actual.TargetUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyServerValidationCaCertificateProviderInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyServerValidationCaCertificateProviderInstance)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyServerValidationCaCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaCertificateProviderInstance or *ClientTlsPolicyServerValidationCaCertificateProviderInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyServerValidationCaCertificateProviderInstance)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyServerValidationCaCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaCertificateProviderInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PluginInstance, actual.PluginInstance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PluginInstance")); len(ds) != 0 || err != nil {
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
func (r *ClientTlsPolicy) urlNormalized() *ClientTlsPolicy {
	normalized := dcl.Copy(*r).(ClientTlsPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Sni = dcl.SelfLinkToName(r.Sni)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *ClientTlsPolicy) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateClientTlsPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ClientTlsPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ClientTlsPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandClientTlsPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ClientTlsPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalClientTlsPolicy decodes JSON responses into the ClientTlsPolicy resource schema.
func unmarshalClientTlsPolicy(b []byte, c *Client) (*ClientTlsPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapClientTlsPolicy(m, c)
}

func unmarshalMapClientTlsPolicy(m map[string]interface{}, c *Client) (*ClientTlsPolicy, error) {

	return flattenClientTlsPolicy(c, m), nil
}

// expandClientTlsPolicy expands ClientTlsPolicy into a JSON request object.
func expandClientTlsPolicy(c *Client, f *ClientTlsPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/clientTlsPolicies/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
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
	if v := f.Sni; !dcl.IsEmptyValueIndirect(v) {
		m["sni"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificate(c, f.ClientCertificate); err != nil {
		return nil, fmt.Errorf("error expanding ClientCertificate into clientCertificate: %w", err)
	} else if v != nil {
		m["clientCertificate"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaSlice(c, f.ServerValidationCa); err != nil {
		return nil, fmt.Errorf("error expanding ServerValidationCa into serverValidationCa: %w", err)
	} else {
		m["serverValidationCa"] = v
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

// flattenClientTlsPolicy flattens ClientTlsPolicy from a JSON request object into the
// ClientTlsPolicy type.
func flattenClientTlsPolicy(c *Client, i interface{}) *ClientTlsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &ClientTlsPolicy{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Sni = dcl.FlattenString(m["sni"])
	res.ClientCertificate = flattenClientTlsPolicyClientCertificate(c, m["clientCertificate"])
	res.ServerValidationCa = flattenClientTlsPolicyServerValidationCaSlice(c, m["serverValidationCa"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandClientTlsPolicyClientCertificateMap expands the contents of ClientTlsPolicyClientCertificate into a JSON
// request object.
func expandClientTlsPolicyClientCertificateMap(c *Client, f map[string]ClientTlsPolicyClientCertificate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateSlice expands the contents of ClientTlsPolicyClientCertificate into a JSON
// request object.
func expandClientTlsPolicyClientCertificateSlice(c *Client, f []ClientTlsPolicyClientCertificate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateMap flattens the contents of ClientTlsPolicyClientCertificate from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificate{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificate{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificate)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateSlice flattens the contents of ClientTlsPolicyClientCertificate from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificate {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificate{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificate{}
	}

	items := make([]ClientTlsPolicyClientCertificate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificate(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificate expands an instance of ClientTlsPolicyClientCertificate into a JSON
// request object.
func expandClientTlsPolicyClientCertificate(c *Client, f *ClientTlsPolicyClientCertificate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandClientTlsPolicyClientCertificateGrpcEndpoint(c, f.GrpcEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding GrpcEndpoint into grpcEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["grpcEndpoint"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificateCertificateProviderInstance(c, f.CertificateProviderInstance); err != nil {
		return nil, fmt.Errorf("error expanding CertificateProviderInstance into certificateProviderInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certificateProviderInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificate flattens an instance of ClientTlsPolicyClientCertificate from a JSON
// response object.
func flattenClientTlsPolicyClientCertificate(c *Client, i interface{}) *ClientTlsPolicyClientCertificate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClientTlsPolicyClientCertificate
	}
	r.GrpcEndpoint = flattenClientTlsPolicyClientCertificateGrpcEndpoint(c, m["grpcEndpoint"])
	r.CertificateProviderInstance = flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c, m["certificateProviderInstance"])

	return r
}

// expandClientTlsPolicyClientCertificateGrpcEndpointMap expands the contents of ClientTlsPolicyClientCertificateGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyClientCertificateGrpcEndpointMap(c *Client, f map[string]ClientTlsPolicyClientCertificateGrpcEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificateGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateGrpcEndpointSlice expands the contents of ClientTlsPolicyClientCertificateGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, f []ClientTlsPolicyClientCertificateGrpcEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificateGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateGrpcEndpointMap flattens the contents of ClientTlsPolicyClientCertificateGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateGrpcEndpointMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificateGrpcEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificateGrpcEndpoint)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificateGrpcEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateGrpcEndpointSlice flattens the contents of ClientTlsPolicyClientCertificateGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificateGrpcEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	items := make([]ClientTlsPolicyClientCertificateGrpcEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificateGrpcEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificateGrpcEndpoint expands an instance of ClientTlsPolicyClientCertificateGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, f *ClientTlsPolicyClientCertificateGrpcEndpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetUri; !dcl.IsEmptyValueIndirect(v) {
		m["targetUri"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificateGrpcEndpoint flattens an instance of ClientTlsPolicyClientCertificateGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, i interface{}) *ClientTlsPolicyClientCertificateGrpcEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificateGrpcEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClientTlsPolicyClientCertificateGrpcEndpoint
	}
	r.TargetUri = dcl.FlattenString(m["targetUri"])

	return r
}

// expandClientTlsPolicyClientCertificateCertificateProviderInstanceMap expands the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyClientCertificateCertificateProviderInstanceMap(c *Client, f map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificateCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateCertificateProviderInstanceSlice expands the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, f []ClientTlsPolicyClientCertificateCertificateProviderInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificateCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateCertificateProviderInstanceMap flattens the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateCertificateProviderInstanceMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateCertificateProviderInstanceSlice flattens the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificateCertificateProviderInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	items := make([]ClientTlsPolicyClientCertificateCertificateProviderInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificateCertificateProviderInstance expands an instance of ClientTlsPolicyClientCertificateCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, f *ClientTlsPolicyClientCertificateCertificateProviderInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PluginInstance; !dcl.IsEmptyValueIndirect(v) {
		m["pluginInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificateCertificateProviderInstance flattens an instance of ClientTlsPolicyClientCertificateCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, i interface{}) *ClientTlsPolicyClientCertificateCertificateProviderInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificateCertificateProviderInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClientTlsPolicyClientCertificateCertificateProviderInstance
	}
	r.PluginInstance = dcl.FlattenString(m["pluginInstance"])

	return r
}

// expandClientTlsPolicyServerValidationCaMap expands the contents of ClientTlsPolicyServerValidationCa into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaMap(c *Client, f map[string]ClientTlsPolicyServerValidationCa) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyServerValidationCa(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyServerValidationCaSlice expands the contents of ClientTlsPolicyServerValidationCa into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaSlice(c *Client, f []ClientTlsPolicyServerValidationCa) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyServerValidationCa(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyServerValidationCaMap flattens the contents of ClientTlsPolicyServerValidationCa from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaMap(c *Client, i interface{}) map[string]ClientTlsPolicyServerValidationCa {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyServerValidationCa{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyServerValidationCa{}
	}

	items := make(map[string]ClientTlsPolicyServerValidationCa)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyServerValidationCa(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyServerValidationCaSlice flattens the contents of ClientTlsPolicyServerValidationCa from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaSlice(c *Client, i interface{}) []ClientTlsPolicyServerValidationCa {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyServerValidationCa{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyServerValidationCa{}
	}

	items := make([]ClientTlsPolicyServerValidationCa, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyServerValidationCa(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyServerValidationCa expands an instance of ClientTlsPolicyServerValidationCa into a JSON
// request object.
func expandClientTlsPolicyServerValidationCa(c *Client, f *ClientTlsPolicyServerValidationCa) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandClientTlsPolicyServerValidationCaGrpcEndpoint(c, f.GrpcEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding GrpcEndpoint into grpcEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["grpcEndpoint"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c, f.CertificateProviderInstance); err != nil {
		return nil, fmt.Errorf("error expanding CertificateProviderInstance into certificateProviderInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certificateProviderInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyServerValidationCa flattens an instance of ClientTlsPolicyServerValidationCa from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCa(c *Client, i interface{}) *ClientTlsPolicyServerValidationCa {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyServerValidationCa{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClientTlsPolicyServerValidationCa
	}
	r.GrpcEndpoint = flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c, m["grpcEndpoint"])
	r.CertificateProviderInstance = flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c, m["certificateProviderInstance"])

	return r
}

// expandClientTlsPolicyServerValidationCaGrpcEndpointMap expands the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaGrpcEndpointMap(c *Client, f map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyServerValidationCaGrpcEndpointSlice expands the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, f []ClientTlsPolicyServerValidationCaGrpcEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyServerValidationCaGrpcEndpointMap flattens the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaGrpcEndpointMap(c *Client, i interface{}) map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	items := make(map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyServerValidationCaGrpcEndpointSlice flattens the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, i interface{}) []ClientTlsPolicyServerValidationCaGrpcEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	items := make([]ClientTlsPolicyServerValidationCaGrpcEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyServerValidationCaGrpcEndpoint expands an instance of ClientTlsPolicyServerValidationCaGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, f *ClientTlsPolicyServerValidationCaGrpcEndpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetUri; !dcl.IsEmptyValueIndirect(v) {
		m["targetUri"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyServerValidationCaGrpcEndpoint flattens an instance of ClientTlsPolicyServerValidationCaGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, i interface{}) *ClientTlsPolicyServerValidationCaGrpcEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyServerValidationCaGrpcEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClientTlsPolicyServerValidationCaGrpcEndpoint
	}
	r.TargetUri = dcl.FlattenString(m["targetUri"])

	return r
}

// expandClientTlsPolicyServerValidationCaCertificateProviderInstanceMap expands the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaCertificateProviderInstanceMap(c *Client, f map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice expands the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, f []ClientTlsPolicyServerValidationCaCertificateProviderInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceMap flattens the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceMap(c *Client, i interface{}) map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	items := make(map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice flattens the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, i interface{}) []ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	items := make([]ClientTlsPolicyServerValidationCaCertificateProviderInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyServerValidationCaCertificateProviderInstance expands an instance of ClientTlsPolicyServerValidationCaCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, f *ClientTlsPolicyServerValidationCaCertificateProviderInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PluginInstance; !dcl.IsEmptyValueIndirect(v) {
		m["pluginInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyServerValidationCaCertificateProviderInstance flattens an instance of ClientTlsPolicyServerValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, i interface{}) *ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyServerValidationCaCertificateProviderInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClientTlsPolicyServerValidationCaCertificateProviderInstance
	}
	r.PluginInstance = dcl.FlattenString(m["pluginInstance"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ClientTlsPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalClientTlsPolicy(b, c)
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

type clientTlsPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clientTlsPolicyApiOperation
}

func convertFieldDiffsToClientTlsPolicyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]clientTlsPolicyDiff, error) {
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
	var diffs []clientTlsPolicyDiff
	// For each operation name, create a clientTlsPolicyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := clientTlsPolicyDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToClientTlsPolicyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToClientTlsPolicyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (clientTlsPolicyApiOperation, error) {
	switch opName {

	case "updateClientTlsPolicyUpdateClientTlsPolicyOperation":
		return &updateClientTlsPolicyUpdateClientTlsPolicyOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractClientTlsPolicyFields(r *ClientTlsPolicy) error {
	return nil
}
