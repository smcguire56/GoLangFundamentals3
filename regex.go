package main

import (
	"fmt"
    "math/rand"
    "time"
	"regexp"
)

var tests = 0

func main() {

	 questions := []string {
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?" } 

	//reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Ask Eliza: ")
	//text, _ := reader.ReadString('\n')
	var text = questions[tests]
	fmt.Println(ElizaResponse(text, tests))
	tests++
	main()
}

func ElizaResponse(text string, tests int) string{
	fmt.Println(tests)
    r, _ := regexp.Compile("father")

	if r.MatchString(text) && tests < 3 {
		return "Why don’t you tell me more about your father?"
	} else {
	var response[3] string 
	response[0] = "I’m not sure what you’re trying to say. Could you explain it to me?"
	response[1] = "How does that make you feel?"
	response[2] = "Why do you say that?"
	return response[randInt(0,3)]
	}
}

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}