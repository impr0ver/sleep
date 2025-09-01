package main
import (
        "fmt"
        "os"
        "time"
)
func main() {
        if len(os.Args) < 2 {
                fmt.Println("Error argument, please use:", os.Args[0]+" <3m>")
                os.Exit(1)
        }
        suffix := os.Args[1]
        Sleep(suffix)
}
func Sleep(suffix string) {
        t, _ := time.ParseDuration(suffix)
        fmt.Println("Now sleep: ", t, ".....")
        time.Sleep(t)
}