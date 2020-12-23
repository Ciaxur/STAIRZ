package main

import (
	"fmt"
	"time"
	"bufio"
	"os"

	"github.com/stianeikeland/go-rpio"
)

// SensorConfig - Sensor Configurations
type SensorConfig struct {
	// Sensor Activation Distance (cm)
	s1_activeDist float64
	s2_activeDist float64
}

// Global Variables
// Constant Pins
const (
	// Sensor Pinout
	SENSOR1_TRIGGER = rpio.Pin(5)
	SENSOR1_ECHO    = rpio.Pin(6)
	SENSOR2_TRIGGER = rpio.Pin(23)
	SENSOR2_ECHO    = rpio.Pin(24)
	RELAY_TRIGGER   = rpio.Pin(16)
	RELAY_DELAY     = 10 * time.Second
)

// Configuration Setup
var sensorConf = SensorConfig{100, 80}
var writer *bufio.Writer

/**
 * Obtains Sensor Distance data for given
 *  sensor
 * @param trigSensor Trigger Sensor
 * @param echoSensor Echo Sensor
 * @returns Distance Data for Sensor
 */
func getDistance(trigSensor rpio.Pin, echoSensor rpio.Pin) float64 {
	// Clear TriggerPin
	fmt.Println("Waiting for Trigger to settle...")
	trigSensor.Output()
	trigSensor.Low()
	time.Sleep(10 * time.Millisecond)

	// Set Trigger High for 10us
	fmt.Println("Triggering for 10us")
	trigSensor.High()
	time.Sleep(1 * time.Microsecond)
	trigSensor.Low()

	// Get Echo Pin Data
	echoSensor.Input()
	fmt.Println("Reading Echo...")
	var pulseStart, pulseEnd time.Time
	pulseStart = time.Now()

	// Times out after 2000 and 52000 Iterations Respectively
	for i := 0; i < 2000 && SENSOR1_ECHO.Read() == 0; i++ {
		pulseStart = time.Now()
	}

	for i := 0; i < 52000 && SENSOR1_ECHO.Read() == 1; i++ {
		pulseEnd = time.Now()
	}

	// time difference between start and arrival
	elapsedTime := pulseEnd.Sub(pulseStart)
	// fmt.Printf("Elapsed Time %d\n", elapsedTime.Milliseconds())
	// multiply with the sonic speed (34300 cm/s)
	// and divide by 2, because there and back
	distance := 17150 * elapsedTime.Seconds()
	return distance
}

/**
 * Listens to Sensors given a delay
 * @param delay Time Delay
 */
func listenToSensors(delay time.Duration) {
	// Local Sensor State
	relayIsOn := false

	for true {
		// Wait...
		time.Sleep(delay)

		// Relay ON Delay
		if relayIsOn {
			// Log Activation
			writer.Write([]byte(time.Now().String() + ": Relay ON\n"))
			writer.Flush()

			time.Sleep(RELAY_DELAY)
			relayIsOn = false
			RELAY_TRIGGER.Low()
		} else {
			RELAY_TRIGGER.Low()
		}

		// Get Distance from Sensors
		dist1 := getDistance(SENSOR1_TRIGGER, SENSOR1_ECHO)
		fmt.Printf("Distance of Sensor1: %.2fcm\n", dist1)
		dist2 := getDistance(SENSOR2_TRIGGER, SENSOR2_ECHO)
		fmt.Printf("Distance of Sensor2: %.2fcm\n", dist2)

		// Relay On! (Negative means Timed out) - Sensor1 or Sensor2
		if (dist1 <= sensorConf.s1_activeDist && dist1 > 0) {
			relayIsOn = true
			RELAY_TRIGGER.High()
			writer.Write([]byte(fmt.Sprintf("Sensor1 - Triggered [%.2fcm]\n", dist1)))
		} else if (dist2 <= sensorConf.s2_activeDist && dist2 > 0) {
			relayIsOn = true
			RELAY_TRIGGER.High()
			writer.Write([]byte(fmt.Sprintf("Sensor2 - Triggered [%.2fcm]\n", dist2)))
		}

		// DEBUG: Log out Timeout
		if dist1 < 0 {
			fmt.Println("Senser 1 Timed out")
		} else if dist2 < 0 {
			fmt.Println("Senser 2 Timed out")
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Init RPIO
	if err := rpio.Open(); err != nil {
		fmt.Println("Run as Sudo!")
		fmt.Printf("Error: %v\n", err)
	}
	defer rpio.Close()

	// Init Log Output
	filepath := "./out.log"
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	defer file.Close()
	writer = bufio.NewWriter(file)
	
	// Initiate Pin Modes
	RELAY_TRIGGER.Output()

	// Listen to Sensors
	listenToSensors(50 * time.Millisecond)
	// go listenToSensors(50 * time.Millisecond)

	// TODO: Do other tasks here...
	// time.Sleep(1 * time.Hour)
}
