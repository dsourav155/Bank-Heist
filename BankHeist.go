package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")

		fmt.Println("EludedGuards has a value of", eludedGuards)
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
		fmt.Println("EludedGuards has a value of", eludedGuards)
	}
	openedVault := rand.Intn(100)
	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and Go!")
		fmt.Println("EludedGuards has a value of", openedVault)
	} else if isHeistOn == true && openedVault < 70 {
		isHeistOn = false
		fmt.Println("Vault can't be opened!")
		fmt.Println("EludedGuards has a value of", openedVault)
	}
	leftSafety := rand.Intn(5)
	if isHeistOn {
		switch leftSafety {
		case 0:
			isHeistOn = false
			fmt.Println("Security caught us, Heist Failed!")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside, Heist Failed!")
		case 2:
			isHeistOn = false
			fmt.Println("One of the member back stabed the team and told the plan to the security, Heist failed!")
		case 3:
			isHeistOn = false
			fmt.Println("Fingerprint scanner is not accepting any fingerprint, we got stuck, Heist failed")
		default:
			fmt.Println("Start the getaway car!")
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We got a total of", amtStolen, "dollars")
	}
	fmt.Println("What's the current robbery status", isHeistOn)
}
