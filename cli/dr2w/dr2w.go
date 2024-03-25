package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/meteocima/magda_drones2wrf"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dr2w <inputfile> <outputfile>")
		os.Exit(1)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var configFile string
	flag.StringVar(&configFile, "c", "", "config file to use (default: ~/.magda_drones2wrf.yaml)")
	flag.Parse()
	if configFile == "" {
		configFile = path.Join(home, ".magda_drones2wrf.yaml") //"/home/parroit/repos/cima/magda/obstools/magda_drones2wrf/fixtures/config.yaml"
	}
	fmt.Println(flag.Args())
	input := flag.Arg(0)  //"/home/parroit/repos/cima/magda/obstools/magda_drones2wrf/fixtures/20240118225040_Lat_47.5733947_Lon_9.0468587.csv"
	output := flag.Arg(1) //"/home/parroit/repos/cima/magda/obstools/magda_drones2wrf/fixtures/ob.ascii"

	magda_drones2wrf.ReadConfig(configFile)
	fmt.Printf("Converting %s to %s\n", input, output)
	err = magda_drones2wrf.Convert(input, output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Conversion done")
}
