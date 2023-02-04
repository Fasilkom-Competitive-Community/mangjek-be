package user

func (m AuthDriver) IsAdmin() bool {
	return m.Role >= ADMIN
}

func (m AuthDriver) IsUser() bool {
	return m.Role == USER
}

func (m AuthDriver) IsSame(ID string) bool {
	return m.ID == ID
}

func (m *AuthDriver) SetRoleString(role string) {
	m.Role = roleDirectories[role]
}
