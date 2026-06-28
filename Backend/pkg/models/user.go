package models

import "time"

// ID (string, maps to id in DB)
// Email (string, maps to email in DB)
// PasswordHash (string, maps to password_hash in DB)
// WalletAddress (pointer to string *string, maps to wallet_address in DB. Using a pointer allows handling SQL NULL values cleanly in Go)
// CreatedAt (time.Time, maps to created_at in DB)

type User struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"-"`
	WalletAddress *string   `json:"wallet_address"`
	CreatedAt     time.Time `json:"created_at"`
}

/*
Define Request Structs:

Define SignupRequest: A struct representing the payload when a user registers. It should contain Email and Password fields with appropriate JSON tags.
Define LoginRequest: A struct representing the payload when a user logs in. It should contain Email and Password fields with appropriate JSON tags.
*/

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Define Response Structs:
Define UserResponse: A struct representing the user data returned to the client. This struct must exclude the PasswordHash for security reasons. It should contain ID, Email, WalletAddress, and CreatedAt with appropriate JSON tags.
*/

type UserResponse struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	WalletAddress *string   `json:"wallet_address,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}
