/*
 * Copyright 2021 DaSCH - Data and Service Center for the Humanities.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package valueobject

import (
	"fmt"
	"strconv"
)

type ShortCode struct {
	value string
}

// NewShortCode creates a new valid short code object.
func NewShortCode(value string) (ShortCode, error) {
	_, err := strconv.ParseUint(value, 16, 64)
	if err != nil {
		return ShortCode{}, fmt.Errorf("invalid short code, must be two hexadecimal digits and non-empty")
	}

	return ShortCode{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v ShortCode) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v ShortCode) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *ShortCode) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewShortCode(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v ShortCode) Equals(value Value) bool {
	otherValueObject, ok := value.(ShortCode)
	return ok && v.value == otherValueObject.value
}
