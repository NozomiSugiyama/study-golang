package model

type store struct {
	users []User
}

type session struct {
	store store
}

var instance *session = newSession()

func newSession() *session {
	return &session{
		store{
			[]User{
				{1, "Nozomi Sugiyama", "nozomi", "1970/1/1", "090-6809-3158"},
			},
		},
	}
}

// GetSession get session
func GetSession() *session {
	return instance
}
