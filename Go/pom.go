package main

import (
    "fmt"
    "time"
)

func main() {
    pomodoroDuration := 1 * time.Minute
    breakDuration := 2 * time.Minute

    fmt.Println("Starting Pomodoro timer for", pomodoroDuration)
    time.Sleep(pomodoroDuration)
    fmt.Println("Pomodoro complete! Taking a break for", breakDuration)
    time.Sleep(breakDuration)
    fmt.Println("Break complete! Ready to start the next Pomodoro.")
}