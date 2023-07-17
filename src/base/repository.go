package base

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type key struct {
}

type LogDBKey struct {
}

var LogDBKeyContext = LogDBKey{}
var keyDB = key{}

type Repository struct {
	Db *gorm.DB
}

type IBaseRepository interface {
	BeginTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context
	Commit(ctx context.Context) *gorm.DB
	Rollback(ctx context.Context) *gorm.DB
	Transaction(ctx context.Context, callback func(ctx context.Context) error, opts ...*sql.TxOptions) error
}

func (r *Repository) GetDB(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(keyDB).(*gorm.DB)
	var db *gorm.DB
	if !ok {
		db = r.Db
	} else {
		db = val
	}

	check, ok := ctx.Value(LogDBKeyContext).(bool)
	if ok && check {
		db = db.Session(&gorm.Session{
			Logger: logger.Default.LogMode(logger.Info),
		})

	}
	return db
}

func (r *Repository) BeginTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context {
	db := r.GetDB(ctx)
	return context.WithValue(ctx, keyDB, db.Begin(opts...))
}

func (r *Repository) Commit(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Commit()
}

func (r *Repository) Rollback(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Rollback()
}

func (r *Repository) Transaction(ctx context.Context, callback func(ctx context.Context) error, opts ...*sql.TxOptions) error {
	return r.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		//tx.Set("gorm:query_option", "FOR UPDATE")
		txCtx := context.WithValue(ctx, keyDB, tx)
		err := callback(txCtx)
		if err != nil {
			return err
		}
		return nil
	}, opts...)
}
