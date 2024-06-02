package repository

import (
	"context"
	"fmt"
	"time"

	interfaces "github.com/bibin-zoz/social-match-userauth-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type profilePhotoRepo struct {
	DB      *gorm.DB
	MongoDB *mongo.Database
}

func NewProfilePhotoRepo(DB *gorm.DB, MongoDB *mongo.Database) interfaces.ProfilePhotoRepo {
	return &profilePhotoRepo{
		DB:      DB,
		MongoDB: MongoDB,
	}

}

func (ur *profilePhotoRepo) GetProfilePhotos(userID uint32) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result struct {
		ProfilePhotos []string `bson:"profile_photos"`
	}

	filter := bson.M{"userid": userID}
	err := ur.MongoDB.Collection("userimages").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile photos: %v", err)
	}

	return result.ProfilePhotos, nil
}

func (ur *profilePhotoRepo) DeleteProfilePhoto(userID uint32, index int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"userid": userID}

	var result struct {
		ProfilePhotos []string `bson:"profile_photos"`
	}

	err := ur.MongoDB.Collection("userimages").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to get profile photos: %v", err)
	}

	if index < 0 || index >= len(result.ProfilePhotos) {
		return fmt.Errorf("index out of range")
	}

	result.ProfilePhotos = append(result.ProfilePhotos[:index], result.ProfilePhotos[index+1:]...)

	update := bson.M{
		"$set": bson.M{"profile_photos": result.ProfilePhotos},
	}

	_, err = ur.MongoDB.Collection("userimages").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete profile photo: %v", err)
	}

	return nil
}

func (ur *profilePhotoRepo) SetMainPhoto(userID uint32, mainPhotoURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"userid": userID}

	update := bson.M{
		"$set": bson.M{"main_photo": mainPhotoURL},
	}

	_, err := ur.MongoDB.Collection("userimages").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to set main photo: %v", err)
	}

	return nil
}

func (ur *profilePhotoRepo) MaxPhotosReached(userID uint32) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"userid": userID}
	var result struct {
		ProfilePhotos []string `bson:"profile_photos"`
	}

	err := ur.MongoDB.Collection("userimages").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false
	}

	return len(result.ProfilePhotos) >= 5
}
