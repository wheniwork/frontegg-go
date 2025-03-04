package group_types

import (
	"time"

	"github.com/wheniwork/frontegg-go/clients/roles/role_types"
)

type GroupUser struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	Name               string    `json:"name"`
	ProfilePictureURL  string    `json:"profilePictureUrl"`
	CreatedAt          time.Time `json:"createdAt"`
	ActivatedForTenant bool      `json:"activatedForTenant"`
}

type Group struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Color       string            `json:"color"`
	Description string            `json:"description"`
	Metadata    string            `json:"metadata"`
	Roles       []role_types.Role `json:"roles"`
	Users       []GroupUser       `json:"users"`
	ManagedBy   string            `json:"managedBy"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}
