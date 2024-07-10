package column

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func chop(col int, line string) (ret string) {

	fields := strings.FieldsFunc(line, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	for i, field := range fields {
		if i == col {
			ret = field
		}
	}

	return ret
}

func Run() {
	var nFlag = flag.Int("n", 0, "0-indexed column to pick out")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()

		if err != nil || err == io.EOF {
			os.Exit(0)
		}

		fmt.Println(chop(*nFlag, string(line)))
	}

}
