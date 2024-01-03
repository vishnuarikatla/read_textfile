package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type JobInformation struct {
	Title    string
	Duties   string
	Location string
	Pay      string
	Client   string
}

func main() {
	//Replace "input.txt" with the path of the text file
	filename := "C:\\Users\\arika\\OneDrive\\Desktop\\Job_title.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //NewScanner?   to read the file line by line

	var jobInfo JobInformation

	for scanner.Scan() {  //scanner.Scan()?  iterates through each line of the file
		line := scanner.Text()  //scanner.Text()?  retrives current line from scanner

		if title := ExtractJobTitle(line); title != "" {
			jobInfo.Title = title
		}
		if duties := ExtractDuties(line); duties != "" {
			jobInfo.Duties = duties
		}
		if location := ExtractLocation(line); location != "" {
			jobInfo.Location = location
		}
		if pay := ExtractPay(line); pay != "" {
			jobInfo.Pay = pay
		}
		if client := ExtractClient(line); client != "" {
			jobInfo.Client = client
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	// Print the extracted information
	fmt.Println("Job Title:", jobInfo.Title)
	fmt.Println("Duties:", jobInfo.Duties)
	fmt.Println("Location:", jobInfo.Location)
	fmt.Println("Pay:", jobInfo.Pay)
	fmt.Println("Client:", jobInfo.Client)
}

func ExtractJobTitle(line string) string {
	//implement job title extraction logic using regular expressions
	//Replace the pattern accordingly
	pattern := regexp.MustCompile(`(?:Job Title|Title): (.+)`)
	match := pattern.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func ExtractDuties(line string) string {
	//implement Duties extraction logic using regular expressions
	//Replace the pattern accordingly
	pattern := regexp.MustCompile(`(?:Duties|Responsibilities): (.+)`)
	match := pattern.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func ExtractLocation(line string) string {
	//implement job title extraction logic using regular expressions
	//Replace the pattern accordingly
	pattern := regexp.MustCompile(`(?:Location|Location): (.+)`)
	match := pattern.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func ExtractPay(line string) string {
	//implement job title extraction logic using regular expressions
	//Replace the pattern accordingly
	pattern := regexp.MustCompile(`(?:Pay|Salary): (.+)`)
	match := pattern.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func ExtractClient(line string) string {
	//implement job title extraction logic using regular expressions
	//Replace the pattern accordingly
	pattern := regexp.MustCompile(`(?:Client|Customer): (.+)`)
	match := pattern.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
