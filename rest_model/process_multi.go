// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProcessMulti process multi
//
// swagger:model processMulti
type ProcessMulti struct {

	// hashes
	Hashes []string `json:"hashes"`

	// os type
	// Required: true
	OsType *OsType `json:"osType"`

	// path
	// Required: true
	Path *string `json:"path"`

	// signer fingerprints
	SignerFingerprints []string `json:"signerFingerprints"`
}

// Validate validates this process multi
func (m *ProcessMulti) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessMulti) validateOsType(formats strfmt.Registry) error {

	if err := validate.Required("osType", "body", m.OsType); err != nil {
		return err
	}

	if err := validate.Required("osType", "body", m.OsType); err != nil {
		return err
	}

	if m.OsType != nil {
		if err := m.OsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osType")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessMulti) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this process multi based on the context it is used
func (m *ProcessMulti) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessMulti) contextValidateOsType(ctx context.Context, formats strfmt.Registry) error {

	if m.OsType != nil {
		if err := m.OsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessMulti) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessMulti) UnmarshalBinary(b []byte) error {
	var res ProcessMulti
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
