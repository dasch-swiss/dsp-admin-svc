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
	"time"
)

type Timestamp struct {
	value time.Time
}

//NewTimestamp creates a new timestamp value object
func NewTimestamp() Timestamp {
	v := time.Now()
	return Timestamp{value: v}
}

//NewTimestampFromUnix creates a new timestamp value object for the supplied unix time in seconds.
func NewTimestampFromUnix(sec int64) Timestamp {
	return Timestamp{value: time.Unix(sec, 0)}
}

//Time returns the Time of the value.
func (v Timestamp) Time() time.Time {
	return v.value
}

//Unix returns the Unix time of the value in number of seconds.
func (v Timestamp) Unix() int64 {
	return v.value.Unix()
}

// String implements the fmt.Stringer interface.
func (v Timestamp) String() string {
	return v.value.String()
}

// MarshalText used to serialize the object
func (v Timestamp) MarshalText() ([]byte, error) {
	ts := v.value.UnixNano()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *Timestamp) UnmarshalText(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	// we marshal it as nanoseconds
	*v = NewTimestampFromUnix(int64(ts))
	return nil
}

// MarshalJSON used to serialize the object
func (v Timestamp) MarshalJSON() ([]byte, error) {
	ts := v.value.Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

// UnmarshalJSON used to deserialize the object and returns an error if it's invalid.
func (v *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	// we marshal it as nanoseconds
	*v = NewTimestampFromUnix(int64(ts))
	return nil
}

//Equals tests for equality with another value object
func (v Timestamp) Equals(value Value) bool {
	otherIdentifier, ok := value.(Timestamp)
	return ok && v.value == otherIdentifier.value
}
