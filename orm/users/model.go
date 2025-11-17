// should contain User struct, role types and validation
package users

import "slices"

/*
User struct contains an ID, Name Surname, StudyGroup and a Role
(e.g chair, representative, vice-president)
*/
type User struct {
	ID         int
	Name       string
	Surname    string
	StudyGroup string
	Role       Role // Change User.Role from string to Role so the struct uses Role type
}

type Role string

const (
	RoleAdmin          Role = "admin"
	RoleChair          Role = "chair"
	RoleRepresentative Role = "representative"
	RoleVicePresident  Role = "vicepresident"
)

// IsOneOf returns true if r equals any of the provided roles.
func (r Role) IsOneOf(roles ...Role) bool {
	return slices.Contains(roles, r)
}

// Permissions part begin

const (
	PermManageUsers    = "manage:users"
	PermCreateMeetings = "create:meetings"
	PermManageVoting   = "manage:voting"
)

// rolePermission maps each role to permission set.
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

// Purpose: determine whether a given Role has a specific permission string.
// 1) Assigning all permissions regarding given role to `perms`
// 2) After this checking if perms don't exist in the rolePermission constants
// 3) Return true/false from the rolePermission constants, by the needed permit from (perm string) in line 65
func (r Role) HasPermission(perm string) bool {
	perms := rolePermission[r] // perms now contains permissions for the 'r' - role
	if perms == nil {
		return false
	}
	return perms[perm]
}
