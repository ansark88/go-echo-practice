package repository

import (
	"go-echo-practice/domain/model"
)

// ProfileRepository は ProfileにおけるRepositoryのインターフェイス
type ProfileRepository interface {
	ListProfile(query string) (*[]model.Profile, error)
	GetProfile(name string) (*model.Profile, error)
	AddProfile(p *model.Profile) error
}

// ProfileRepository.goでは実装を行わない。実装するのはpersistenceのほう
