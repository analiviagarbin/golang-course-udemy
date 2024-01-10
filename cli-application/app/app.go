package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Create returns the CLI application ready to be executed
func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and Server Names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "linkedin.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search for Server Names",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	// net
	ips, errors := net.LookupIP(host)
	if errors != nil {
		log.Fatal(errors)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	server, errors := net.LookupNS(host)
	if errors != nil {
		log.Fatal(errors)
	}

	for _, server := range server {
		fmt.Println(server)
	}

}
