package usecase

import (
	"yamato/internal/domain"
	"yamato/internal/repository"
)

type MountainUseCase interface {
	FetchMountainByID(id int) (*domain.Mountain, error)
	FetchAllMountains() ([]*domain.Mountain, error)
}

type mountainUseCase struct {
	repo repository.MountainRepository
}

func NewMountainUseCase(repo repository.MountainRepository) MountainUseCase {
	return &mountainUseCase{repo: repo}
}

func (uc *mountainUseCase) FetchMountainByID(id int) (*domain.Mountain, error) {
	return uc.repo.GetMountainByID(id)
}

func (uc *mountainUseCase) FetchAllMountains() ([]*domain.Mountain, error) {
	return uc.repo.GetAllMountains()
}
