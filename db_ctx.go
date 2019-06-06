package gorm

import "context"

func (s *DB) WithContext(ctx context.Context) *DB {
	db := s.clone()
	db.ctx = ctx
	return db
}

func (s *DB) contextOrBackground() context.Context {
	if s.ctx != nil {
		return s.ctx
	}
	return context.Background()
}
