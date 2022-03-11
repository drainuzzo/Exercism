package main

import "fmt"

func main() {
	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	var petDogIsPresent = false

	fmt.Println("CanFastAttack?", CanFastAttack(knightIsAwake))
	fmt.Println("CanSpy?", CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println("CanSignalPrisoner?", CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println("CanFreePrisoner?", CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
	// CanFastAttack can be executed only when the knight is sleeping

}

func CanFastAttack(knightIsAwake bool) bool {
	//Fast attack: a fast attack can be made if the knight is
	//sleeping, as it takes time for him to get his armor on, so he will be vulnerable.
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	//Spy: the group can be spied upon if at least one of them is awake.
	//Otherwise, spying is a waste of time.
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	//Signal prisoner: the prisoner can be signaled using bird sounds if the prisoner is awake and the
	//archer is sleeping, as archers are trained in bird signaling so they could intercept the message.
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//Free prisoner: Annalyn can try sneaking into the camp to free the prisoner. This is a risky thing to do
	//and can only succeed in one of two ways:
	//If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep.
	//The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.
	//If Annalyn does not have her dog then she and the prisoner must be very sneaky! Annalyn can
	//free the prisoner if the prisoner is awake and the knight and archer are both sleeping, but if the prisoner is sleeping they
	//can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.
	return ((!knightIsAwake && !archerIsAwake && prisonerIsAwake) || (petDogIsPresent && !archerIsAwake))
}
