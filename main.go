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

const version = "0.1.0"

func main() {

	flag.Usage = func() {
		Usage()
	}

	err := run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func run() error { 

	digits := flag.Int("d", 3, "Number of digits to append")
	words := flag.Int("w", 3, "Number of words in the phrase")
	versionFlag := flag.Bool("v", false, "Print version information and exit")
	helpFlag := flag.Bool("h", false, "Print usage and exit")
	count := flag.Bool("c", false, "Show character count")
	name := flag.Bool("n", false, "Use a name in the phrase")
	underscore := flag.Bool("u", false, "Use underscores instead of dashes")
	flag.Parse()

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

	if *name {
		dashphrase = getNamedDashPhrase(*words, *digits)
	} else {
		dashphrase = getDashPhrase(*words, *digits)
	}

	if *count {
		dashphrase = fmt.Sprintf("%s\t\t(%d ch)", dashphrase, len(dashphrase))
	}

	dashphrase = string(dashphrase[0]-32) + dashphrase[1:] // capitalize first char

	if *underscore {
		dashphrase = strings.ReplaceAll(dashphrase, "-", "_")
	}

	fmt.Println(dashphrase)

	return nil
}

func getDashPhrase(words int, digits int) string {
	phrase := getPhrase(words)
	phrase += "-" + getNum(digits)
	return phrase
}

func getNamedDashPhrase(words int, digits int) string {
	phrase := getNamedPhrase(words)
	phrase += "-" + getNum(digits)
	return phrase
}


func getNum(digits int) string {
	num := ""
	for i := 0; i < digits; i++ {
		digit := rand.Intn(10)
		num += fmt.Sprintf("%d", digit)
	}
	return num
}


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

func getNamedPhrase(words int) string {
	phrase := getName() + "-" + getPhrase(words-1)
	return phrase
}

func getWord() string {
	wordlist := wordlists.EffLarge
	randomIndex := rand.Intn(len(wordlist))
	return wordlist[randomIndex]
}

func getName() string {
	wordlist := wordlists.NamesMixed
	randomIndex := rand.Intn(len(wordlist))
	return wordlist[randomIndex]
}


func Version() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
}

func Usage() {
	fmt.Println("")
	Version()
	fmt.Println("\nA simple dash-separated passphrase generator.\n")
    fmt.Println("Usage:\n")
	flag.PrintDefaults()
}
