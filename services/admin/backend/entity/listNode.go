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

import "time"

type ListNode struct {
	ID        ID         `json:"id"`
	Name      string     `json:"name"`
	Label     string     `json:"label"`
	Comment   string     `json:"comment"`
	RootNode  string     `json:"rootNode"`
	Position  int        `json:"position"`
	Children  []ListNode `json:"children"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

//NewList create a new ListNode entity
func NewListNode(name string, label string, comment string) (*ListNode, error) {
	node := &ListNode{
		ID:        NewID(),
		Name:      name,
		Label:     label,
		Comment:   comment,
		CreatedAt: time.Now(),
	}

	err := node.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return node, nil
}

//Validate validate ListNode entity
func (node *ListNode) Validate() error {
	if node.Name == "" || node.Label == "" {
		return ErrInvalidEntity
	}

	return nil
}
