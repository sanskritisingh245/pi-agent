package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sanskritisingh/pi-agent/internal/agents"
	"github.com/sanskritisingh/pi-agent/internal/config"
	"github.com/sanskritisingh/pi-agent/internal/llm"
)

func main() {
	cfg := config.Load()

	client := llm.NewClient(
		cfg.ApiKey,
		cfg.Model,
	)

	agent, err := agent.New(client)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Pi Terminal")
	fmt.Println("Type exit to quit.")

	for{
		fmt.Println("\nYou > ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		if input == "exit"{
			fmt.Println("Goodbye!")
			break
		}

		reply, err := agent.Respond(input)
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Println("\nPi >", reply)
	}

}
