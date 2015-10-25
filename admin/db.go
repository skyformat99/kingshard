// The struct of db manages the mysql database
package admin

import (
	"fmt"
	"net/http"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
)

type db struct {
	*base
}

// NewDBManage create the db instance with cfg
func NewDB(cfg *config.Config) *db {
	d := &db{
		base: &base{
			Config: cfg,
		},
	}

	d.openDB()

	return d
}

// Manage do the database manage
func (d *db) Manage(w http.ResponseWriter, r *http.Request) {
	opt := r.PostFormValue("opt")
	node := r.PostFormValue("node")
	k := r.PostFormValue("k")
	v := r.PostFormValue("v")
	if opt == "" || node == "" || k == "" || v == "" {
		fmt.Fprint(w, `{"code":1,"msg":"opt/node/k/v is empty!"}`)
		return
	}

	strSql := fmt.Sprintf("admin node(opt,node,k,v) values('%s','%s','%s','%s')", opt, node, k, v)
	fmt.Println(strSql)
	result, err := d.Exec(strSql)
	if err != nil {
		golog.Error("admin", "Manage", err.Error(), 0)
		fmt.Fprint(w, `{"code":2,"msg":"proxy error(exec)"}`)
		return
	}

	_, err = result.RowsAffected()
	if err != nil {
		golog.Error("admin", "Manage", err.Error(), 0)
		fmt.Fprint(w, `{"code":3,"msg":"proxy error(rows affected)"}`)
		return
	}

	fmt.Fprint(w, `{"code":0,"msg":"ok"}`)
	return
}
