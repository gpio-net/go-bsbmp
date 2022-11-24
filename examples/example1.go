package main

import (
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	lg "github.com/d2r2/go-logger"
)

func main() {
	// Create new connection to i2c-bus on 1 line with address 0x76.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.NewI2C(0x76, 5, "localhost:3001")
	if err != nil {
		lg.Fatal(err)
	}
	defer i2c.Close()

	lg.Notify("***************************************************************************************************")

	// sensor, err := bsbmp.NewBMP(bsbmp.BMP180, i2c) // signature=0x55
	// sensor, err := bsbmp.NewBMP(bsbmp.BMP280, i2c) // signature=0x58
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60
	// sensor, err := bsbmp.NewBMP(bsbmp.BMP388, i2c) // signature=0x50
	if err != nil {
		lg.Fatal(err)
	}

	id, err := sensor.ReadSensorID()
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("This Bosch Sensortec sensor has signature: 0x%x", id)

	err = sensor.IsValidCoefficients()
	if err != nil {
		lg.Fatal(err)
	}

	// Read temperature in celsius degree
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("Temprature = %v*C", t)

	// Read atmospheric pressure in pascal
	p, err := sensor.ReadPressurePa(bsbmp.ACCURACY_LOW)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("Pressure = %v Pa", p)

	// Read atmospheric pressure in mmHg
	p, err = sensor.ReadPressureMmHg(bsbmp.ACCURACY_LOW)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("Pressure = %v mmHg", p)

	// Read atmospheric pressure in mmHg
	supported, h1, err := sensor.ReadHumidityRH(bsbmp.ACCURACY_LOW)
	if supported {
		if err != nil {
			lg.Fatal(err)
		}
		lg.Infof("Humidity = %v %%", h1)
	}

	// Read atmospheric altitude in meters above sea level, if we assume
	// that pressure at see level is equal to 101325 Pa.
	a, err := sensor.ReadAltitude(bsbmp.ACCURACY_LOW)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("Altitude = %v m", a)
}
