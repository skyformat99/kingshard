package admin

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"

	_ "github.com/go-sql-driver/mysql"
)

type fileHandler struct {
	root string
}

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at root.
func FileServer(root string) http.Handler {
	return &fileHandler{root}
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath

	}
	upath = path.Clean(upath)

	// See https://golang.org/pkg/net/http/#StripPrefix
	_, r.URL.Path = filepath.Split(upath)
	r.URL.Path = "/" + r.URL.Path

	name := f.root + upath

	// consist js: history.pushState()
	_, err := os.Stat(f.root + upath)
	if err != nil {
		r.URL.Path = "/"
		name = f.root + "/"
	}

	http.ServeFile(w, r, name)
}

type base struct {
	*sql.DB
	*config.Config
}

func (b *base) openDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", b.User, b.Password, b.Addr, b.Schemas[0].DB)
	b.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		golog.Error("admin", "openDB", err.Error(), 0)
	} else {
		b.SetMaxOpenConns(8)
		b.SetMaxIdleConns(4)
	}
}
