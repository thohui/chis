package main

import (
	"os/exec"

	"github.com/c-bata/go-prompt"
	"github.com/thohui/chis/history"
)

var (
	his *history.History
)

func run(t string) {
	cancel := []string{":q", "quit", "exit"}
	for i := 0; i < len(cancel); i++ {
		if t == cancel[i] {
			return
		}
	}
	cmd := exec.Command("zsh", "-c", t)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			break
		}
	}
}

func completer(t prompt.Document) []prompt.Suggest {
	result := his.Find(t.Text)
	var suggestions []prompt.Suggest
	for i := 0; i < len(result); i++ {
		suggestions = append(suggestions, prompt.Suggest{Text: result[i]})
	}
	return suggestions
}

func main() {
	his = history.New()
	t := prompt.Input("> ", completer)
	run(t)
}
