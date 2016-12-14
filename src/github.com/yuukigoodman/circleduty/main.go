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
	viper.AutomaticEnv()

	authtoken := viper.GetString("auth_token")
	var opts pagerduty.ListEscalationPoliciesOptions
	pd := pagerduty.NewClient(authtoken)

	app := cli.NewApp()
	app.Name = "circleduty"
	app.Usage = "circleduty fail"
	app.Action = func(c *cli.Context) error {
		if eps, err := pd.ListEscalationPolicies(opts); err != nil {
			panic(err)
		} else {
			for _, p := range eps.EscalationPolicies {
				fmt.Println(p.Name)
			}
		}

		message := fmt.Sprintf("PagerDuty Incident created: %s", authtoken)
		fmt.Println(message)
		return nil
	}

	app.Run(os.Args)
}
