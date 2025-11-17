package orm

import "slices"

type Role string

const (
	RoleAdmin          Role = "admin"
	RoleChair          Role = "chair"
	RoleRepresentative Role = "representative"
	RoleVicePresident  Role = "vicepresident"
)

var AllRoles = []Role{
	RoleAdmin,
	RoleChair,
	RoleRepresentative,
	RoleVicePresident,
}

func (r Role) Valid() bool {
	return slices.Contains(AllRoles, r)
}

const (
	PermManageUsers    = "manage:users"
	PermCreateMeetings = "create:meetings"
	PermManageVoting   = "manage:voting"
)

var rolePermission = map[Role]map[string]bool{
	RoleAdmin: {
		PermManageUsers:    true,
		PermCreateMeetings: true,
		PermManageVoting:   true,
	},
	RoleChair: {
		PermManageUsers:    false,
		PermCreateMeetings: true,
		PermManageVoting:   true,
	},
	RoleRepresentative: {
		PermManageUsers:    false,
		PermCreateMeetings: false,
		PermManageVoting:   false,
	},
	RoleVicePresident: {
		PermManageUsers:    false,
		PermCreateMeetings: true,
		PermManageVoting:   true,
	},
}

func (r Role) HasPermission(perm string) bool {
	perms := rolePermission[r]
	if perms == nil {
		return false
	}
	return perms[perm]
}
