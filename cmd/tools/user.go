package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/yoomall/yoomall/apps/auth/model"
	authservice "github.com/yoomall/yoomall/apps/auth/service"

	"github.com/charmbracelet/log"
	"github.com/lazyfury/pulse/helper/utils"
	"github.com/spf13/cobra"
)

func seedingUsers() *cobra.Command {
	var count int
	cmd := &cobra.Command{
		Use: "seeding:users",
		Run: func(cmd *cobra.Command, args []string) {
			startTime := time.Now()
			service := authservice.NewAuthService(getDB())

			for i := 0; i < count; i++ {
				email := getRandomEmail()
				if err := service.CreateUser(&model.User{
					UserName: utils.StringUtils.HiddenEmail(email),
					Password: "yoo123456",
					Email:    email,
					Phone:    getRandomPhone(),
				}); err != nil {
					log.Error(err.Error())
				}
			}

			endTime := time.Now()
			log.Info(fmt.Sprintf("seeding users success, cost %f 分钟", endTime.Sub(startTime).Minutes()))
		},
	}

	cmd.Flags().IntVarP(&count, "count", "c", 10, "number of users")
	return cmd
}

func getRandomEmail() string {
	return fmt.Sprintf("%s@%s.com", getRandomStr(8), "yoomall")
}

func getRandomPhone() string {
	return fmt.Sprintf("%s%s", "133", getRandomNumber(8))
}

func getRandomNumber(length int) string {
	seed := "1234567890"
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = seed[rand.IntN(len(seed))]
	}
	return string(randStr)
}

func getRandomStr(length int) string {
	seed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = seed[rand.IntN(len(seed))]
	}
	return string(randStr)
}

func createSuperUser() *cobra.Command {
	var username string
	var password string
	var email string
	var phone string
	reader := bufio.NewReader(os.Stdin)

	createSuperUserCmd := &cobra.Command{
		Use: "create:superuser",
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" {
				fmt.Print("username: ")
				username, _ = reader.ReadString('\n')
				username = trimShellInputString(username)
			}

			if password == "" {
				fmt.Print("password: ")
				password, _ = reader.ReadString('\n')
				password = trimShellInputString(password)
			}

			if email == "" {
				fmt.Print("email: ")
				email, _ = reader.ReadString('\n')
				email = trimShellInputString(email)
			}

			if phone == "" {
				fmt.Print("phone: ")
				phone, _ = reader.ReadString('\n')
				phone = trimShellInputString(phone)
			}

			log.Info("create-super-user", "username", username, "password", password)
			service := authservice.NewAuthService(getDB())

			if err := service.CheckPasswordStrength(password); err != nil {
				log.Error("create-super-user", "err", err)
			}

			err := service.CreateUser(&model.User{
				UserName: username,
				Password: password,
				Role:     0,
				Email:    email,
				Phone:    phone,
			})

			if err != nil {
				log.Error("create-super-user", "err", err)
			}

		},
	}

	createSuperUserCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	createSuperUserCmd.Flags().StringVarP(&password, "password", "p", "", "password")
	createSuperUserCmd.Flags().StringVarP(&email, "email", "e", "", "email")
	createSuperUserCmd.Flags().StringVarP(&phone, "phone", "t", "", "phone")
	return createSuperUserCmd
}
