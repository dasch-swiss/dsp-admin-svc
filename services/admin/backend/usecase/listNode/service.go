package listNode

import "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"

//Service interface
type Service struct {
	repo Repository
}

//NewService create a new listNode use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateListNode create a ListNode
func (s *Service) CreateListNode(label string, comment string) (entity.ID, error) {
	e, err := entity.NewListNode(label, comment)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetListNode get a ListNode by id
func (s *Service) GetListNode(id entity.ID) (*entity.ListNode, error) {
	return s.repo.Get(id)
}
