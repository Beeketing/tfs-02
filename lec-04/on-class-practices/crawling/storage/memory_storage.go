package storage

type MemoryStorage struct {
	m map[string][]byte
}

func (s *MemoryStorage) Save(url string, b []byte) error {
	if _, ok := s.m[url]; !ok {
		s.m[url] = b
	}
	return nil
}
