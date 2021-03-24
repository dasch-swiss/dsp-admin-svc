package entity_test

import (
	"testing"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewListNode(t *testing.T) {
	node, err := entity.NewListNode("TEST label", "TEST comment")
	assert.Nil(t, err)
	assert.Equal(t, node.Label, "TEST label")
	assert.NotNil(t, node.ID)
	assert.False(t, node.CreatedAt.IsZero())
	assert.True(t, node.UpdatedAt.IsZero())
}
