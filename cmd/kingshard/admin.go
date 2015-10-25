package main

import (
	"net/http"
	"strings"

	"github.com/flike/kingshard/config"
	"github.com/studygolang/kingshard/admin"
)

func StartAdminDashboard(cfg *config.Config) {

	configure := admin.NewConfigure(cfg)
	db := admin.NewDB(cfg)

	http.HandleFunc("/api/configure/global", configure.Global)
	http.HandleFunc("/api/configure/node_state", configure.NodeState)
	http.HandleFunc("/api/configure/schema", configure.Schema)
	http.HandleFunc("/api/db/manage", db.Manage)

	http.Handle("/", admin.FileServer("dashboard"))

	addrs := strings.Split(cfg.Addr, ":")
	addr, port := addrs[0], "1"+addrs[1]
	http.ListenAndServe(addr+":"+port, nil)
}
