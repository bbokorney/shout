package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

const (
	defaultTemplateDir     = "templates"
	defaultUsersConfigFile = "users.yml"
)

func main() {
	app := cli.NewApp()

	app.Name = "shout"
	app.Usage = "Be heard"
	app.UsageText = "shout [options]"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Baker Bokorney",
			Email: "bbokorney@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "users, u",
			Value:  "users.yml",
			Usage:  "User mapping `FILE`",
			EnvVar: "SHOUT_USERS_FILE",
		},
		cli.StringFlag{
			Name:   "templates, t",
			Value:  "templates",
			Usage:  "Template directory `DIR`",
			EnvVar: "SHOUT_TEMPLATES_DIR",
		},
		cli.StringFlag{
			Name:   "slack-url, s",
			Usage:  "Slack incoming webhook `URL`",
			EnvVar: "SHOUT_SLACK_URL",
		},
		cli.IntFlag{
			Name:   "port, p",
			Usage:  "`PORT`",
			Value:  8080,
			EnvVar: "SHOUT_PORT",
		},
	}

	app.Action = func(c *cli.Context) error {

		parsedTemplates, err := ParseTemplates(c.String("templates"))
		if err != nil {
			log.Fatal(err)
		}

		usersConfig, err := ReadUsersFile(c.String("users"))
		if err != nil {
			log.Fatal(err)
		}

		slackURL := c.String("slack-url")
		if slackURL == "" {
			log.Fatal("Slack url must be defined")
		}

		users := NewUsers(usersConfig)
		templates := NewTemplates(parsedTemplates)
		notifications := NewNotifications(slackURL)
		shouter := NewShouter(users, templates, notifications)
		shoutHandler := NewShoutHandler(shouter)

		http.HandleFunc("/shout", shoutHandler)

		log.Printf("Listening on port %d", c.Int("port"))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), nil))
		return nil
	}

	app.Run(os.Args)
}
