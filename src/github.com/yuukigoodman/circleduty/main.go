package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func main() {
	viper.SetEnvPrefix("circleduty")
	viper.BindEnv("auth_token")
	viper.AutomaticEnv()

	app := cli.NewApp()
	app.Name = "circleduty"
	app.Usage = "circleduty fail"
	app.Action = func(c *cli.Context) error {
		message := fmt.Sprintf("PagerDuty Incident created: %s", viper.Get("auth_token"))
		fmt.Println(message)
		return nil
	}

	app.Run(os.Args)
}
