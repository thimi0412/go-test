package main

import (
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli" // cli.goをインポート
)

func main() {
	app := cli.NewApp()
	app.Name = "cal"              // ヘルプを表示する際に使用される
	app.Usage = "print arguments" // ヘルプを表示する際に使用される
	app.Version = "1.0.0"         // ヘルプを表示する際に使用される
	app.Commands = []cli.Command{
		{
			Name:  "plus",
			Usage: "number1 plus number2",
			Action: func(c *cli.Context) error {
				var num1, _ = strconv.Atoi(c.Args()[0])
				var num2, _ = strconv.Atoi(c.Args()[1])
				println(num1 + num2)
				return nil
			},
		},
		{
			Name:  "minus",
			Usage: "number1 subtracted number2",
			Action: func(c *cli.Context) error {
				var num1, _ = strconv.Atoi(c.Args()[0])
				var num2, _ = strconv.Atoi(c.Args()[1])
				println(num1 - num2)
				return nil
			},
		},
		{
			Name:  "times",
			Usage: "number1 multiplied number2",
			Action: func(c *cli.Context) error {
				var num1, _ = strconv.Atoi(c.Args()[0])
				var num2, _ = strconv.Atoi(c.Args()[1])
				println(num1 * num2)
				return nil
			},
		},
		{
			Name:  "div",
			Usage: "number1 divided number2",
			Action: func(c *cli.Context) error {
				var num1, _ = strconv.Atoi(c.Args()[0])
				var num2, _ = strconv.Atoi(c.Args()[1])
				print(num1 / num2)
				print(" mod:")
				println(num1 % num2)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
