package usecases

import (
    "context"

    "github.com/018bf/example/internal/domain/models"
)

//nolint: lll
//go:generate mockgen -build_flags=-mod=mod -destination mock/arch_mock.go github.com/018bf/example/internal/domain/usecases ArchUseCase

type ArchUseCase interface {
    Get(
        ctx context.Context,
        id models.UUID,
    ) (*models.Arch, error)
    List(
        ctx context.Context,
        filter *models.ArchFilter,
    ) ([]*models.Arch, uint64, error)
    Create(
        ctx context.Context,
        create *models.ArchCreate,
    ) (*models.Arch, error)
    Update(
        ctx context.Context,
        update *models.ArchUpdate,
    ) (*models.Arch, error)
    Delete(
        ctx context.Context,
        id models.UUID,
    ) error
}
