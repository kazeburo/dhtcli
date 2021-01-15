package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MichaelS11/go-dht"
	"github.com/jessevdk/go-flags"
)

type commandOpts struct {
	GPIO       int  `long:"gpio" default:"26" description:"GPIO pin number"`
	Retry      int  `long:"retry" default:"5" description:"number of retry to read result"`
	DHT11      bool `long:"dht11" description:"use DHT11 sensor. Use DHT22/AM2302 by default"`
	Fahrenheit bool `short:"F" long:"fahrenheit" description:"Use Fahrenheit for temperature Unit. Celsius is used by default"`
}

var opts commandOpts

func main() {
	opts = commandOpts{}
	psr := flags.NewParser(&opts, flags.Default)
	_, err := psr.Parse()
	if err != nil {
		os.Exit(1)
	}

	err = dht.HostInit()
	if err != nil {
		log.Fatalf("HostInit error: %v", err)
	}
	pin := fmt.Sprintf("GPIO%d", opts.GPIO)
	sensorType := ""
	if opts.DHT11 {
		sensorType = "dht11"
	}
	temperatureUnit := dht.Celsius
	if opts.Fahrenheit {
		temperatureUnit = dht.Fahrenheit
	}
	dht, err := dht.NewDHT(pin, temperatureUnit, sensorType)
	if err != nil {
		log.Fatalf("NewDHT error: %v", err)
	}

	humidity, temperature, err := dht.ReadRetry(opts.Retry)
	if err != nil {
		log.Fatalf("Read error: %v", err)
	}

	json.NewEncoder(os.Stdout).Encode(map[string]interface{}{"humidity": humidity, "temperature": temperature})
	os.Exit(0)
}
