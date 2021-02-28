package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/stianeikeland/go-rpio"
	"github.com/joho/godotenv"
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
var db *sql.DB
var sensorConf = SensorConfig{100, 70}

// Obtains Sensor Distance data for given sensor
// @param trigSensor Trigger Sensor
// @param echoSensor Echo Sensor
// @returns Distance Data for Sensor
func getDistance(trigSensor rpio.Pin, echoSensor rpio.Pin) float64 {
	// Clear TriggerPin
	trigSensor.Output()
	trigSensor.Low()
	time.Sleep(10 * time.Millisecond)

	// Set Trigger High for 10us
	trigSensor.High()
	time.Sleep(1 * time.Microsecond)
	trigSensor.Low()

	// Get Echo Pin Data
	echoSensor.Input()
	var pulseStart, pulseEnd time.Time
	pulseStart = time.Now()

	// Times out after 2000 and 52000 Iterations Respectively
	for i := 0; i < 2000 && echoSensor.Read() == 0; i++ {
		pulseStart = time.Now()
	}

	for i := 0; i < 52000 && echoSensor.Read() == 1; i++ {
		pulseEnd = time.Now()
	}

	// time difference between start and arrival
	elapsedTime := pulseEnd.Sub(pulseStart)
	// multiply with the sonic speed (34300 cm/s)
	// and divide by 2, because there and back
	distance := 17150 * elapsedTime.Seconds()
	return distance
}

// Listens to Sensors given a delay
// @param delay Time Delay
func listenToSensors(delay time.Duration) {
	// Local Sensor State
	relayIsOn := false

	for true {
		// Wait...
		time.Sleep(delay)

		// Relay ON Delay
		if relayIsOn {
			time.Sleep(RELAY_DELAY)
			relayIsOn = false
			RELAY_TRIGGER.Low()
		} else {
			RELAY_TRIGGER.Low()
		}

		// Get Distance from Sensors (Average two Samples)
		dist1 := getDistance(SENSOR1_TRIGGER, SENSOR1_ECHO)
		time.Sleep(25 * time.Millisecond)
		dist1 += getDistance(SENSOR1_TRIGGER, SENSOR1_ECHO)
		dist1 /= 2

		dist2 := getDistance(SENSOR2_TRIGGER, SENSOR2_ECHO)
		time.Sleep(25 * time.Millisecond)
		dist2 += getDistance(SENSOR2_TRIGGER, SENSOR2_ECHO)
		dist2 /= 2

		// Relay On! (Negative means Timed out) - Sensor1 or Sensor2
		if dist1 <= sensorConf.s1_activeDist && dist1 > 0 {
			relayIsOn = true
			RELAY_TRIGGER.High()
			id, _ := uuid.NewUUID()

			CreateEventDb(db, id.String(), 1, EventActive, dist1)
		} else if dist2 <= sensorConf.s2_activeDist && dist2 > 0 {
			relayIsOn = true
			RELAY_TRIGGER.High()
			id, _ := uuid.NewUUID()

			CreateEventDb(db, id.String(), 2, EventActive, dist2)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Load in .env Variables
	CheckError(godotenv.Load(), "dotenv Configure Error Loading .env File", true)

	// Initialize Database
	db = InitDatabase(DatabaseOptions{
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PSWD"),
		os.Getenv("PSQL_DB"),
	})
	defer db.Close()

	// Try to Create Relays in DB
	CreateRelayDb(db, 1, "Bottom Relay", "Sensor at thy bottom")
	CreateRelayDb(db, 2, "Top Relay", "Sensor at thy top")

	// Init RPIO
	if err := rpio.Open(); err != nil {
		fmt.Println("Run as Sudo!")
		fmt.Printf("Error: %v\n", err)
	}
	defer rpio.Close()

	// Initiate Pin Modes
	RELAY_TRIGGER.Output()

	// Listen to Sensors
	listenToSensors(50 * time.Millisecond)
	// go listenToSensors(50 * time.Millisecond)

	// TODO: Do other tasks here...
	// time.Sleep(1 * time.Hour)
}
