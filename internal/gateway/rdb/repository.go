package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go-layered-architecture-sample/internal/config"
	"go-layered-architecture-sample/pkg/logger"

	"github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db     *sqlx.DB
	logger logger.Logger
	conf   *config.Database
}

func NewRepository(conf *config.Database, logger logger.Logger) (*Repository, *Repository, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, nil, err
	}

	primaryConfig := mysql.Config{
		DBName:    conf.DatabaseName,
		User:      conf.User,
		Passwd:    conf.Password,
		Addr:      fmt.Sprintf("%s:%d", conf.PrimaryHost, conf.PrimaryPort),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	primaryDB, err := sqlx.Open("mysql", primaryConfig.FormatDSN())
	if err != nil {
		return nil, nil, err
	}

	replicaConfig := mysql.Config{
		DBName:    conf.DatabaseName,
		User:      conf.User,
		Passwd:    conf.Password,
		Addr:      fmt.Sprintf("%s:%d", conf.ReplicaHost, conf.ReplicaPort),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	replicaDB, err := sqlx.Open("mysql", replicaConfig.FormatDSN())
	if err != nil {
		return nil, nil, err
	}

	return &Repository{
			db:     primaryDB,
			logger: logger,
			conf:   conf,
		}, &Repository{
			db:     replicaDB,
			logger: logger,
			conf:   conf,
		}, nil
}

func (r *Repository) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	r.executeWithLogging(ctx, "QueryxContext", query, func() {
		rows, err = r.db.QueryxContext(ctx, query, args...)
	})
	return rows, err
}

func (r *Repository) QueryRowxContext(ctx context.Context, query string, args ...interface{}) (row *sqlx.Row) {
	r.executeWithLogging(ctx, "QueryRowxContext", query, func() {
		row = r.db.QueryRowxContext(ctx, query, args...)
	})
	return row
}

func (r *Repository) NamedExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	r.executeWithLogging(ctx, "NamedExecContext", query, func() {
		result, err = r.db.NamedExecContext(ctx, query, args)
	})
	return result, err
}

func (r *Repository) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	r.executeWithLogging(ctx, "ExecContext", query, func() {
		result, err = r.db.ExecContext(ctx, query, args...)
	})
	return result, err
}

func (r *Repository) executeWithLogging(ctx context.Context, action string, query string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	r.logger.Printf("SQL %s: %s; Duration: %v", action, query, duration)
}

func (r *Repository) WithTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				r.logger.Errorf("failed to rollback transaction: %v\n", rollbackErr)
			}
			panic(p)

		} else if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				r.logger.Errorf("failed to rollback transaction: %v\n", rollbackErr)
			}

		} else {
			commitErr := tx.Commit()
			if commitErr != nil {
				r.logger.Errorf("Failed to commit transaction: %v\n", commitErr)
			}
		}
	}()

	if err := fn(tx); err != nil {
		return err
	}

	return nil
}

func NewBaseStatementBuilder() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)
}
