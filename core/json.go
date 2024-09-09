package core

type JSON interface {
	UnmarshalJSON([]byte) error
	MarshalJSON() ([]byte, error)
}

type MarshalJSON interface {
	MarshalJSON() ([]byte, error)
}
