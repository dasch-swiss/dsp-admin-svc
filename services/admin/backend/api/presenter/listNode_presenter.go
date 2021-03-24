package presenter

import (
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//ListNode data
type ListNode struct {
	ID        entity.ID  `json:"id"`
	Label     string     `json:"label"`
	Comment   string     `json:"comment"`
	RootNode  string     `json:"rootNode"`
	Position  int        `json:"position"`
	Children  []ListNode `json:"children"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
