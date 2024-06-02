package interfaces

type ProfilePhotoRepo interface {
	GetProfilePhotos(userID uint32) ([]string, error)
	SetMainPhoto(userID uint32, mainPhotoURL string) error
	MaxPhotosReached(userID uint32) bool
	DeleteProfilePhoto(userID uint32, index int) error
}
