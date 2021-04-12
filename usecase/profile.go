package usecase

import (
	"errors"
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

	// 該当ユーザーない時はRepositoryはエラー扱いにせずnilを返す。エラー扱いにするのはUseCase(ユースケース都合でエラーにしてるから)
	if profile == nil {
		return nil, errors.New("該当するユーザーがいません")
	}

	return profile, nil
}

func (pu profileUseCase) AddProfile(p *model.Profile) (err error) {
	// 必要に応じてここでバリデーション
	// 重複しないかチェック
	name := p.Name
	profile, err := pu.profileRepository.GetProfile(name)
	if err != nil {
		return err
	}

	if profile != nil {
		return errors.New("既に存在する名前です")
	}

	err = pu.profileRepository.AddProfile(p)
	return err
}
