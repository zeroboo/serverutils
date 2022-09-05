package serverutils

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeBeginOfDay_CurrentTime_CorrectValue(t *testing.T) {
	log.Println("TestGetTimeBeginOfDay_CurrentTime_CorrectValue")
	currentTime := time.Now()
	bod := GetTimeBeginOfDay(currentTime)
	log.Printf("BeginOfDay of %v is %v", currentTime, bod)

	assert.Equal(t, 0, bod.Hour(), "BOD: 0 hour")
	assert.Equal(t, 0, bod.Minute(), "BOD: 0 minute")
	assert.Equal(t, 0, bod.Second(), "BOD: 0 second")

}

const TestTimeLayout = "2006-01-02T15:04:05.000"

func TestGetTimeBeginOfDay_PresetDate_CorrectValue(t *testing.T) {
	log.Println("TestGetTimeBeginOfDay_CorrectValue")

	testTime, errTestTime := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")
	bodTime, errBODTime := time.Parse(TestTimeLayout, "1945-09-02T00:00:00.000")
	log.Printf("Prepare test done, errTestTime=%v, errBODTime=%v", errTestTime, errBODTime)
	assert.Equal(t, nil, errTestTime, "No parsing error")
	assert.Equal(t, nil, errBODTime, "No parsing error")

	bod := GetTimeBeginOfDay(testTime)
	log.Printf("BeginOfDay of %v is %v", testTime, bodTime)

	assert.Equal(t, bod, bodTime, "Expected bod")
}

func TestGetSecondsSinceBeginOfDay_PresetDate_CorrectValue(t *testing.T) {
	log.Println("TestGetTimeBeginOfDay_CorrectValue")
	log.Println("TestGetTimeBeginOfDay_CorrectValue")

	testTime, errTestTime := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")
	log.Printf("Prepare test done, errTestTime=%v", errTestTime)
	secondsSinceBOD := GetSecondsSinceBeginOfDay(testTime)
	log.Printf("Seconds since begin of day of `%v` is %v", testTime, secondsSinceBOD)

	assert.Equal(t, int64(16219), secondsSinceBOD, "Expected seconds")
}
