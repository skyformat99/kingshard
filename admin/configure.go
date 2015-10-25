package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
)

type configure struct {
	*base
}

// NewConfigure create the configure instance with cfg
func NewConfigure(cfg *config.Config) *configure {
	c := &configure{
		base: &base{
			Config: cfg,
		},
	}

	c.openDB()

	return c
}

type globalConfig struct {
	Key   string
	Value string
}

// Global output global configures
func (c *configure) Global(w http.ResponseWriter, r *http.Request) {
	rows, err := c.Query("admin server(opt,k,v) values('show','proxy','config')")
	if err != nil {
		golog.Error("admin", "Global", err.Error(), 0)
		fmt.Fprint(w, `{"code":1,"msg":"proxy error"}`)
		return
	}
	defer rows.Close()

	globalConfigs := make([]globalConfig, 0, 10)
	var middleConfig globalConfig

	for rows.Next() {
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			golog.Error("admin", "Global", err.Error(), 0)
			continue
		}
		middleConfig.Key = key
		middleConfig.Value = value

		globalConfigs = append(globalConfigs, middleConfig)
	}

	dataBytes, err := json.Marshal(globalConfigs)
	if err != nil {
		golog.Error("admin", "Global", err.Error(), 0)
		fmt.Fprint(w, `{"code":2,"msg":"proxy error"}`)
		return
	}

	fmt.Fprint(w, `{"code":0,"msg":"ok","data":`+string(dataBytes)+`}`)
}

type nodeState struct {
	Node        string
	Address     string
	Type        string
	State       string
	LastPing    string
	MaxIdleConn string
	IdleConn    string
}

// NodeState output node's state
func (c *configure) NodeState(w http.ResponseWriter, r *http.Request) {
	rows, err := c.Query("admin server(opt,k,v) values('show','node','config')")
	if err != nil {
		golog.Error("admin", "NodeState", err.Error(), 0)
		fmt.Fprint(w, `{"code":1,"msg":"proxy error"}`)
		return
	}
	defer rows.Close()

	nodeStates := make([]nodeState, 0, 10)
	var middleConfig nodeState

	for rows.Next() {
		var node, addr, typ, state, lastPing, maxIdleConn, idleConn string
		err = rows.Scan(&node, &addr, &typ, &state, &lastPing, &maxIdleConn, &idleConn)
		if err != nil {
			golog.Error("admin", "NodeState", err.Error(), 0)
			continue
		}
		middleConfig.Node = node
		middleConfig.Address = addr
		middleConfig.Type = typ
		middleConfig.State = state
		middleConfig.LastPing = lastPing
		middleConfig.MaxIdleConn = maxIdleConn
		middleConfig.IdleConn = idleConn

		nodeStates = append(nodeStates, middleConfig)
	}

	dataBytes, err := json.Marshal(nodeStates)
	if err != nil {
		golog.Error("admin", "NodeState", err.Error(), 0)
		fmt.Fprint(w, `{"code":2,"msg":"proxy error"}`)
		return
	}

	fmt.Fprint(w, `{"code":0,"msg":"ok","data":`+string(dataBytes)+`}`)
}

type schema struct {
	DB            string
	Table         string
	Type          string
	Key           string
	Nodes_List    string
	Locations     string
	TableRowLimit string
}

// Schema output the schema
func (c *configure) Schema(w http.ResponseWriter, r *http.Request) {
	rows, err := c.Query("admin server(opt,k,v) values('show','schema','config')")
	if err != nil {
		golog.Error("admin", "Schema", err.Error(), 0)
		fmt.Fprint(w, `{"code":1,"msg":"proxy error"}`)
		return
	}
	defer rows.Close()

	schemas := make([]schema, 0, 10)
	var middleConfig schema

	for rows.Next() {
		var db, table, typ, key, nodesList, locations, tableRowLimit string
		err = rows.Scan(&db, &table, &typ, &key, &nodesList, &locations, &tableRowLimit)
		if err != nil {
			golog.Error("admin", "Schema", err.Error(), 0)
			continue
		}
		middleConfig.DB = db
		middleConfig.Table = table
		middleConfig.Type = typ
		middleConfig.Key = key
		middleConfig.Nodes_List = nodesList
		middleConfig.Locations = locations
		middleConfig.TableRowLimit = tableRowLimit

		schemas = append(schemas, middleConfig)
	}

	dataBytes, err := json.Marshal(schemas)
	if err != nil {
		golog.Error("admin", "Schema", err.Error(), 0)
		fmt.Fprint(w, `{"code":2,"msg":"proxy error"}`)
		return
	}

	fmt.Fprint(w, `{"code":0,"msg":"ok","data":`+string(dataBytes)+`}`)
}
