package usecase

import (
	"go-echo-practice/domain/model"
	"go-echo-practice/domain/repository"
)

// ProfileUseCase は ProfileのUseCaseのインターフェイス、Repositoryと同じものを記述
type ProfileUseCase interface {
	GetProfile(name string) (*model.Profile, error)
	AddProfile(p *model.Profile) error
}

// UseCaseは自分に対応するリポジトリを知っている
type profileUseCase struct {
	profileRepository repository.ProfileRepository
}

// NewProfileUseCase は ProfileUseCaseのコンストラクタ
func NewProfileUseCase(pr repository.ProfileRepository) ProfileUseCase {
	return &profileUseCase{
		profileRepository: pr,
	}
}

// インターフェイスの実装
func (pu profileUseCase) GetProfile(name string) (profile *model.Profile, err error) {
	profile, err = pu.profileRepository.GetProfile(name)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (pu profileUseCase) AddProfile(p *model.Profile) (err error) {
	// 必要に応じてここでバリデーション

	err = pu.profileRepository.AddProfile(p)
	return err
}
