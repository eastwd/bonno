package controllers

// modelsはあとで作成します
import (
	"bonno/app/models"
	"database/sql"
	"github.com/go-gorp/gorp/v3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

// DbMapという変数にgorpのポインタを持たせる
var (
	DbMap *gorp.DbMap
)

//テーブルの初期化処理

func InitDB() {
	dbdir := revel.Config.StringDefault("db.directory", "./app.db")
	dbdriver := revel.Config.StringDefault("db.driver", "sqlite3")
	db, err := sql.Open(dbdriver, dbdir)
	if err != nil {
		panic(err.Error())
	}
	DbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	t := DbMap.AddTable(models.Movie{}).SetKeys(true, "Id")
	t.ColMap("Name").MaxSize = 50

	DbMap.CreateTables()
}

// GorpContrllerの定義
type GorpController struct {
	*revel.Controller
	Transaction *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := DbMap.Begin()
	if err != nil {
		panic(err)
	}
	c.Transaction = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Transaction == nil {
		return nil
	}
	err := c.Transaction.Commit() // SQLによる変更をDBに反映
	if err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Transaction = nil // 正常終了した場合はROLLBACK処理に入らないようにする
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Transaction == nil {
		return nil
	}
	err := c.Transaction.Rollback() // 問題があった場合変更前の状態に戻す
	if err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Transaction = nil
	return nil
}
