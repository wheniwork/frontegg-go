package identity_types

type IdentityConfiguration struct {
	ID                                  string `json:"id"`
	DefaultTokenExpiration              int    `json:"defaultTokenExpiration"`
	DefaultRefreshTokenExpiration       int    `json:"defaultRefreshTokenExpiration"`
	PublicKey                           string `json:"publicKey"`
	CookieSameSite                      string `json:"cookieSameSite"`
	AllowSignups                        bool   `json:"allowSignups"`
	APITokensEnabled                    bool   `json:"apiTokensEnabled"`
	AllowOverridePasswordComplexity     bool   `json:"allowOverridePasswordComplexity"`
	AllowOverridePasswordExpiration     bool   `json:"allowOverridePasswordExpiration"`
	AllowOverrideEnforcePasswordHistory bool   `json:"allowOverrideEnforcePasswordHistory"`
	JwtAlgorithm                        string `json:"jwtAlgorithm"`
	JwtSecret                           string `json:"jwtSecret"`
	AllowNotVerifiedUsersLogin          bool   `json:"allowNotVerifiedUsersLogin"`
	ForcePermissions                    bool   `json:"forcePermissions"`
	AuthStrategy                        string `json:"authStrategy"`
	DefaultPasswordlessTokenExpiration  int    `json:"defaultPasswordlessTokenExpiration"`
	ForceSameDeviceOnAuth               bool   `json:"forceSameDeviceOnAuth"`
	AllowTenantInvitations              bool   `json:"allowTenantInvitations"`
	RotateRefreshTokens                 bool   `json:"rotateRefreshTokens"`
	MachineToMachineAuthStrategy        string `json:"machineToMachineAuthStrategy"`
	AddRolesToJwt                       bool   `json:"addRolesToJwt"`
	AddPermissionsToJwt                 bool   `json:"addPermissionsToJwt"`
	RefreshTokensRotationLimit          int    `json:"refreshTokensRotationLimit"`
	AddSamlAttributesToJwt              bool   `json:"addSamlAttributesToJwt"`
}
