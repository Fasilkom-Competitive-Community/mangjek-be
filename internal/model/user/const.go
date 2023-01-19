package user

const (
	DEFAULT uint32 = 1
	DRIVER  uint32 = 2
	ADMIN   uint32 = 3
)

var (
	roleDirectories map[string]uint32 = map[string]uint32{
		"":       DEFAULT,
		"driver": DRIVER,
		"admin":  ADMIN,
	}
)
