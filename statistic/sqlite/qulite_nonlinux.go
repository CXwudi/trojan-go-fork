package sqlite

import "github.com/Potterli20/trojan-go-fork/statistic"

const Name = "sqlite"

type Persistencer struct{}

func NewSqlitePersistencer(path string) (*Persistencer, error) {
	return &Persistencer{}, nil
}

func (p *Persistencer) SaveUser(u statistic.Metadata) error {
	return nil
}

func (p *Persistencer) LoadUser(hash string) (statistic.Metadata, error) {
	var u User
	return &u, nil
}

func (p *Persistencer) DeleteUser(hash string) error {
	return nil
}

func (p *Persistencer) ListUser(f func(hash string, u statistic.Metadata) bool) error {
	return nil
}

func (p *Persistencer) UpdateUserTraffic(hash string, sent, recv uint64) error {
	return nil
}
