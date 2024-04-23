// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2TriggerData struct {
	Event      string                 `json:"event"`
	Filter     *string                `json:"filter,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Vars       map[string]interface{} `json:"vars,omitempty"`
	WorkflowID string                 `json:"workflowID"`
}

func (o *V2TriggerData) GetEvent() string {
	if o == nil {
		return ""
	}
	return o.Event
}

func (o *V2TriggerData) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *V2TriggerData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *V2TriggerData) GetVars() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Vars
}

func (o *V2TriggerData) GetWorkflowID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowID
}
