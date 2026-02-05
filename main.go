package main    

import (
"os"
"fmt"
"flag"
"strings"
"math/rand"
"path/filepath"
"github.com/wordgen/wordlists"
)

const version = "0.1.3"

var (
	digitsFlag     *int 
	wordsFlag      *int
	versionFlag    *bool
	helpFlag       *bool
	countFlag      *bool
	nameFlag       *bool
	underscoreFlag *bool
)

func main() {

	flag.Usage = func() {
		Usage()
	}

	digitsFlag     = flag.Int("d", 3, "Number of digits to append")
	wordsFlag      = flag.Int("w", 3, "Number of words in the phrase")
	versionFlag    = flag.Bool("v", false, "Print version information and exit")
	helpFlag       = flag.Bool("h", false, "Print usage and exit")
	countFlag      = flag.Bool("c", false, "Show character count")
	nameFlag       = flag.Bool("n", false, "Use a name in the phrase")
	underscoreFlag = flag.Bool("u", false, "Use underscores instead of dashes")
	flag.Parse()

	err := run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// Main execution function
func run() error { 


	if *wordsFlag < 1 || *wordsFlag > 20 {
		return fmt.Errorf("number of words must be at least 1")
	}

	if *digitsFlag < 1 || *digitsFlag > 999 {
		return fmt.Errorf("number of digits must be at least 1")
	}

	// Handle version and help flags
	if *versionFlag {
		Version()
		return nil
	}
	if *helpFlag {
		Usage()
		return nil
	}

	var (
		dashphrase string
		err        error
	)

	if *nameFlag {
		dashphrase, err = getNamedDashPhrase(*wordsFlag, *digitsFlag)
	} else {
		dashphrase, err = getDashPhrase(*wordsFlag, *digitsFlag)
	}

	if err != nil {
		return err
	}

	if *countFlag {
		dashphrase = fmt.Sprintf("%s\t\t(%d ch)", dashphrase, len(dashphrase))
	}

	dashphrase = string(dashphrase[0]-32) + dashphrase[1:] // capitalize first char

	if *underscoreFlag {
		dashphrase = strings.ReplaceAll(dashphrase, "-", "_")
	}

	fmt.Println(dashphrase)
	return nil
}

// Generate a dash-separated passphrase with specified number of words and digits
func getDashPhrase(words int, digits int) (string, error) {

	if words < 1 || words > 20 {
		return "", fmt.Errorf("number of words must be between 1 and 20")
	}
	phrase, err := getPhrase(words)
	if err != nil {
		return "", err
	}

	number, e := getNum(digits)
	if e != nil {
		return "", e
	}
	phrase += "-" + number
	return phrase, nil
}

// Generate a dash-separated passphrase with a name, specified number of words and digits
func getNamedDashPhrase(words int, digits int) (string, error) {
	if words < 1 || words > 20 {
		return "", fmt.Errorf("number of words must be between 1 and 20")
	}

	phrase, err := getNamedPhrase(words)
	if err != nil {
		return "", err
	}

	number, e := getNum(digits)
	if e != nil {
		return "", e
	}
	phrase += "-" + number
	return phrase, nil
}


// Generate a string of random digits of specified length
func getNum(digits int) (string, error) {

	if digits < 1 || digits > 999 {
		return "", fmt.Errorf("number of digits must be between 1 and 20")
	}

	num := ""
	for i := 0; i < digits; i++ {
		digit := rand.Intn(10)
		num += fmt.Sprintf("%d", digit)
	}
	return num, nil
}

// Generate a dash-separated phrase with specified number of words
func getPhrase(words int ) (string, error) {
	if words < 1 || words > 20 {
		return "", fmt.Errorf("number of words must be between 1 and 20")
	}

	phrase := ""
	for i := 0; i < words; i++ {
		if i > 0 {
			phrase += "-"
		}
		phrase += getWord()
	}
	return phrase, nil
}

// Generate a dash-separated phrase starting with a name followed by specified number of words
func getNamedPhrase(words int) (string, error) {
	if words < 1 || words > 20 {
		return "", fmt.Errorf("number of words must be between 1 and 20")
	}

	phrase, err := getPhrase(words - 1)
	if err != nil {
		return "", err
	}

	named_phrase := getName() + "-" + phrase
	return named_phrase, nil
}

// Get a random word from the wordlist
func getWord() string {
	wordlist := wordlists.EffLarge
	randomIndex := rand.Intn(len(wordlist))
	return wordlist[randomIndex]
}

// Get a random name from the name wordlist
func getName() string {
	wordlist := wordlists.NamesMixed
	randomIndex := rand.Intn(len(wordlist))
	return wordlist[randomIndex]
}


// Print version information
func Version() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
}

// Print usage information
func Usage() {
	fmt.Println("")
	Version()
	fmt.Println("\nA simple dash-separated passphrase generator.")
    fmt.Println("\nUsage:")
	fmt.Println("")
	flag.PrintDefaults()
}
