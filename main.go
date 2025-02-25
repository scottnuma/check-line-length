package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type violation struct {
	filename string
	lineNum  int
	length   int
}

var (
	maxLength int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "line-length-checker [files...]",
		Short: "Check line length in text files",
		Long:  "Check if any lines in the provided files exceed a specified maximum length",
		RunE:  run,
	}

	rootCmd.Flags().IntVar(&maxLength, "max-line-length", 80, "Maximum allowed line length")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no files provided")
	}

	var violations []violation

	// Check each file
	for _, filename := range args {
		fileViolations, err := checkFile(filename, maxLength)
		if err != nil {
			return fmt.Errorf("error checking %s: %v", filename, err)
		}
		violations = append(violations, fileViolations...)
	}

	// Report violations and exit with appropriate code
	if len(violations) > 0 {
		for _, v := range violations {
			fmt.Printf("%s:%d: Line too long (%d > %d characters)\n",
				v.filename, v.lineNum, v.length, maxLength)
		}
		os.Exit(1)
	}

	return nil
}

func checkFile(filename string, maxLength int) ([]violation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var violations []violation
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		lineLength := len(line)

		if lineLength > maxLength {
			violations = append(violations, violation{
				filename: filename,
				lineNum:  lineNum,
				length:   lineLength,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return violations, nil
}
