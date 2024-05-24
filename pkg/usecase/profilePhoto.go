package usecase

import (
	"fmt"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/helper"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
	"github.com/google/uuid"
)

func (ur *userUseCase) UpdateProfilePhoto(images models.UserProfilePhoto) (models.ProfilePhoto, error) {
	uploadedUrl := []string{}

	for _, imageData := range images.ImageData {
		// Generate a unique file name
		fileUID := uuid.New()
		fileName := fileUID.String()

		// Upload to S3
		url, err := helper.AddImageToAwsS3(imageData, fileName)
		if err != nil {
			fmt.Println("error uploading to S3:", err)
			return models.ProfilePhoto{}, err
		}

		uploadedUrl = append(uploadedUrl, url)
	}

	images.ImageURL = uploadedUrl
	err := ur.userRepository.SaveProfilePhoto(uint32(images.UserID), images.ImageURL)
	if err != nil {
		return models.ProfilePhoto{}, err
	}

	return models.ProfilePhoto{
		UserID:   images.UserID,
		ImageURL: images.ImageURL,
	}, nil
}

func (uc *userUseCase) AddProfilePhoto(userID uint32, imageData []byte) (models.ProfilePhoto, error) {
	if uc.profilePhotoRepo.MaxPhotosReached(userID) {
		return models.ProfilePhoto{}, fmt.Errorf("maximum number of profile photos reached")
	}

	// Generate a unique file name
	fileUID := uuid.New()
	fileName := fileUID.String()

	// Upload to S3
	url, err := helper.AddImageToAwsS3(imageData, fileName)
	if err != nil {
		return models.ProfilePhoto{}, fmt.Errorf("error uploading to S3: %v", err)
	}

	profilePhotos, err := uc.profilePhotoRepo.GetProfilePhotos(userID)
	if err != nil {
		return models.ProfilePhoto{}, fmt.Errorf("error retrieving profile photos: %v", err)
	}

	profilePhotos = append(profilePhotos, url)

	err = uc.userRepository.SaveProfilePhoto(userID, profilePhotos)
	if err != nil {
		return models.ProfilePhoto{}, fmt.Errorf("error saving profile photos: %v", err)
	}

	return models.ProfilePhoto{
		UserID:   int(userID),
		ImageURL: profilePhotos,
	}, nil
}

func (uc *userUseCase) DeleteProfilePhotoByID(userID uint32, imageURL string) (models.ProfilePhoto, error) {
	profilePhotos, err := uc.profilePhotoRepo.GetProfilePhotos(userID)
	if err != nil {
		return models.ProfilePhoto{}, fmt.Errorf("error retrieving profile photos: %v", err)
	}

	index := -1
	for i, url := range profilePhotos {
		if url == imageURL {
			index = i
			break
		}
	}

	if index == -1 {
		return models.ProfilePhoto{}, fmt.Errorf("photo not found")
	}

	err = uc.profilePhotoRepo.DeleteProfilePhoto(userID, index)
	if err != nil {
		return models.ProfilePhoto{}, fmt.Errorf("error deleting profile photo: %v", err)
	}

	profilePhotos = append(profilePhotos[:index], profilePhotos[index+1:]...)

	return models.ProfilePhoto{
		UserID:   int(userID),
		ImageURL: profilePhotos,
	}, nil
}
