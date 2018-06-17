package message

type LdapError struct {
	Msg string
}

func (e LdapError) Error() string {
	return e.Msg
}
