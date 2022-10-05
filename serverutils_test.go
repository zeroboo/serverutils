package serverutils

import (
	"log"
	"os"
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

func TestFormatLogTime_CorrectValue(t *testing.T) {
	log.Println("TestFormatLogTime_CorrectValue")
	testTime, _ := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")

	formattedTime := Time.FormatTimeLogFile(testTime)

	//formatted
	assert.Equal(t, "19450902_043019", formattedTime, "Correct value")
}

func TestGetDate_CorrectValue(t *testing.T) {
	log.Println("TestGetDate_CorrectValue")
	testTime, _ := time.Parse(TestTimeLayout, "1945-09-02T04:30:19.750")

	formattedTime := Time.GetDate(testTime)

	//formatted
	assert.Equal(t, "1945-09-02", formattedTime, "Correct formatted date")
}

// go test -timeout 30s -run ^TestFormatLogTime$ github.com/zeroboo/serverutils
func TestRandomString_CorrectSize(t *testing.T) {
	log.Println("TestRandomString_CorrectSize")

	assert.Equal(t, 4, len(Random.GetRandomString(4)), "Correct size")
	assert.Equal(t, 16, len(Random.GetRandomString(16)), "Correct size")
}

// go test -timeout 30s -run ^TestGetIntegerEnvOrDefault_CorrectValue$ github.com/zeroboo/serverutils -v
func TestGetIntegerEnvOrDefault_CorrectValue(t *testing.T) {
	log.Println("TestGetIntegerEnvOrDefault_CorrectValue")
	key := "TEST_VALUE"

	os.Setenv(key, "10")
	value, _ := GetIntegerEnvOrDefault(key, 9)
	assert.Equal(t, int64(10), value, "Correct result for valid number")

	os.Setenv(key, "10-")
	value, _ = GetIntegerEnvOrDefault(key, 9)
	assert.Equal(t, int64(9), value, "Default result for invalid number")

}

// go test -timeout 30s -run ^TestGetHashes_Correct$ github.com/zeroboo/serverutils -v
func TestGetHashes_Correct(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog"
	md5Hash := Hash.CreateMD5Hash([]byte(text))
	sha1Hash := Hash.CreateSHA1Hash([]byte(text))
	sha256Hash := Hash.CreateSHA256Hash([]byte(text))
	assert.Equal(t, "9e107d9d372bb6826bd81d3542a419d6", md5Hash, "Correct md5")
	assert.Equal(t, "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12", sha1Hash, "Correct sha1")
	assert.Equal(t, "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592", sha256Hash, "Correct sha256")

}
