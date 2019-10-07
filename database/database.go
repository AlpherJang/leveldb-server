package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDB struct {
	dbPath string
	db     *leveldb.DB
}

func NewLevelDB(dbPath string) *LevelDB {
	return &LevelDB{dbPath: dbPath}
}

func (l *LevelDB) OpenDataBase(o *opt.Options) error {
	db, err := leveldb.OpenFile(l.dbPath, nil)
	if err != nil {
		return err
	}
	l.db = db
	return nil
}

func (l *LevelDB) PutOne(key string, value []byte) error {
	return l.db.Put([]byte(key), value, nil)
}

func (l *LevelDB) Get(key string) ([]byte, error) {
	return l.db.Get([]byte(key), nil)
}
