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
	"context"
	"strconv"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Returns true if m and n represent the same floating point value of seconds.
func canonicalizeOSPolicyAssignmentRolloutMinWaitDuration(m, n interface{}) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil || n == nil {
		return false
	}
	mStr := dcl.ValueOrEmptyString(m)
	mFloat, err := strconv.ParseFloat(mStr[:len(mStr)-1], 64)
	if err != nil {
		return false
	}
	nStr := dcl.ValueOrEmptyString(n)
	nFloat, err := strconv.ParseFloat(nStr[:len(nStr)-1], 64)
	if err != nil {
		return false
	}
	return mFloat == nFloat
}

// Waits for os policy assignment to be done reconciling before deletion.
func (r *OSPolicyAssignment) waitForNotReconciling(ctx context.Context, client *Client) error {
	return dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		nr, err := client.GetOSPolicyAssignment(ctx, r)
		if err != nil {
			return nil, err
		}
		if dcl.ValueOrEmptyBool(nr.Reconciling) {
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, nil
	}, client.Config.RetryProvider)
}
