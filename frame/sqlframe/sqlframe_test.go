// Copyright 2015 The Neugram Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlframe

import (
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"neugram.io/ng/frame/internal/frametest"
)

func createDB(t *testing.T) (db *sql.DB, cleanup func()) {
	dbfile, err := ioutil.TempFile("", "sqlframe-sqlite-")
	if err != nil {
		t.Fatal(err)
	}
	dbfile.Close()
	os.Remove(dbfile.Name())

	db, err = sql.Open("sqlite3", dbfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		db.Close()
		os.Remove(dbfile.Name())
	}
}

func TestLoadPresidents(t *testing.T) {
	db, cleanup := createDB(t)
	defer cleanup()

	txt := `
	create table Presidents (
		ID integer not null primary key,
		Name text,
		Term1 int,
		Term2 int
	);`
	if _, err := db.Exec(txt); err != nil {
		t.Fatal(err)
	}
	f, err := Load(db, "Presidents")
	if err != nil {
		t.Fatal(err)
	}
	frametest.LoadPresidents(t, f)
}

func TestLoadPresidentsNoPK(t *testing.T) {
	db, cleanup := createDB(t)
	defer cleanup()

	txt := `
	create table Presidents (
		ID integer,
		Name text,
		Term1 int,
		Term2 int
	);`
	if _, err := db.Exec(txt); err != nil {
		t.Fatal(err)
	}
	f, err := Load(db, "Presidents")
	if err != nil {
		t.Fatal(err)
	}
	frametest.LoadPresidents(t, f)
}

// TODO wrap sqlframe in dummy frame, use default impls.
// func TestLoadPresidentsNoSpec(t *testing.T)
