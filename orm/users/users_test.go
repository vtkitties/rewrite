package users

import "testing"

// TestRoleIsOneOf tests the IsOneOf method with various role combinations
func TestRoleIsOneOf(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		roles    []Role
		expected bool
	}{
		{
			name:     "role matches single item",
			role:     RoleAdmin,
			roles:    []Role{RoleAdmin},
			expected: true,
		},
		{
			name:     "role matches in middle of list",
			role:     RoleChair,
			roles:    []Role{RoleAdmin, RoleChair, RoleRepresentative},
			expected: true,
		},
		{
			name:     "role does not match",
			role:     RoleVicePresident,
			roles:    []Role{RoleAdmin, RoleChair},
			expected: false,
		},
		{
			name:     "empty roles list",
			role:     RoleAdmin,
			roles:    []Role{},
			expected: false,
		},
		{
			name:     "single role matches last in list",
			role:     RoleRepresentative,
			roles:    []Role{RoleAdmin, RoleChair, RoleRepresentative},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.role.IsOneOf(tt.roles...)
			if result != tt.expected {
				t.Errorf("IsOneOf(%v) = %v, want %v", tt.roles, result, tt.expected)
			}
		})
	}
}

// TestRoleHasPermission tests the HasPermission method across roles and permissions
func TestRoleHasPermission(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		perm     string
		expected bool
	}{
		{
			name:     "admin has manage users permission",
			role:     RoleAdmin,
			perm:     PermManageUsers,
			expected: true,
		},
		{
			name:     "chair cannot manage users",
			role:     RoleChair,
			perm:     PermManageUsers,
			expected: false,
		},
		{
			name:     "chair can create meetings",
			role:     RoleChair,
			perm:     PermCreateMeetings,
			expected: true,
		},
		{
			name:     "representative has no permissions",
			role:     RoleRepresentative,
			perm:     PermManageVoting,
			expected: false,
		},
		{
			name:     "unknown permission returns false",
			role:     RoleAdmin,
			perm:     "unknown:permission",
			expected: false,
		},
		{
			name:     "unknown role returns false",
			role:     Role("unknown"),
			perm:     PermManageUsers,
			expected: false,
		},
		{
			name:     "vice president can manage voting",
			role:     RoleVicePresident,
			perm:     PermManageVoting,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.role.HasPermission(tt.perm)
			if result != tt.expected {
				t.Errorf("HasPermission(%q) = %v, want %v", tt.perm, result, tt.expected)
			}
		})
	}
}
