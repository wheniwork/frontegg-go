package user_types

type UpdateUserRequest struct {
	PhoneNumber       *string `json:"phoneNumber,omitempty"`
	ProfilePictureURL *string `json:"profilePictureUrl,omitempty"`
	Metadata          *string `json:"metadata,omitempty"`
	VendorMetadata    *string `json:"vendorMetadata,omitempty"`
	MfaBypass         *bool   `json:"mfaBypass,omitempty"`
	Name              *string `json:"name,omitempty"`
}
