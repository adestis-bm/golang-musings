package readfilelinebyline

import (
	"bufio"
	"os"
	"testing"
)

// TestReadFileLineByLine is based on <https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go>.
func TestReadFileLineByLine(t *testing.T) {
	filename := "example.txt"
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("os.Open(%q) failed with: %s", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t.Logf(" - %q", line)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scanner.Err() resulted in: %s", err)
	}
}
