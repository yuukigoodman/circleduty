package main

import (
	"fmt"
	"os"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func main() {
	viper.SetEnvPrefix("circleduty")
	viper.BindEnv("auth_token")
	viper.BindEnv("service_key")
	viper.AutomaticEnv()

	app := cli.NewApp()
	app.Name = "circleduty"
	app.Usage = "circleduty fail"
	app.Action = func(c *cli.Context) error {
		e := pagerduty.Event{
			Type:        "trigger",
			ServiceKey:  viper.GetString("service_key"),
			Description: "test desctiption",
		}

		if res, err := pagerduty.CreateEvent(e); err != nil {
			panic(err)
		} else {
			message := fmt.Sprintf("PagerDuty Incident created: %s", res.IncidentKey)
			fmt.Println(message)
		}

		return nil
	}

	app.Run(os.Args)
}
