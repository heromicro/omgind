package uid

type IDProvider interface {
	ID() (string, error)
}
