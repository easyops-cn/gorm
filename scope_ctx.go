package gorm

import "database/sql"

func (scope *Scope) sqldbExec(query string, args ...interface{}) (sql.Result, error) {
	return scope.SQLDB().ExecContext(scope.db.contextOrBackground(), query, args...)
}

func (scope *Scope) sqldbPrepare(query string) (*sql.Stmt, error) {
	return scope.SQLDB().PrepareContext(scope.db.contextOrBackground(), query)
}

func (scope *Scope) sqldbQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return scope.SQLDB().QueryContext(scope.db.contextOrBackground(), query, args...)
}

func (scope *Scope) sqldbQueryRow(query string, args ...interface{}) *sql.Row {
	return scope.SQLDB().QueryRowContext(scope.db.contextOrBackground(), query, args...)
}
