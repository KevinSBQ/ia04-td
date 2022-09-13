package probleme

import (
	"os"
	"log"
	"bufio"
)

// Use os.open() function to open the file.
// Use bufio.NewScanner() function to create the file scanner.
// Use bufio.ScanLines() function with the scanner to split the file into lines.
// Then use the scanner Scan() function in a for loop to get each line and process it.

func DictFromFile(filename string) (dict []string) {
	file, err := os.Open(filename)
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		dict = append(dict, fileScanner.Text())
	}
	return
}