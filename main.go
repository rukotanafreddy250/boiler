package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

var (
	workingDir    string
	projectName   string
	language      int
	isRubocop     string
	isTests       string
	testFramework int
	isGithub      string
)

func main() {
	fmt.Println("Welcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")

	// working directory
	fmt.Println("Enter the working directory:")
	fmt.Scan(&workingDir)

	// project name
	fmt.Println("\n\nWhat is the project name you want to use?")
	fmt.Scan(&projectName)

	// choose a language
	fmt.Println("Choose a number which correspond to the language you will be using:\n1.Ruby")
	fmt.Scan(&language)
	if language != 1 {
		for i := 0; i < 5; i++ {
			fmt.Println("Choose a number which correspond to the language you will be using:\n1.Ruby")
			fmt.Scan(&language)
			if language == 1 {
				break
			}
		}
		fmt.Println("The language you chose is not supported")
		return
	}
	fmt.Println("Will you use Rubocop as a linter? Enter y for yes or any other key for no")
	fmt.Scan(&isRubocop)

	// will you run tests?
	fmt.Println("Will you write some unit tests for your project? Enter y for yes or any other key for no")
	fmt.Scan(&isTests)
	if isTests == "y" || isTests == "Y" {
		fmt.Println("Choose a number which corresponds to the testing framework you will be using:\n1.RSpec")
		fmt.Scan(&testFramework)
		if testFramework != 1 {
			for i := 0; i < 5; i++ {
				fmt.Println("Choose a number which corresponds to the testing framework you will be using:\n1.RSpec")
				fmt.Scan(&testFramework)
				if testFramework == 1 {
					break
				}
			}
			fmt.Println("The the testing framework you chose is not supported")
		}
	}

	// will you use github?
	fmt.Println("Will you use github as a collaboration tool? Enter y for yes or any other key for no")
	fmt.Scan(&isGithub)

	fmt.Println("\n\n\nThe following are your preferences, we will setup your project depending on your preferences.")
	fmt.Printf("Working dir : %v\n", workingDir)
	fmt.Printf("Project name : %v\n", projectName)
	fmt.Printf("Language name : %v\n", language)
	fmt.Printf("Will you write unit test? : %v\n", isTests)
	fmt.Printf("Testing framework : %v\n", testFramework)
	fmt.Printf("Will you use github? : %v\n\n", isGithub)

	// changing working dir
	fmt.Printf("Getting your home directory")
	homeDirectory, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	wrkDr := homeDirectory + "/" + workingDir + "/" + projectName

	// create a project directory
	fmt.Printf("Creating directory to %s...\n", projectName)
	os.Mkdir(wrkDr, 0755)

	// initialize rubocop
	fmt.Printf("Initializing rubocop in %s directory...\n", projectName)
	copy("./lib/.ruby/.rubocop.yml", wrkDr+"/.rubocop.yml")

	// // initialize github actions
	// defer fmt.Printf("Initializing github actions in %s directory...\n", projectName)
	// defer os.Mkdir(projectName+"/.github", 0755)
	// defer os.Mkdir(projectName+"/.github/workflows", 0755)
	// defer copy("./lib/.ruby/.github/workflows/linters.yml", projectName+"/.github/workflows/linters.yml")
	// defer copy("./lib/.ruby/.github/workflows/tests.yml", projectName+"/.github/workflows/tests.yml")

	// // create a readme file
	// defer fmt.Printf("Creating README file in %s directory...\n", projectName)
	// defer copy("./lib/.ruby/README.md", projectName+"/README.md")

	// // change working dir
	// defer os.Chdir(projectName)

	// // initialize gemfile
	// defer fmt.Printf("Initializing gem in %s directory...\n", projectName)
	// defer exec.Command("bundle", "init").Run()

	// // initialize rspec
	// defer fmt.Printf("Initializing rspec in %s directory...\n", projectName)
	// defer exec.Command("rspec", "--init").Run()

	// // initialize git
	// defer fmt.Printf("Initializing git in %s directory...\n", projectName)
	// defer exec.Command("git", "init").Run()

	// // Displaying last commands
	// fmt.Println("\n\nYour project has been initialized successfully")
	// fmt.Println("The remaining task is to go on github and create a repository and copy its url")
	// fmt.Printf("Come back in the root directory of %s\n", projectName)
	// fmt.Println("\nRun the following commands respectifuly")
	// fmt.Println("1. git remote add .")
	// fmt.Println("2. git commit -m \"Initial commit\"")
	// fmt.Println("3. git remote add origin [Paste the url you copied from github]")
	// fmt.Println("4. git push -u origin master")
	// fmt.Print("\n\nCongratulations and good luck for your new project\n\n\n")
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
