# MAGDA Drones observations to WRF

This module can be used to convert drones observations in MAGDA format into ascii
WRF format.

## Usage on CIMA Typhoon

`dr2w` is already present in /data/safe/home/wrfprod/bin/dr2w

## Command line usage

This module implements a console command
that can be used to convert observations from
CSV to ascii WRF format.

Usage of `dr2w`:

```
Usage: dr2w [options] <inputfile> <outputfile>
Options:
  -c string
        config file to use (default "~/.magda_drones2wrf.yaml")

```

* <input file> is the path of the dataset source in csv format. You can see an example of the format [here](fixtures/source2.csv)
* <output file> is the path of WRF ascii file to create.
* An example of the configuration file is available [here](fixtures/config.yaml) 


## Build

1) If not already installed, follow the instructions here to install the Go Workbench::
[Install-go-environment](https://github.com/meteocima/documentazione/wiki/Installare-ambiente-di-lavoro-go)

2) Clone this repository 

```bash
git clone https://github.com/MAGDA-Project/magda_drones2wrf.git
cd magda_drones2wrf
```

5) Compile the executable with this command:

```bash
go build ./cli/dr2w
```