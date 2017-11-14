package main

import (
	"fmt"
    "math/rand"
	"regexp"
)

var tests = 0

func main() {

	 questions := []string {
		"I am happy.",
		"I am not happy father with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am not sure that you understand the father effect that your questions are having on me.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?" } 

	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')


    for i := 0; i < len(questions); i++ {
		fmt.Print("Ask Eliza: ")
		var text = questions[tests]
		fmt.Println(questions[tests])

		fmt.Println(ElizaResponse(text, tests))
		tests++
	}
}

func ElizaResponse(text string, tests int) string{
	response := []string { 
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?"}
	
	rand.Seed(int64(tests))

    r, _ := regexp.Compile("father")

	if r.MatchString(text) && tests < 3 {
		return "Why don’t you tell me more about your father?"
	} else {
	return response[rand.Intn(len(response))]
	}
}
