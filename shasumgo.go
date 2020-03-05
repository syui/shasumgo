package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"github.com/urfave/cli/v2"
	//"encoding/hex"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "shasumgo"
	app.Usage = "shasumgo .bashrc .zshrc"
	return app
}

func Action(c *cli.Context) {
	app := App()
	if c.Args().Get(0) == "" {
		help := []string{"", "--help"}
		app.Run(help)
		os.Exit(1)
	}
	return
}

var (
	algorithm = flag.String("a", "1", "algorithm md5, 1, 224, 256, 384, 512, 512224, 512256")
)

func shasum(data *[]byte, algorithm string) (interface{}, error) {
	var sum interface{}
	switch strings.ToLower(algorithm) {
	case "md5":
		sum = md5.Sum(*data)
	case "1":
		sum = sha1.Sum(*data)
	case "224":
		sum = sha256.Sum224(*data)
	case "256":
		sum = sha256.Sum256(*data)
	case "384":
		sum = sha512.Sum384(*data)
	case "512":
		sum = sha512.Sum512(*data)
	case "512224":
		sum = sha512.Sum512_224(*data)
	case "512256":
		sum = sha512.Sum512_256(*data)
	default:
		return nil, errors.New("unsupported algorithm")
	}
	return sum, nil
}

func main() {
	app := &cli.App{
		Version: "0.1.1",
		Name: "shasumgo",
		Usage: "$ shasumgo go.mod go.sum",
		Action: func(c *cli.Context) error {
			if c.Args().Get(0) == "" {
				help := []string{"shasumgo", "--help"}
				fmt.Printf("%s", help)
			} else if c.NArg() > 1 {
				filea, _ := ioutil.ReadFile(c.Args().Get(0))
				suma, err := shasum(&filea, *algorithm)
				if err != nil {
					fmt.Println(err)
				}
				fileb, _ := ioutil.ReadFile(c.Args().Get(1))
				sumb, err := shasum(&fileb, *algorithm)
				if err != nil {
					fmt.Println(err)
				}
				if  suma == sumb {
					fmt.Printf("ok")
				} else {
					fmt.Printf("%x %x", suma, sumb)
				}
			} else {
				file, _ := ioutil.ReadFile(c.Args().Get(0))
				sum, err := shasum(&file, *algorithm)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("%x", sum)
				//fmt.Printf("%x\t%s\n", sum, file)
			}
			return nil
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "shasumgo c ./go.mod ./go.mod [success:result display is none]",
			Action:  func(c *cli.Context) error {
				filea, _ := ioutil.ReadFile(c.Args().Get(0))
				suma, err := shasum(&filea, *algorithm)
				if err != nil {
					fmt.Println(err)
				}
				fileb, _ := ioutil.ReadFile(c.Args().Get(1))
				sumb, err := shasum(&fileb, *algorithm)
				if err != nil {
					fmt.Println(err)
				}
				if  suma != sumb {
					fmt.Printf("%x %x", suma, sumb)
				}
				return nil
			},
		},
		{
			Name:    "s",
			Aliases: []string{"s"},
			Usage:   "shasumgo s xxxxxxxx ./go.mod",
			Subcommands: []*cli.Command{
				{
					Name:  "c",
					Aliases: []string{"c"},
					Usage: "xq s c xxxxxx ./go.mod",
					Action: func(c *cli.Context) error {
						h := c.Args().Get(0)
						filea, _ := ioutil.ReadFile(c.Args().Get(1))
						suma, err := shasum(&filea, *algorithm)
						if err != nil {
							fmt.Println(err)
						}
						s := fmt.Sprintf("%x", suma)
						if h != s {
							fmt.Printf("%s %s", h, s)
						}
						return nil
					},
				},
			},
			Action:  func(c *cli.Context) error {
				h := c.Args().Get(0)
				filea, _ := ioutil.ReadFile(c.Args().Get(1))
				suma, err := shasum(&filea, *algorithm)
				if err != nil {
					fmt.Println(err)
				}
				s := fmt.Sprintf("%x", suma)
				if h != s {
					fmt.Printf("%s %s", h, s)
				} else {
					fmt.Printf("ok")
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}

