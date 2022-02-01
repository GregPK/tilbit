package core

import (
	"database/sql"
	"errors"
	"log"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	db, err := sql.Open("sqlite3", "file:foobar?mode=memory&cache=shared")
	// db, err := sql.Open("sqlite3", "file:foobar.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	repo := SQLiteRepository{db: db}
	repo.Setup()

	return &repo
}

func (r *SQLiteRepository) Setup() error {
	query := `
	CREATE TABLE IF NOT EXISTS tilbits(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL,
			hash TEXT NOT NULL,
			sourceAuthor TEXT,
			sourceName TEXT,
			sourceUrl TEXT,
			sourceAddedOn TEXT,
			sourceLocationUrl TEXT,
			sourceLocationLineNumber TEXT
	);
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM tilbits;")
	return err
}

func (r *SQLiteRepository) Import(tilbits []Tilbit) error {
	for _, tilbit := range tilbits {
		_, err := r.Create(tilbit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SQLiteRepository) Create(tilbit Tilbit) (*Tilbit, error) {
	res, err := r.db.Exec("INSERT INTO tilbits(text, hash, sourceAuthor, sourceName, sourceUrl, sourceAddedOn, sourceLocationUrl, sourceLocationLineNumber) values(?,?,?,?,?,?,?,?)",
		tilbit.Text, tilbit.Hash(), tilbit.Data.Author, tilbit.Data.Source, tilbit.Data.Url, tilbit.Data.AddedOn, tilbit.Location.Uri, tilbit.Location.LineNumber)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	tilbit.DbID = id

	return &tilbit, nil
}

func (r *SQLiteRepository) All() (all []Tilbit, err error) {
	return r.scanRows("1=1")
}

func (r *SQLiteRepository) ByIds(hashes []string) (foundBits []Tilbit, err error) {
	for _, hash := range hashes {
		bits, errsc := r.scanRows("hash LIKE '%" + hash + "%'")
		if errsc != nil {
			return nil, errsc
		}
		foundBits = append(foundBits, bits...)
	}

	return foundBits, nil
}

func (r *SQLiteRepository) scanRows(where string) (all []Tilbit, err error) {
	rows, err := r.db.Query("SELECT * FROM tilbits WHERE " + where)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tilbit Tilbit
		var source SourceMetadata
		var sourceLoc SourceLocation

		err := rows.Scan(&tilbit.DbID, &tilbit.Text, &tilbit.DbHash,
			&source.Author, &source.Source, &source.Url, &source.AddedOn,
			&sourceLoc.Uri, &sourceLoc.LineNumber)
		if err != nil {
			return nil, err
		}

		tilbit.Data = source
		tilbit.Location = sourceLoc

		all = append(all, tilbit)
	}
	return
}

func (r *SQLiteRepository) ById(hash string) (tilbit Tilbit, err error) {
	tilbits, err := r.ByIds([]string{hash})
	if len(tilbits) > 0 {
		tilbit = tilbits[0]
	}

	return
}

func (r *SQLiteRepository) ByQuery(query string) (tilbits []Tilbit, err error) {

	if query == "all" {
		tilbits, _ = r.All()
	} else if query == "random" {
		allBits, _ := r.All()
		randTil := getRandomBit(allBits)
		tilbits = append(tilbits, randTil)
	} else {
		ids := ParseIdsFromString(query)

		tilbits, err = r.ByIds(ids)
	}
	return
}

func (r *SQLiteRepository) SetSourceURI(source string) error {
	// no-op, will be used later for passing in file
	return nil
}
