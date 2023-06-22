/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class GetVersionsResponse {
    
    public String contentType;

    public GetVersionsResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * OK
     */
    
    public com.formance.formance_sdk.models.shared.GetVersionsResponse getVersionsResponse;

    public GetVersionsResponse withGetVersionsResponse(com.formance.formance_sdk.models.shared.GetVersionsResponse getVersionsResponse) {
        this.getVersionsResponse = getVersionsResponse;
        return this;
    }
    
    
    public Integer statusCode;

    public GetVersionsResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public GetVersionsResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    public GetVersionsResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}