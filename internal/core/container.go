// internal/core/repositories.go
package core

import (
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/repositories"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	ReportBuilder repositories.ReportBuilderRepository
	FilterBuilder repositories.FilterBuilderRepository
	LookupBranch  repositories.LookupBranchRepository
	// ... tambahkan repo lain di sini
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {

	return &RepositoryContainer{
		ReportBuilder: repositories.NewReportBuilderRepositories(db),
		FilterBuilder: repositories.NewFilterBuilderRepository(db),
		LookupBranch:  repositories.NewLookupBranchRepository(db),
	}
}
