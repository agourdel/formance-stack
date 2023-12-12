// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type GetPolicyRequest struct {
	// The policy ID.
	PolicyID string `pathParam:"style=simple,explode=false,name=policyID"`
}

func (o *GetPolicyRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

type GetPolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PolicyResponse *shared.PolicyResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error response
	ReconciliationErrorResponse *shared.ReconciliationErrorResponse
}

func (o *GetPolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPolicyResponse) GetPolicyResponse() *shared.PolicyResponse {
	if o == nil {
		return nil
	}
	return o.PolicyResponse
}

func (o *GetPolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPolicyResponse) GetReconciliationErrorResponse() *shared.ReconciliationErrorResponse {
	if o == nil {
		return nil
	}
	return o.ReconciliationErrorResponse
}