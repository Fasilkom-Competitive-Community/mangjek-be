package user

const (
	USER    uint32 = 1
	DEFAULT uint32 = 2
	ADMIN   uint32 = 3
)

var (
	roleDirectories map[string]uint32 = map[string]uint32{
		"user":  USER,
		"":      DEFAULT,
		"admin": ADMIN,
	}
)
