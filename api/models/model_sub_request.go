package models

type SensorSubRequest struct {
	Address string `json:"address"`

	MWM int8 `json:"mwm"`

	TickRate int8 `json:"tickRate"`

	Node string `json:"node"`

	Id string `json:"id,omitempty"`
}
