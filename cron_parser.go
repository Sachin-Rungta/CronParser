package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CronParserData struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

func ParseField(field, allowedRange string) string {
	parts := strings.Split(field, ",")
	result := make([]string, 0)

	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			for i := start; i <= end; i++ {
				result = append(result, strconv.Itoa(i))
			}
		} else if strings.Contains(part, "/") {
			stepParts := strings.Split(part, "/")
			step, _ := strconv.Atoi(stepParts[1])
			start, _ := strconv.Atoi(strings.Split(allowedRange, "-")[0])
			end, _ := strconv.Atoi(strings.Split(allowedRange, "-")[1])
			for i := start; i <= end; i += step {
				result = append(result, strconv.Itoa(i))
			}
		} else if part == "*" {
			start, _ := strconv.Atoi(strings.Split(allowedRange, "-")[0])
			end, _ := strconv.Atoi(strings.Split(allowedRange, "-")[1])
			for i := start; i <= end; i++ {
				result = append(result, strconv.Itoa(i))
			}
		} else {
			val, _ := strconv.Atoi(part)
			result = append(result, strconv.Itoa(val))
		}
	}

	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: your-program \"cron_string\"")
		os.Exit(1)
	}

	cronString := os.Args[1]
	fields := strings.Fields(cronString)

	if len(fields) != 6 {
		fmt.Println("Invalid cron string")
		os.Exit(1)
	}

	printTable(&CronParserData{
		Minute:     ParseField(fields[0], "0-59"),
		Hour:       ParseField(fields[1], "0-23"),
		DayOfMonth: ParseField(fields[2], "1-31"),
		Month:      ParseField(fields[3], "1-12"),
		DayOfWeek:  ParseField(fields[4], "0-6"),
		Command:    fields[5],
	})
}

func printTable(cronParser *CronParserData) {
	fmt.Printf("%-14s%s\n", "minute", cronParser.Minute)
	fmt.Printf("%-14s%s\n", "hour", cronParser.Hour)
	fmt.Printf("%-14s%s\n", "day of month", cronParser.DayOfMonth)
	fmt.Printf("%-14s%s\n", "month", cronParser.Month)
	fmt.Printf("%-14s%s\n", "day of week", cronParser.DayOfWeek)
	fmt.Printf("%-14s%s\n", "command", cronParser.Command)
}
