package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de Servidores na internet"
	flag := []cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca ips de endereços na internet",
			Flags: flag,
			Action: buscarIps,
		},

		{
			Name: "servidor",
			Usage: "Busca o nome de servidores na internet",
			Flags: flag,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}