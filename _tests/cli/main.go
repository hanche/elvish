// A test program for the cli package.
package main

import (
	"fmt"
	"io"
	"unicode"

	"github.com/elves/elvish/cli"
	"github.com/elves/elvish/edit/ui"
	"github.com/elves/elvish/styled"
)

func highlight(code string) styled.Text {
	t := styled.Text{}
	for _, r := range code {
		style := ""
		if unicode.IsDigit(r) {
			style = "green"
		}
		t = append(t, styled.MakeText(string(r), style)...)
	}
	return t
}

func main() {
	app := cli.NewApp(&cli.AppConfig{
		Prompt:      cli.ConstPlainPrompt("> "),
		Highlighter: cli.FuncHighlighterNoError(highlight),
		InsertConfig: cli.InsertModeConfig{
			Binding: cli.MapBinding(map[ui.Key]cli.KeyHandler{
				ui.K('D', ui.Ctrl): cli.CommitEOF,
				ui.Default:         cli.DefaultInsert,
			}),
		},
	})

	for {
		code, err := app.ReadCode()
		if err != nil {
			if err != io.EOF {
				fmt.Println("error:", err)
			}
			break
		}
		fmt.Println("got:", code)
	}
}
