package sync

import (
	"context"
	"fmt"

	"time"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

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

func (s *Service) SyncProformas(ctx context.Context) (err error) {
	start := time.Now()
	recordsCount := 0
	status := "success"
	logMessage := ""

	defer func() {
		logErr := s.repo.CreateSyncLog(ctx, &database.SyncLog{
			EntityType:   "recruitment_cycles",
			ExternalID:   "active",
			Action:       database.SyncAction("synced"),
			RecordsCount: recordsCount,
			Status:       status,
			ErrorMessage: logMessage,
			SyncDuration: int(time.Since(start).Milliseconds()),
		})

		if logErr != nil && err == nil {
			err = logErr
		}
	}()

	rcs, err := s.rasRepo.GetActiveRecruitmentCycles(ctx)

	if err != nil {
		status = "failed"
		logMessage = err.Error()
		return err
	}

	for _, rasRc := range rcs {
		pibsRc := mapRecruitmentCycle(rasRc)

		if err := s.repo.UpsertRecruitmentCycle(ctx, &pibsRc); err != nil {
			status = "failed"
			logMessage = fmt.Sprintf("recruitment cycle %d: %v", rasRc.ID, err)
			return err
		}

		recordsCount++
	}

	return nil
}

func (s *Service) SyncStudents(ctx context.Context) (err error) {
	start := time.Now()
	status := "success"
	logMessage := ""

	defer func() {
		logErr := s.repo.CreateSyncLog(ctx, &database.SyncLog{
			EntityType:   "students",
			ExternalID:   "bulk",
			Action:       database.SyncAction("synced"),
			RecordsCount: 0,
			Status:       status,
			ErrorMessage: logMessage,
			SyncDuration: int(time.Since(start).Milliseconds()),
		})

		if logErr != nil && err == nil {
			err = logErr
		}
	}()

	return nil
}

func mapRecruitmentCycle(rc RASRecruitmentCycle) database.RecruitmentCycle {
	return database.RecruitmentCycle{
		ID:           rc.ID,
		AcademicYear: rc.AcademicYear,
		Type:         rc.Type,
		Phase:        rc.Phase,
		IsActive:     rc.IsActive,
	}
}

func mapProforma(p RASProforma) database.Proforma {
	return database.Proforma{
		ID:                 p.ID,
		RecruitmentCycleID: p.RecruitmentCycleID,

		CompanyID: p.CompanyID,

		Title:       p.CompanyName,
		RoleOffered: p.Role,
		Description: p.Profile,

		ProformaType:      "",
		IsInterviewActive: false,
		LastSyncedAt:      time.Now(),
	}
}
