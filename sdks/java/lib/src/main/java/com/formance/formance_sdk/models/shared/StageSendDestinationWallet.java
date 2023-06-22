/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class StageSendDestinationWallet {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("balance")
    public String balance;

    public StageSendDestinationWallet withBalance(String balance) {
        this.balance = balance;
        return this;
    }
    
    @JsonProperty("id")
    public String id;

    public StageSendDestinationWallet withId(String id) {
        this.id = id;
        return this;
    }
    
    public StageSendDestinationWallet(@JsonProperty("id") String id) {
        this.id = id;
  }
}