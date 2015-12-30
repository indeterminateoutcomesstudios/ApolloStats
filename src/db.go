package apollostats

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const MAX_ROWS = 200

// NOTE: DON'T USE ANY WRITE OPERATIONS ON THE DATABASE!
// We're interfacing with an external, live game database!

type DB struct {
	*gorm.DB
}

func OpenDB(DSN string) (*DB, error) {
	tmp := fmt.Sprintf("%s?parseTime=True&timeout=30s", DSN)
	db, e := gorm.Open("mysql", tmp)
	if e != nil {
		return nil, e
	}
	return &DB{&db}, nil
}

func (db *DB) AllBans() []*Ban {
	var tmp []*Ban
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) AllAccountItems() []*AccountItem {
	var tmp []*AccountItem
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) AllDeaths() []*Death {
	var tmp []*Death
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) AllRounds() []*RoundStats {
	var tmp []*RoundStats
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) GetRound(id int64) *RoundStats {
	var tmp RoundStats
	db.First(&tmp, id)
	return &tmp
}

func (db *DB) GetLatestRound() *RoundStats {
	var tmp RoundStats
	db.Order("id desc").Limit(1).First(&tmp)
	return &tmp
}

func (db *DB) GetAntags(id int64) []*RoundAntags {
	var tmp []*RoundAntags
	db.Order("round_id desc, name asc").Where("round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetAILaws(id int64) []*RoundAILaws {
	var tmp []*RoundAILaws
	db.Order("round_id desc, law asc").Where("round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetDeaths(id int64) []*Death {
	var tmp []*Death
	db.Order("round_id desc, name asc").Where("round_id = ?", id).Find(&tmp)
	return tmp
}
