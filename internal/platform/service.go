package platform

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
)

func ListPlatforms(showInactive bool) (platforms []models.Platform, err error) {
	platforms, err = GetPlatformListFromDb(showInactive)
	if err != nil {
		return nil, fmt.Errorf("failed to get platform list: %w", err)
	}

	return platforms, nil
}

func AddPlatform(platform models.Platform) (err error) {
	err = AddPlatformToDb(platform)
	if err != nil {
		return fmt.Errorf("failed to add platform: %w", err)
	}
	return nil
}

func RemovePlatform(platform models.Platform) (err error) {
	err = RemovePlatformFromDb(platform)
	if err != nil {
		return fmt.Errorf("failed to remove platform: %w", err)
	}
	return nil
}
