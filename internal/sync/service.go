package sync

import "github.com/gin-gonic/gin"

type Service struct {
	repo    *Repository
	rasRepo *RASRepository
}

func NewService(repo *Repository, rasRepo *RASRepository) *Service {
	return &Service{
		repo:    repo,
		rasRepo: rasRepo,
	}
}

func (s *Service) SyncProformas(c *gin.Context) error {
	return nil
}

func (s *Service) SyncStudents(c *gin.Context) error {
	return nil
}
