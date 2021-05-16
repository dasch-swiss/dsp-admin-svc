/*
 * Copyright © 2021 the contributors.
 *
 *  This file is part of the DaSCH Service Platform.
 *
 *  The DaSCH Service Platform is free software: you can
 *  redistribute it and/or modify it under the terms of the
 *  GNU Affero General Public License as published by the
 *  Free Software Foundation, either version 3 of the License,
 *  or (at your option) any later version.
 *
 *  The DaSCH Service Platform is distributed in the hope that
 *  it will be useful, but WITHOUT ANY WARRANTY; without even
 *  the implied warranty of MERCHANTABILITY or FITNESS FOR
 *  A PARTICULAR PURPOSE.  See the GNU Affero General Public
 *  License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public
 *  License along with the DaSCH Service Platform.  If not, see
 *  <http://www.gnu.org/licenses/>.
 *
 */

package entity

import (
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"time"
)

//Organization is the domain entity.
type Organization struct {
	ID              ID
	Type            string
	Name            map[string]bool
	Email           valueobject.Email
	URL             valueobject.URL
	PostalAddresses address
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//postalAddress represents the postal address of the organization domain entity.
type address struct {
	StreetAddress   string
	PostalCode      string
	AddressLocality string
}

//NewOrganization creates a new organization entity.
func NewOrganization(newName string) (*Organization, error) {

	name := make(map[string]bool) // idiomatic way of implementing a set
	name[newName] = true

	org := &Organization{
		ID:        NewID(),
		Type:      "http://ns.dasch.swiss/repository#Organization",
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := org.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return org, nil
}

//AddName adds a name to the organization (up to a maximum of 3).
func (org *Organization) AddName(newName string) error {

	// dont allow adding if we already have three
	if len(org.Name) == 3 {
		return ErrCannotAddName
	}

	exists := org.Name[newName]
	if exists {
		// found the name in the list
		// return error because we don't allow adding a name twice
		return ErrCannotAddExistingName
	} else {
		// didn't find it, add to set
		org.Name[newName] = true
	}

	org.UpdatedAt = time.Now()
	return nil
}

//RemoveName removes a name from an organization if more than one is set
func (org *Organization) RemoveName(name string) error {

	// dont allow deleting name if there is only one left
	if len(org.Name) == 1 {
		return ErrCannotDeleteName
	}

	exists := org.Name[name]
	if exists {
		// found it, remove it
		delete(org.Name, name)
	} else {
		// not found, returning error
		return ErrCannotDeleteNotFoundName
	}

	org.UpdatedAt = time.Now()
	return nil
}

//GetNames returns all organization names.
func (org *Organization) GetNames() []string {
	var names []string

	for name := range org.Name {
		names = append(names, name)
	}

	return names
}

//AddAddress adds a postal address to the organization.
func (org *Organization) AddAddress(streetAddress string, postalCode string, addressLocality string) error {

	address := address{
		StreetAddress:   streetAddress,
		PostalCode:      postalCode,
		AddressLocality: addressLocality,
	}

	if address.StreetAddress == "" {
		return ErrInvalidEntity
	}

	if address.PostalCode == "" {
		return ErrInvalidEntity
	}

	if address.AddressLocality == "" {
		return ErrInvalidEntity
	}

	org.PostalAddresses = address
	org.UpdatedAt = time.Now()
	return nil
}

//RemoveAddress removes the postal address from the organization.
func (org *Organization) RemoveAddress() error {
	if org.PostalAddresses == (address{}) {
		return ErrPostalAddressNotSet
	} else {
		org.PostalAddresses = address{}
	}
	org.UpdatedAt = time.Now()
	return nil
}

//AddEmail adds the email address to the organization and overwrites an existing one.
func (org *Organization) AddEmail(emailAddress valueobject.Email) error {
	org.Email = emailAddress
	org.UpdatedAt = time.Now()
	return nil
}

//RemoveEmail removes the email address from the organization.
func (org *Organization) RemoveEmail() error {
	if org.Email == valueobject.ZeroEmail() {
		return ErrEmailNotSet
	} else {
		org.Email = valueobject.ZeroEmail()
	}
	org.UpdatedAt = time.Now()
	return nil
}

//AddURL adds an URL to the organization and overwrites an existing one.
func (org *Organization) AddURL(url valueobject.URL) error {
	org.URL = url
	org.UpdatedAt = time.Now()
	return nil
}

//RemoveURL removes an URL from the organization.
func (org *Organization) RemoveURL() error {
	if org.URL == valueobject.ZeroURL() {
		return ErrURLNotSet
	} else {
		org.URL = valueobject.ZeroURL()
	}
	org.UpdatedAt = time.Now()
	return nil
}

//Validate validates the organization entity.
func (org *Organization) Validate() error {
	if len(org.Name) < 1 {
		return ErrInvalidEntity
	}

	if len(org.Name) > 3 {
		return ErrInvalidEntity
	}

	return nil
}
