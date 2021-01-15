# dhtcli

Tiny cli for DHT11/DHT22 using github.com/MichaelS11/go-dht

## Example

```
$ dhtcli
{"humidity":43.5,"temperature":21.8}
```

## Usage

```
Usage:
  dhtcli [OPTIONS]

Application Options:
      --gpio=       GPIO pin number (default: 26)
      --retry=      number of retry to read result (default: 5)
      --dht11       use DHT11 sensor. Use DHT22/AM2302 by default
  -F, --fahrenheit  Use Fahrenheit for temperature Unit. Celsius is used by default

Help Options:
  -h, --help        Show this help message
```
