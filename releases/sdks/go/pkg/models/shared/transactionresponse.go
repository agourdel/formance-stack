// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransactionResponse struct {
	Data Transaction `json:"data"`
}

func (o *TransactionResponse) GetData() Transaction {
	if o == nil {
		return Transaction{}
	}
	return o.Data
}