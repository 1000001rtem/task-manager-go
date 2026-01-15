package command

type Command struct {
	Name        string
	Description string
	Secure      bool
	Action      func() error
}
