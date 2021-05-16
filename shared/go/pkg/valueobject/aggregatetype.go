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

import "fmt"

type AggregateType struct {
	value string
}

// NewAggregateType creates a new valid aggregate type value object.
func NewAggregateType(value string) (AggregateType, error) {
	if len(value) > 50 || value == "" {
		return AggregateType{}, fmt.Errorf("invalid aggregate type, must be within 50 characters and non-empty")
	}
	return AggregateType{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v AggregateType) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v AggregateType) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *AggregateType) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewAggregateType(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v AggregateType) Equals(value Value) bool {
	otherValueObject, ok := value.(AggregateType)
	return ok && v.value == otherValueObject.value
}
