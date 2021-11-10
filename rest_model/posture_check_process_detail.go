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
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureCheckProcessDetail posture check process detail
//
// swagger:model postureCheckProcessDetail
type PostureCheckProcessDetail struct {
	linksField Links

	createdAtField *strfmt.DateTime

	idField *string

	nameField *string

	roleAttributesField *Attributes

	tagsField *Tags

	updatedAtField *strfmt.DateTime

	versionField *int64

	// process
	// Required: true
	Process *Process `json:"process"`
}

// Links gets the links of this subtype
func (m *PostureCheckProcessDetail) Links() Links {
	return m.linksField
}

// SetLinks sets the links of this subtype
func (m *PostureCheckProcessDetail) SetLinks(val Links) {
	m.linksField = val
}

// CreatedAt gets the created at of this subtype
func (m *PostureCheckProcessDetail) CreatedAt() *strfmt.DateTime {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this subtype
func (m *PostureCheckProcessDetail) SetCreatedAt(val *strfmt.DateTime) {
	m.createdAtField = val
}

// ID gets the id of this subtype
func (m *PostureCheckProcessDetail) ID() *string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *PostureCheckProcessDetail) SetID(val *string) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *PostureCheckProcessDetail) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PostureCheckProcessDetail) SetName(val *string) {
	m.nameField = val
}

// RoleAttributes gets the role attributes of this subtype
func (m *PostureCheckProcessDetail) RoleAttributes() *Attributes {
	return m.roleAttributesField
}

// SetRoleAttributes sets the role attributes of this subtype
func (m *PostureCheckProcessDetail) SetRoleAttributes(val *Attributes) {
	m.roleAttributesField = val
}

// Tags gets the tags of this subtype
func (m *PostureCheckProcessDetail) Tags() *Tags {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *PostureCheckProcessDetail) SetTags(val *Tags) {
	m.tagsField = val
}

// TypeID gets the type Id of this subtype
func (m *PostureCheckProcessDetail) TypeID() string {
	return "PROCESS"
}

// SetTypeID sets the type Id of this subtype
func (m *PostureCheckProcessDetail) SetTypeID(val string) {
}

// UpdatedAt gets the updated at of this subtype
func (m *PostureCheckProcessDetail) UpdatedAt() *strfmt.DateTime {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this subtype
func (m *PostureCheckProcessDetail) SetUpdatedAt(val *strfmt.DateTime) {
	m.updatedAtField = val
}

// Version gets the version of this subtype
func (m *PostureCheckProcessDetail) Version() *int64 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *PostureCheckProcessDetail) SetVersion(val *int64) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckProcessDetail) UnmarshalJSON(raw []byte) error {
	var data struct {

		// process
		// Required: true
		Process *Process `json:"process"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Links Links `json:"_links"`

		CreatedAt *strfmt.DateTime `json:"createdAt"`

		ID *string `json:"id"`

		Name *string `json:"name"`

		RoleAttributes *Attributes `json:"roleAttributes"`

		Tags *Tags `json:"tags"`

		TypeID string `json:"typeId"`

		UpdatedAt *strfmt.DateTime `json:"updatedAt"`

		Version *int64 `json:"version"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckProcessDetail

	result.linksField = base.Links

	result.createdAtField = base.CreatedAt

	result.idField = base.ID

	result.nameField = base.Name

	result.roleAttributesField = base.RoleAttributes

	result.tagsField = base.Tags

	if base.TypeID != result.TypeID() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid typeId value: %q", base.TypeID)
	}
	result.updatedAtField = base.UpdatedAt

	result.versionField = base.Version

	result.Process = data.Process

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckProcessDetail) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// process
		// Required: true
		Process *Process `json:"process"`
	}{

		Process: m.Process,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Links Links `json:"_links"`

		CreatedAt *strfmt.DateTime `json:"createdAt"`

		ID *string `json:"id"`

		Name *string `json:"name"`

		RoleAttributes *Attributes `json:"roleAttributes"`

		Tags *Tags `json:"tags"`

		TypeID string `json:"typeId"`

		UpdatedAt *strfmt.DateTime `json:"updatedAt"`

		Version *int64 `json:"version"`
	}{

		Links: m.Links(),

		CreatedAt: m.CreatedAt(),

		ID: m.ID(),

		Name: m.Name(),

		RoleAttributes: m.RoleAttributes(),

		Tags: m.Tags(),

		TypeID: m.TypeID(),

		UpdatedAt: m.UpdatedAt(),

		Version: m.Version(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this posture check process detail
func (m *PostureCheckProcessDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckProcessDetail) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links()); err != nil {
		return err
	}

	if m.Links() != nil {
		if err := m.Links().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt()); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateRoleAttributes(formats strfmt.Registry) error {

	if err := validate.Required("roleAttributes", "body", m.RoleAttributes()); err != nil {
		return err
	}

	if m.RoleAttributes() != nil {
		if err := m.RoleAttributes().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags()); err != nil {
		return err
	}

	if m.Tags() != nil {
		if err := m.Tags().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt()); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessDetail) validateProcess(formats strfmt.Registry) error {

	if err := validate.Required("process", "body", m.Process); err != nil {
		return err
	}

	if m.Process != nil {
		if err := m.Process.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("process")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this posture check process detail based on the context it is used
func (m *PostureCheckProcessDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckProcessDetail) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Links().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("_links")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("_links")
		}
		return err
	}

	return nil
}

func (m *PostureCheckProcessDetail) contextValidateRoleAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleAttributes() != nil {
		if err := m.RoleAttributes().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckProcessDetail) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if m.Tags() != nil {
		if err := m.Tags().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckProcessDetail) contextValidateProcess(ctx context.Context, formats strfmt.Registry) error {

	if m.Process != nil {
		if err := m.Process.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("process")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckProcessDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckProcessDetail) UnmarshalBinary(b []byte) error {
	var res PostureCheckProcessDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
