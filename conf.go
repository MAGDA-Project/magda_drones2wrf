package magda_drones2wrf

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigValues = struct {
	Elevation          float64 `yaml:"Elevation"`
	PressureError      float64 `yaml:"PressureError"`
	WindSpeedError     float64 `yaml:"WindSpeedError"`
	WindDirectionError float64 `yaml:"WindDirectionError"`
	AltitudeError      float64 `yaml:"AltitudeError"`
	TemperatureError   float64 `yaml:"TemperatureError"`
	DewpointError      float64 `yaml:"DewpointError"`
	HumidityError      float64 `yaml:"HumidityError"`
}{}

func ReadConfig(cfgFile string) {
	fmt.Printf("Reading configuration from %s\n", cfgFile)
	cfg, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Panicf("Error reading configuration file: %s", err)
	}
	err = yaml.Unmarshal(cfg, &ConfigValues)
	if err != nil {
		log.Panicf("Error parsing configuration file: %s", err)
	}

}
