package users

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
)

func GetPostsCount(id string) (int, error) {
	var result	int
	row := mysql.DB.QueryRow("SELECT COUNT(id) FROM posts WHERE user_id = ?;", id)
	err := row.Scan(&result)
	if err != nil {
		fmt.Print(err.Error())
	}
	return result, err
}
