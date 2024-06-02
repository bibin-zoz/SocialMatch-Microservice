package models

type Intrests struct {
	ID   uint
	Name string `json:"interest_name"  binding:"min=4"`
}
type Preferences struct {
	ID   uint
	Name string `json:"preference_name" binding:"min=4"`
}
type AddPreferences struct {
	Name string `json:"preference_name" binding:"min=4"`
}
type AddInterest struct {
	Name string `json:"interest_name"  binding:"min=4"`
}
type AddUserInterestRequest struct {
	UserID     uint64 `json:"user_id"`
	InterestID int    `json:"interest_id"`
}

type EditUserInterestRequest struct {
	UserID          uint64 `json:"user_id"`
	InterestID      uint64 `json:"interest_id" binding:"min=4"`
	NewInterestName string `json:"new_interest_name" binding:"min=4"`
}
type DeleteUserInterestRequest struct {
	UserID     uint64 `json:"user_id"`
	InterestID uint64 `json:"interest_id"`
}

type AddUserPreferenceRequest struct {
	UserID       uint64 `json:"user_id"`
	PreferenceID int    `json:"Preference_id"`
}

type EditUserPreferenceRequest struct {
	UserID            uint64 `json:"user_id"`
	PreferenceID      uint64 `json:"Preference_id"`
	NewPreferenceName string `json:"new_Preference_name" binding:"min=4"`
}
type DeleteUserPreferenceRequest struct {
	UserID       uint64 `json:"user_id"`
	PreferenceID uint64 `json:"Preference_id" binding:"min=4"`
}

type InterestResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PreferenceResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
