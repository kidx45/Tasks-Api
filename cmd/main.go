// @title Task API
// @version 1.0
// @description Task Manager
// @host localhost:8080
// @BasePath /

package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	httpy "task-api/internal/adapter/inbound/http"
	"task-api/internal/adapter/outbound/persistence"
	service "task-api/internal/application"
	_ "task-api/docs"
)

func getdata(s string, r *bufio.Reader) (string, error) {
	fmt.Println(s)
	i, err := r.ReadString('\n')
	return strings.TrimSpace(i), err
}

func main() {
	user := bufio.NewReader(os.Stdin)
	s, _ := getdata("For MySQL:1, For Postgres:2", user)
	switch s {
	case "2":
		db := persistence.ConnectToPostgres()
		database := persistence.CallPsql(db)
		service := service.NewConnect(database)
		httpy.Handler(service)
	case "1":
		db := persistence.ConnectToMysql()
		database := persistence.CallMysql(db)
		service := service.NewConnect(database)
		httpy.Handler(service)
	}
}
