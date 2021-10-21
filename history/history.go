package history

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/thohui/chis/trie"
)

type History struct {
	trie *trie.HistoryTrie
}

func New() *History {
	history := &History{trie.NewTrie()}
	history.initialize()
	return history
}

func (h *History) initialize() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	p := path.Join(home, ".zsh_history")
	file, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(file), "\n")
	for i := 0; i < len(split); i++ {
		str := split[i]
		index := strings.IndexRune(str, ';')
		if index != -1 {
			h.trie.Insert(split[i][index+1:])
		}
	}
}

func (h *History) Find(entry string) []string {
	r := h.trie.AutoComplete(entry)
	var results []string
	for k := range r {
		results = append(results, k)
	}
	return results
}
