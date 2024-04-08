package service

import (
	"encoding/json"
	"game-store-api/database"
)

type PlatformService struct {
	platform database.IPlatform
}

type IPlatformService interface {
	GetPlatforms() ([]byte, error)
	GetPlatform(id string) ([]byte, error)
	CreatePlatform(platform database.Platform) ([]byte, error)
	UpdatePlatform(platform database.Platform) ([]byte, error)
	DeletePlatform(id string) ([]byte, error)
}

func NewPlatformService(platform database.IPlatform) *PlatformService {
	return &PlatformService{platform: platform}
}

func (s *PlatformService) GetPlatforms() ([]byte, error) {
	platforms, err := s.platform.GetPlatforms()
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(platforms)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *PlatformService) GetPlatform(id string) ([]byte, error) {
	platform, err := s.platform.GetPlatform(id)
	if err != nil {
		return nil, err
	}

	jsonM, err := json.Marshal(platform)
	if err != nil {
		return nil, err
	}

	return jsonM, nil
}

func (s *PlatformService) CreatePlatform(platform database.Platform) ([]byte, error) {
	err := s.platform.CreatePlatform(platform)
	if err != nil {
		return nil, err
	}

	return []byte("Platform created"), nil
}

func (s *PlatformService) UpdatePlatform(platform database.Platform) ([]byte, error) {
	err := s.platform.UpdatePlatform(platform)
	if err != nil {
		return nil, err
	}

	return []byte("Platform updated"), nil
}

func (s *PlatformService) DeletePlatform(id string) ([]byte, error) {
	err := s.platform.DeletePlatform(id)
	if err != nil {
		return nil, err
	}

	return []byte("Platform deleted"), nil
}