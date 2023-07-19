// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Run run
//
// swagger:model run
type Run struct {

	// cluster
	// Required: true
	// Min Length: 1
	Cluster string `json:"cluster"`

	// exit code
	ExitCode *int32 `json:"exitCode,omitempty"`

	// finished
	// Format: date-time
	Finished *strfmt.DateTime `json:"finished,omitempty"`

	// job run state
	// Required: true
	// Enum: [RUN_PENDING RUN_RUNNING RUN_SUCCEEDED RUN_FAILED RUN_TERMINATED RUN_PREEMPTED RUN_UNABLE_TO_SCHEDULE RUN_LEASE_RETURNED RUN_LEASE_EXPIRED RUN_MAX_RUNS_EXCEEDED RUN_LEASED]
	JobRunState string `json:"jobRunState"`

	// leased
	// Min Length: 1
	// Format: date-time
	Leased *strfmt.DateTime `json:"leased,omitempty"`

	// node
	Node *string `json:"node,omitempty"`

	// pending
	// Min Length: 1
	// Format: date-time
	Pending *strfmt.DateTime `json:"pending,omitempty"`

	// run Id
	// Required: true
	// Min Length: 1
	RunID string `json:"runId"`

	// started
	// Format: date-time
	Started *strfmt.DateTime `json:"started,omitempty"`
}

// Validate validates this run
func (m *Run) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobRunState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeased(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePending(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Run) validateCluster(formats strfmt.Registry) error {

	if err := validate.RequiredString("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if err := validate.MinLength("cluster", "body", m.Cluster, 1); err != nil {
		return err
	}

	return nil
}

func (m *Run) validateFinished(formats strfmt.Registry) error {
	if swag.IsZero(m.Finished) { // not required
		return nil
	}

	if err := validate.FormatOf("finished", "body", "date-time", m.Finished.String(), formats); err != nil {
		return err
	}

	return nil
}

var runTypeJobRunStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUN_PENDING","RUN_RUNNING","RUN_SUCCEEDED","RUN_FAILED","RUN_TERMINATED","RUN_PREEMPTED","RUN_UNABLE_TO_SCHEDULE","RUN_LEASE_RETURNED","RUN_LEASE_EXPIRED","RUN_MAX_RUNS_EXCEEDED","RUN_LEASED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runTypeJobRunStatePropEnum = append(runTypeJobRunStatePropEnum, v)
	}
}

const (

	// RunJobRunStateRUNPENDING captures enum value "RUN_PENDING"
	RunJobRunStateRUNPENDING string = "RUN_PENDING"

	// RunJobRunStateRUNRUNNING captures enum value "RUN_RUNNING"
	RunJobRunStateRUNRUNNING string = "RUN_RUNNING"

	// RunJobRunStateRUNSUCCEEDED captures enum value "RUN_SUCCEEDED"
	RunJobRunStateRUNSUCCEEDED string = "RUN_SUCCEEDED"

	// RunJobRunStateRUNFAILED captures enum value "RUN_FAILED"
	RunJobRunStateRUNFAILED string = "RUN_FAILED"

	// RunJobRunStateRUNTERMINATED captures enum value "RUN_TERMINATED"
	RunJobRunStateRUNTERMINATED string = "RUN_TERMINATED"

	// RunJobRunStateRUNPREEMPTED captures enum value "RUN_PREEMPTED"
	RunJobRunStateRUNPREEMPTED string = "RUN_PREEMPTED"

	// RunJobRunStateRUNUNABLETOSCHEDULE captures enum value "RUN_UNABLE_TO_SCHEDULE"
	RunJobRunStateRUNUNABLETOSCHEDULE string = "RUN_UNABLE_TO_SCHEDULE"

	// RunJobRunStateRUNLEASERETURNED captures enum value "RUN_LEASE_RETURNED"
	RunJobRunStateRUNLEASERETURNED string = "RUN_LEASE_RETURNED"

	// RunJobRunStateRUNLEASEEXPIRED captures enum value "RUN_LEASE_EXPIRED"
	RunJobRunStateRUNLEASEEXPIRED string = "RUN_LEASE_EXPIRED"

	// RunJobRunStateRUNMAXRUNSEXCEEDED captures enum value "RUN_MAX_RUNS_EXCEEDED"
	RunJobRunStateRUNMAXRUNSEXCEEDED string = "RUN_MAX_RUNS_EXCEEDED"

	// RunJobRunStateRUNLEASED captures enum value "RUN_LEASED"
	RunJobRunStateRUNLEASED string = "RUN_LEASED"
)

// prop value enum
func (m *Run) validateJobRunStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, runTypeJobRunStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Run) validateJobRunState(formats strfmt.Registry) error {

	if err := validate.RequiredString("jobRunState", "body", m.JobRunState); err != nil {
		return err
	}

	// value enum
	if err := m.validateJobRunStateEnum("jobRunState", "body", m.JobRunState); err != nil {
		return err
	}

	return nil
}

func (m *Run) validateLeased(formats strfmt.Registry) error {
	if swag.IsZero(m.Leased) { // not required
		return nil
	}

	if err := validate.MinLength("leased", "body", m.Leased.String(), 1); err != nil {
		return err
	}

	if err := validate.FormatOf("leased", "body", "date-time", m.Leased.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Run) validatePending(formats strfmt.Registry) error {
	if swag.IsZero(m.Pending) { // not required
		return nil
	}

	if err := validate.MinLength("pending", "body", m.Pending.String(), 1); err != nil {
		return err
	}

	if err := validate.FormatOf("pending", "body", "date-time", m.Pending.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Run) validateRunID(formats strfmt.Registry) error {

	if err := validate.RequiredString("runId", "body", m.RunID); err != nil {
		return err
	}

	if err := validate.MinLength("runId", "body", m.RunID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Run) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this run based on context it is used
func (m *Run) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Run) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Run) UnmarshalBinary(b []byte) error {
	var res Run
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
