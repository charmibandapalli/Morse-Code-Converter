package main
import "fmt"

var morseMap = map[rune]string{
    'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..",
}

func toMorse(text string) string {
    var morse string
    for _, char := range text {
        morse += morseMap[char] + " "
    }
    return morse
}

func main() {
    fmt.Println(toMorse("ABCD"))
}
