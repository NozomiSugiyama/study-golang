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
				{1, "Nozomi Sugiyama", "sugiyama@example.com", "1970/1/1", "XXX-XXXX-XXXX"},
				{2, "Hoge Huga", "hoge@example.com", "1992/03/16", "YYY-YYYY-YYYY"},
				{2, "Foo Bar", "foo@example.com", "1977/06/25", "ZZZ-ZZZZ-ZZZZ"},
			},
		},
	}
}

// GetSession get session
func GetSession() *session {
	return instance
}
