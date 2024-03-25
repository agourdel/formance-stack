// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/sdkerrors"
	"net/http"
)

type V2UpdateLedgerMetadataRequest struct {
	RequestBody map[string]string `request:"mediaType=application/json"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *V2UpdateLedgerMetadataRequest) GetRequestBody() map[string]string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2UpdateLedgerMetadataRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type V2UpdateLedgerMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	V2ErrorResponse *sdkerrors.V2ErrorResponse
}

func (o *V2UpdateLedgerMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2UpdateLedgerMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2UpdateLedgerMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2UpdateLedgerMetadataResponse) GetV2ErrorResponse() *sdkerrors.V2ErrorResponse {
	if o == nil {
		return nil
	}
	return o.V2ErrorResponse
}
