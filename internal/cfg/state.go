package cfg

import (
	"database/sql"
	"fmt"

	"github.com/PlinyTheYounger0/stoplight/internal/database"
	"github.com/spf13/viper"
)

type State struct {
	Queries *database.Queries
	DB *sql.DB
}

func SetCurrentUser(userName string) error {

	viper.Set("current_user_name", userName)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("Failed To Login User %s: %w\n", userName, err)
	}

	fmt.Printf("User %s Logged In Successfully\n", userName)
	return nil

}
