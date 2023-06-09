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
--dump, -d                      dump raw hashfs files
--aditional-hashfs-paths, -p    additional hashfs paths
--help, -h                      show help
--version, -v                   print the version
```

## Other workaround

This utility should be able to extract most of the files in the headless SCS. But if anything missing, 
you can use [ConverterPIX](https://github.com/mwl4/ConverterPIX) to convert the model (.pmg/.pmd), and
it should give you some warnings to indicate which files are missing. Then you can add the file paths to `-p`
which can help the utility to extract the missing files.

Here is an example

```
converter_pix.exe -m /vehicle/truck/xxx/yyy -b d:/ets2/mode/base

 ******************************************
 **        Converter PMX to PIX          **
 **       Copyright (C) 2017 mwl4        **
 ******************************************

<warning> [material] /automat/31/31d0e65efaab21fa.mat: Unable to open material!
<warning> [material] /automat/31/31d0e65efaab21fa.mat: Unable to open material!
<warning> [material] /automat/31/31d0e65efaab21fa.mat: Unable to open material!
[model] gps_tmp: pim:yes pit:yes pis:no pic:no pip:no vertices:93 indices:84 materials:3
```

`/vehicle/truck/upgrade/xx_yyy` is the hashfs path pointing to `/vehicle/truck/upgrade/xx_yyy.pmg` and/or `/vehicle/truck/upgrade/xx_yyy.pmd`

From the warning, I know `/automat/31/31d0e65efaab21fa.mat` is missing. So I add it to the additional paths

```
scs-extract.exe -d -p /automat/31/31d0e65efaab21fa.mat odd.scs e:/dest
```

Once it is done, you can run `converter_pix.exe` again, it should succeed