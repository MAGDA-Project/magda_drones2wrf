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
        config file to use (default "/home/parroit/.magda_drones2wrf.yaml")

```

* <input file> is the path of the dataset source in csv format. You can see an example of the format [here](magda_drones2wrf/fixtures/source.csv)
* <output file> is the path of WRF ascii file to create.