package main

import (
	"fmt"

	"github.com/devakxhay/ollama-playground/internal/modules"
)

func main() {
	currentModule := 4

	if currentModule == 1 {
		fmt.Println("------------ GENERATE API EXAMPLE ------------")

		modules.GenerateApiExample("what is TLS?")
		modules.GenerateApiExample("Can you elaborate it?")

		fmt.Println("------------ CHAT API EXAMPLE ------------")

		modules.ChatApiExample("what is TLS?")
		modules.ChatApiExample("Can you elaborate it?")
	}

	if currentModule == 2 {
		fmt.Println("------------ CHAT STREAM API EXAMPLE ------------")

		fmt.Println("1/2 -- What is TLS?")
		modules.ChatStreamApiExample("what is TLS?")

		fmt.Println("2/2 -- Can you elaborate it?")
		modules.ChatStreamApiExample("Can you elaborate it?")
	}

	if currentModule == 3 {
		fmt.Println("------------ CHAT JSON RESPONSE API EXAMPLE ------------")
		modules.GenerateJsonResponseExample("Analyse the text: Book the flight to New Delhi for 2 people on 24th of December. Return the JSON object with the following keys: departure, destination, number_of_passengers, flight_date")

		modules.GenerateJsonResponseExample("Analyse the text: I spend 200rs on food. Return the JSON object with the following keys: amount,currency,expense_type")
	}

	if currentModule == 4 {
		fmt.Println("------------ SYSTEM ROLE EXAMPLE ------------")
		modules.InitChat("You are a database query designer. Return only raw SQL queries.")

		modules.SystemRoleExample("Show me the list of all users.")
		modules.SystemRoleExample("Top 5 users having most followers.")
	}
}
