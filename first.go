package main

import "fmt"

func checkVotingEligibility(age int){
    fmt.Println("Checking age", age, "...")
    if age >= 18{
        fmt.Println("You can vote!")
    } else{
        fmt.Println("You can't vote.")
    }
}

func main(){
 fmt.Println("Hello")
    name := "anas"
    age := 22
    isLearningGo := true

    fmt.Println("My name is", name)
    fmt.Println("I am ", age ," years old")
    fmt.Println("Am i learning go? " , isLearningGo)

    age = 20

    if age >= 18 {
        fmt.Println("You can vote!")
    } else {
        fmt.Println("Sorry you can't vote")
    }

    checkVotingEligibility(22)
    checkVotingEligibility(5)
    checkVotingEligibility(100)
}