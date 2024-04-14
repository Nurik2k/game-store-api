package service

import (
	"game-store-api/database"
)

type PlatformService struct {
	platform database.IPlatform
}

type IPlatformService interface {
	GetPlatforms() ([]database.Platform, error)
	GetPlatform(id int) (database.Platform, error)
	CreatePlatform(platform database.Platform) (string, error)
	UpdatePlatform(platform database.Platform) (string, error)
	DeletePlatform(id int) (string, error)
}

func NewPlatformService(platform database.IPlatform) *PlatformService {
	return &PlatformService{platform: platform}
}

func (s *PlatformService) GetPlatforms() ([]database.Platform, error) {
	platforms, err := s.platform.GetPlatforms()
	if err != nil {
		return nil, err
	}

	return platforms, nil
}

func (s *PlatformService) GetPlatform(id int) (database.Platform, error) {
	platform, err := s.platform.GetPlatform(id)
	if err != nil {
		return platform, err
	}

	return platform, nil
}

func (s *PlatformService) CreatePlatform(platform database.Platform) (string, error) {
	err := s.platform.CreatePlatform(platform)
	if err != nil {
		return "", err
	}

	return "Platform created", nil
}

func (s *PlatformService) UpdatePlatform(platform database.Platform) (string, error) {
	err := s.platform.UpdatePlatform(platform)
	if err != nil {
		return "", err
	}

	return"Platform updated", nil
}

func (s *PlatformService) DeletePlatform(id int) (string, error) {
	err := s.platform.DeletePlatform(id)
	if err != nil {
		return "", err
	}

	return "Platform deleted", nil
}