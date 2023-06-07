package main

import (
	. "github.com/chaoyangnz/scs-extract/scs"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

func main() {
	dump := false

	app := cli.NewApp()
	app.Name = "scs-extract"
	app.Usage = "Extract SCS file to a folder"
	app.Description = ""
	app.UsageText = "scs-extract [-options] scs_file dest_path [args...]"
	app.Version = "1.0.0"
	app.Author = "Chao Yang"
	app.Email = "chao@yang.so"
	app.Description = "Try the best to extract SCS file, especially useful for some odd SCS missing roots"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "dump, d",
			Usage:       "dump raw hashfs files",
			Destination: &dump,
		},
	}
	app.Action = func(c *cli.Context) error {
		args := c.Args()
		if c.NArg() < 1 {
			return nil
		}

		scsFile := args.Get(0)
		destPath := args.Get(1)

		scs := NewSCS(scsFile)
		if dump {
			scs.Dump(filepath.Join(destPath, "dump"))
		}
		scs.TryExtract(filepath.Join(destPath, "extracted"))

		return nil
	}
	app.Run(os.Args)
}
