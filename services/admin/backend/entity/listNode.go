package entity

import "time"

type ListNode struct {
	ID        ID
	Label     string
	Comment   string
	RootNode  string
	Position  int
	Children  []ListNode
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewList create a new ListNode entity
func NewListNode(label string, comment string) (*ListNode, error) {
	node := &ListNode{
		ID:        NewID(),
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
	if node.Label == "" {
		return ErrInvalidEntity
	}

	return nil
}
