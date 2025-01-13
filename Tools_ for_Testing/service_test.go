package Tools__for_Testing

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// Demonstrates a basic test function
func TestMyTest(t *testing.T) {
	t.Log("TestMyTest executed successfully")
}

// Demonstrates running tests concurrently
func TestMyParaTest(t *testing.T) {
	t.Parallel()
	t.Log("Running concurrent test")
}

// Fails the current test but allows other tests to continue
func TestMyFailedTest(t *testing.T) {
	t.Error("Intentional failure for demonstration")
}

// Fails the test and stops further execution
func TestMyBailTest(t *testing.T) {
	t.Fatal("Test execution halted due to fatal error")
}

// Logs output, visible if the test fails or with verbose mode
func TestLogTest(t *testing.T) {
	t.Log("This is a log message for demonstration")
}

// Helper function demonstration
func TestAddHelperFuncTest(t *testing.T) {
	t.Helper()
	t.Log("Helper function example")
}

// Creates a temporary directory for testing
func TestAddTempDirTest(t *testing.T) {
	tempDir := t.TempDir()
	t.Logf("Temporary directory created: %s", tempDir)
}

// Cleanup demonstration after a test completes
func TestAddCleanUpTest(t *testing.T) {
	t.Cleanup(func() {
		t.Log("Cleanup completed")
	})
}

// Compare structs or interfaces using cmp.Equal and cmp.Diff
func TestAddCompareTest(t *testing.T) {
	a := struct{ Name string }{"Alice"}
	b := struct{ Name string }{"Bob"}
	if !cmp.Equal(a, b) {
		t.Logf("Structures differ: %v", cmp.Diff(a, b))
	}
}

/* ### Command Usage:
1. **Test All Directories:**
`go test ./...`

2. **Verbose Test:**
`go test -v`

3. **Run Tests in Current Directory:**
`go test .`

4. **Rerun Tests by Overriding Cache:**
`go test -count=1`

5. **Run a Specific Test:**
`go test -run TestFunctionName`

*/
