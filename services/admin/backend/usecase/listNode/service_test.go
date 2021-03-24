package listNode

import (
	"testing"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureListNode() *entity.ListNode {
	return &entity.ListNode{
		ID:        entity.NewID(),
		Label:     "TEST node label",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	node := newFixtureListNode()
	_, err := service.CreateListNode(node.Label, node.Comment)
	assert.Nil(t, err)
	assert.False(t, node.CreatedAt.IsZero())
	assert.True(t, node.UpdatedAt.IsZero())
}
