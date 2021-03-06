// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order Placement, Cancellation, Amending, and History
// swagger:model Order
type Order struct {

	// account
	Account int64 `json:"account,omitempty"`

	// avg px
	AvgPx float64 `json:"avgPx,omitempty"`

	// cl ord ID
	ClOrdID string `json:"clOrdID,omitempty"`

	// cl ord link ID
	ClOrdLinkID string `json:"clOrdLinkID,omitempty"`

	// contingency type
	ContingencyType string `json:"contingencyType,omitempty"`

	// cum qty
	CumQty int64 `json:"cumQty,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// display qty
	DisplayQty int64 `json:"displayQty,omitempty"`

	// ex destination
	ExDestination string `json:"exDestination,omitempty"`

	// exec inst
	ExecInst string `json:"execInst,omitempty"`

	// leaves qty
	LeavesQty int64 `json:"leavesQty,omitempty"`

	// multi leg reporting type
	MultiLegReportingType string `json:"multiLegReportingType,omitempty"`

	// ord rej reason
	OrdRejReason string `json:"ordRejReason,omitempty"`

	// ord status
	OrdStatus string `json:"ordStatus,omitempty"`

	// ord type
	OrdType string `json:"ordType,omitempty"`

	// order ID
	// Required: true
	OrderID *string `json:"orderID"`

	// order qty
	OrderQty int64 `json:"orderQty,omitempty"`

	// peg offset value
	PegOffsetValue float64 `json:"pegOffsetValue,omitempty"`

	// peg price type
	PegPriceType string `json:"pegPriceType,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// settl currency
	SettlCurrency string `json:"settlCurrency,omitempty"`

	// side
	Side string `json:"side,omitempty"`

	// simple cum qty
	SimpleCumQty float64 `json:"simpleCumQty,omitempty"`

	// simple leaves qty
	SimpleLeavesQty float64 `json:"simpleLeavesQty,omitempty"`

	// simple order qty
	SimpleOrderQty float64 `json:"simpleOrderQty,omitempty"`

	// stop px
	StopPx float64 `json:"stopPx,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// time in force
	TimeInForce string `json:"timeInForce,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// transact time
	// Format: date-time
	TransactTime strfmt.DateTime `json:"transactTime,omitempty"`

	// triggered
	Triggered string `json:"triggered,omitempty"`

	// working indicator
	WorkingIndicator bool `json:"workingIndicator,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("orderID", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateTransactTime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactTime) { // not required
		return nil
	}

	if err := validate.FormatOf("transactTime", "body", "date-time", m.TransactTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
