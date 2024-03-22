package main

import (
	"fmt"
	"os"

	"github.com/meteocima/magda_drones2wrf"
)

func main() {
	/*if len(os.Args) < 2 {
		fmt.Println("Usage: dr2w <inputfile> <outputfile>")
		os.Exit(1)
	}*/
	input := "/home/parroit/repos/cima/magda/obstools/magda_drones2wrf/fixtures/20240118225040_Lat_47.5733947_Lon_9.0468587.csv"
	output := "/home/parroit/repos/cima/magda/obstools/magda_drones2wrf/fixtures/ob.ascii"
	configFile := "/home/parroit/repos/cima/magda/obstools/magda_drones2wrf/fixtures/config.yaml"

	magda_drones2wrf.ReadConfig(configFile)
	fmt.Printf("Converting %s to %s\n", input, output)
	err := magda_drones2wrf.Convert(input, output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Conversion done")
}
