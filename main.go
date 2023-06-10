package main

import (
	"fmt"
	. "github.com/chaoyangnz/scs-extract/scs"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

func main() {
	dump := false
	additionalHashfsPaths := make(cli.StringSlice, 0)

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
		cli.StringSliceFlag{
			Name:  "additional-hashfs-paths, p",
			Usage: "additional hashfs paths if auto-discovery doesn't find",
			Value: &additionalHashfsPaths,
		},
	}
	app.Action = func(c *cli.Context) error {
		args := c.Args()
		if c.NArg() < 2 {
			cli.ShowAppHelp(c)
			return nil
		}

		scsFile := args.Get(0)
		destPath := args.Get(1)

		_, fileName, _ := Unjoin(scsFile)

		scs := NewSCS(scsFile)

		if _, ok := scs.GetEntryByPath(""); ok {
			fmt.Println("root found, please use official scs-extractor instead")
			return nil
		}

		if dump {
			scs.Dump(filepath.Join(destPath, fileName+"_dump"))
		}

		additionalPaths := NormalizePaths(additionalHashfsPaths)
		fmt.Printf("Additional hashfs paths: %+v\n", additionalPaths)

		scs.TryExtract(filepath.Join(destPath, fileName+"_extracted"), additionalPaths...)

		return nil
	}
	app.Run(os.Args)
}
