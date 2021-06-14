package storage

type Storage interface {
	Save(string, []byte) error
}
