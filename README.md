# scs-extract

`scs-extract` is a CLI tool to list / extract files from SCS archives used in Euro Truck Simulator 2 / American Truck Simulator.

For some SCS, the root may be missing or intentionally prevent others to extract. When you use the official SCS extractor, you may encounter
```
[hashfs] xxx.scs: Mounted ok, 12345 entries
*** ERROR *** : Root directory not found, can not extract this archive
```
This utility will try the best to extract the rootless hashfs files.

## Usage

`scs-extract [options] <scs_file> [dest_folder]`

```console
# scs-extract ets2/accessory_pack.scs .
```

```
# scs-extract --help

NAME:
scs-extract - Extract SCS file to a folder

USAGE:
scs-extract [-options] scs_file dest_path [args...]

VERSION:
1.0.0

DESCRIPTION:
Try the best to extract SCS file, especially useful for some odd SCS missing roots

COMMANDS:
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
--dump, -d     dump raw hashfs files
--help, -h     show help
--version, -v  print the version
```
