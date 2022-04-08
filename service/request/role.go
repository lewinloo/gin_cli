package request

type CreateRoleParams struct {
	RoleId        string `json:"roleId"`
	RoleName      string `json:"roleName"`
	ParentId      string `json:"pid"`
	DefaultRouter string `json:"defaultRouter"`
}
