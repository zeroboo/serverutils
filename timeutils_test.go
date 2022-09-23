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
	bod := Time.GetTimeBeginOfDay(currentTime)
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

	bod := Time.GetTimeBeginOfDay(testTime)
	log.Printf("BeginOfDay of %v is %v", testTime, bodTime)

	assert.Equal(t, bod, bodTime, "Expected bod")
}

func TestGetSecondsSinceBeginOfDay_PresetDate_CorrectValue(t *testing.T) {
	log.Println("TestGetTimeBeginOfDay_CorrectValue")
	log.Println("TestGetTimeBeginOfDay_CorrectValue")

	testTime, errTestTime := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")
	log.Printf("Prepare test done, errTestTime=%v", errTestTime)
	secondsSinceBOD := Time.GetSecondsSinceBeginOfDay(testTime)
	log.Printf("Seconds since begin of day of `%v` is %v", testTime, secondsSinceBOD)

	assert.Equal(t, int64(16219), secondsSinceBOD, "Expected seconds")
}

func TestFormatLogTime(t *testing.T) {
	log.Println("TestFormatLogTime")
	testTime, _ := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")

	formattedTime := Time.FormatTimeLogFile(testTime)

	//formatted
	assert.Equal(t, "19450902_043019", formattedTime, "Correct value")
}
func TestFormatLogTime(t *testing.T) {
	log.Println("TestFormatLogTime")
	testTime, _ := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")

	formattedTime := Time.GetDate(testTime)

	//formatted
	assert.Equal(t, "19450902_043019", formattedTime, "Correct value")
}

// go test -timeout 30s -run ^TestFormatLogTime$ github.com/zeroboo/serverutils
func TestRandomString_CorrectSize(t *testing.T) {
	log.Println("TestRandomString_CorrectSize")

	assert.Equal(t, 4, len(Random.GetRandomString(4)), "Correct size")
	assert.Equal(t, 16, len(Random.GetRandomString(16)), "Correct size")
}
