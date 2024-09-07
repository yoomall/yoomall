package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"lazyfury.github.com/yoomall-server/apps/app/model"
	"lazyfury.github.com/yoomall-server/apps/app/service"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/driver"
)

func main() {
	var rootCmd = &cobra.Command{Use: "seed"}
	rootCmd.AddCommand(createSuperUser())
	rootCmd.Execute()
}

func createSuperUser() *cobra.Command {
	var username string
	var password string
	reader := bufio.NewReader(os.Stdin)

	createSuperUserCmd := &cobra.Command{
		Use: "create-super-user",
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" {
				fmt.Print("username: ")
				username, _ = reader.ReadString('\n')
				username = strings.Trim(username, "\n")
			}

			if password == "" {
				fmt.Print("password: ")
				password, _ = reader.ReadString('\n')
				password = strings.Trim(password, "\n")
			}

			log.Info("create-super-user", "username", username, "password", password)
			service := service.NewAuthService(driver.NewDB(config.Config.MysqlDsn))

			err := service.CreateUser(&model.User{
				UserName: username,
				Password: password,
				Role:     0,
			})

			if err != nil {
				log.Error("create-super-user", "err", err)
			}

		},
	}

	createSuperUserCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	createSuperUserCmd.Flags().StringVarP(&password, "password", "p", "", "password")
	return createSuperUserCmd
}
