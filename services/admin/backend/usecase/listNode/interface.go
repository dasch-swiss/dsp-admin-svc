package listNode

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.ListNode, error)
	// Search(query string) ([]*entity.ListNode, error)
	// List() ([]*entity.ListNode, error)
}

//Writer interface
type Writer interface {
	Create(e *entity.ListNode) (entity.ID, error)
	// Update(e *entity.ListNode) error
	// Delete(e *entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetListNode(id entity.ID) (*entity.ListNode, error)
	CreateListNode(label string, comment string) (entity.ID, error)
}
