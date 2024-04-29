package controllers

import (
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/irtse/sqldb"
)

// Operations about object
type QueryController struct {
	beego.Controller
}

type Query struct {
	Text string `json:"query,omitempty"`
}

// @Title Create
// @Description create object
// @Param	body		body 	Query	true		"The query"
// @Success 200 {string} result
// @Failure 403 body is empty
// @router / [post]
func (qc *QueryController) Post() {
	var query Query
	//query.Query = "select * from dbuser;"
	json.Unmarshal(qc.Ctx.Input.RequestBody, &query)

	conn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		fmt.Println(err.Error())
	}
	db := sqldb.Open("sqlserver", conn)

	defer db.Close()

	data, err := db.QueryAssociativeArray(query.Text)
	if err != nil {
		fmt.Println(err.Error())
	}

	qc.Data["json"] = data
	qc.ServeJSON()

}
