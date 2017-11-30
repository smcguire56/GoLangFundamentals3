package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	questions := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!",
		"I am happy.",
		"I am not happy father with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?"}

	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')

	for i := 0; i < len(questions); i++ {
		fmt.Print("You: ")
		var text = questions[i]
		fmt.Println(questions[i])

		fmt.Print("Eliza: ")
		fmt.Println(ElizaResponse(text))
	}
}

func ElizaResponse(text string) string {
	response := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, text); matched {
		return "Why don’t you tell me more about your father?"
	}

	if matched, _ := regexp.MatchString(`(?i)i(?:'|\sa)?m (.*)`, text); matched {
		return "why are you like that?"
	}

	return response[rand.Intn(len(response))]
}

// adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
