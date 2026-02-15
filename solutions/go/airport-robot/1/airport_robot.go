package airportrobot
import "fmt"
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type German struct {}

func (g German) LanguageName() string {
    return "German"
}

func (g German) Greet(text string) string {
    return fmt.Sprintf("Hallo %s!", text)
}

type Italian struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(text string) string {
    return fmt.Sprintf("Ciao %s!", text)
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(text string) string {
    return fmt.Sprintf("Olá %s!", text)
}