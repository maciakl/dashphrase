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

const version = "0.1.1"

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

	digitsFlag = flag.Int("d", 3, "Number of digits to append")
	wordsFlag = flag.Int("w", 3, "Number of words in the phrase")
	versionFlag = flag.Bool("v", false, "Print version information and exit")
	helpFlag = flag.Bool("h", false, "Print usage and exit")
	countFlag = flag.Bool("c", false, "Show character count")
	nameFlag = flag.Bool("n", false, "Use a name in the phrase")
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


	if *wordsFlag < 1 {
		return fmt.Errorf("number of words must be at least 1")
	}

	if *digitsFlag < 1 {
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

	var dashphrase string

	if *nameFlag {
		dashphrase = getNamedDashPhrase(*wordsFlag, *digitsFlag)
	} else {
		dashphrase = getDashPhrase(*wordsFlag, *digitsFlag)
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
func getDashPhrase(words int, digits int) string {
	phrase := getPhrase(words)
	phrase += "-" + getNum(digits)
	return phrase
}

// Generate a dash-separated passphrase with a name, specified number of words and digits
func getNamedDashPhrase(words int, digits int) string {
	phrase := getNamedPhrase(words)
	phrase += "-" + getNum(digits)
	return phrase
}


// Generate a string of random digits of specified length
func getNum(digits int) string {
	num := ""
	for i := 0; i < digits; i++ {
		digit := rand.Intn(10)
		num += fmt.Sprintf("%d", digit)
	}
	return num
}

// Generate a dash-separated phrase with specified number of words
func getPhrase(words int ) string {
	phrase := ""
	for i := 0; i < words; i++ {
		if i > 0 {
			phrase += "-"
		}
		phrase += getWord()
	}
	return phrase
}

// Generate a dash-separated phrase starting with a name followed by specified number of words
func getNamedPhrase(words int) string {
	phrase := getName() + "-" + getPhrase(words-1)
	return phrase
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
