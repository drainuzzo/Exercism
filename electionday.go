package main

import (
	"fmt"
	"strconv"
)

func main() {

	var initialVotes int
	initialVotes = 2

	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println("counter=", counter)
	fmt.Println("*counter=", *counter)
	//*counter == initialVotes // true

	var votes int
	votes = 3
	var voteCounter *int
	voteCounter = &votes

	fmt.Println("votecounter ", VoteCount(voteCounter))
	// Output: 3

	var nilVoteCounter *int
	fmt.Println("votecounter nil ", VoteCount(nilVoteCounter))
	// Output: 0

	//votes = 2
	IncrementVoteCount(voteCounter, 2)
	fmt.Println("votes ", votes)
	fmt.Println("* voteCounter ", *voteCounter)

	result := NewElectionResult("Peter", 3)
	fmt.Println("NewElectionResult ", result)
	//result.Name == "Peter"  // true
	//result.Votes == 3       // true

	result = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}
	fmt.Println("DisplayResult ", DisplayResult(result))
	// Output: John (32)

	finalResults := map[string]int{
		"Mary": 10,
		"John": 51,
	}
	DecrementVotesOfCandidate(finalResults, "Mary")
	fmt.Println("finalResults[Mary] ", finalResults["Mary"])
	// Output: 9
}

//package electionday
type ElectionResult struct {
	// Name of the candidate
	Name string
	// Number of votes the candidate had
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	var votes int = 0
	if counter != nil {
		votes = *counter
	}
	return votes

}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment

}

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{candidateName, votes}
}

// NewElectionResult creates a new election result
// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	return result.Name + " (" + strconv.Itoa(result.Votes) + ")"
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}

/*
So far we've only seen pointers to primitive values. We can also create pointers for structs:

type Person struct {
    Name string
    Age  int
}

var peter Person
peter = Person{Name: "Peter", Age: 22}

var p *Person
p = &peter
We could have also created a new Person and immediately stored a pointer to it:

var p *Person
p = &Person{Name: "Peter", Age: 22}
When we have a pointer to a struct, we don't need to dereference the pointer before accessing one of the fields:

var p *Person
p = &Person{Name: "Peter", Age: 22}

fmt.Println(p.Name) // Output: "Peter"
                    // Go automatically dereferences 'p' to allow
                    // access to the 'Name' field
*/
