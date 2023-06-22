/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * CreateTransactionResponse - OK
 */
public class CreateTransactionResponse {
    @JsonProperty("data")
    public Transaction data;

    public CreateTransactionResponse withData(Transaction data) {
        this.data = data;
        return this;
    }
    
    public CreateTransactionResponse(@JsonProperty("data") Transaction data) {
        this.data = data;
  }
}