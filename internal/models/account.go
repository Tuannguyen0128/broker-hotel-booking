package models

type Account struct {
	Id          string `json:"id,omitempty"`
	StaffId     string `json:"staff_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserRoleId  string `json:"user_role_id"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	DeletedAt   string `json:"deleted_at,omitempty"`
	LastLoginAt string `json:"last_login_at,omitempty"`
}
