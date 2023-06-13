package main

import (
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

		//oldImpl(scsFile, destPath, dump, additionalHashfsPaths)

		newImpl(scsFile, destPath, dump, additionalHashfsPaths)

		return nil
	}
	app.Run(os.Args)
}

func newImpl(scsFile string, destPath string, dump bool, additionalHashfsPaths []string) {
	_, fileName, _ := Unjoin(scsFile)
	additionalPaths := NormalizePaths(additionalHashfsPaths...)

	scs := NewHashfs(scsFile)
	scs.Extract(filepath.Join(destPath, fileName+"_extracted"), additionalPaths...)

	if dump {
		scs.Dump(filepath.Join(destPath, fileName+"_dump"))
	}

	//if entry, ok := scs.NodeByPath("vehicle/truck/tmp_acc/gps_tmp_uk.pmd"); ok {
	//	fmt.Printf("%+v \n", entry)
	//}
}
