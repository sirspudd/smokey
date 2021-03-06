package builtins

import (
	"github.com/CrimsonAS/smokey/lib"
)

// Remove duplicates from input, write the ordered (but unique) inputs to output.
type UniqCmd struct {
}

func (this UniqCmd) Call(inChan, outChan *lib.Channel, arguments []string) {
	dat := make(map[interface{}]lib.ShellData, 1024)
	orderedDat := make([]lib.ShellData, 0, 1024)

	for in, ok := inChan.Read(); ok; in, ok = inChan.Read() {
		_, ok := dat[in]
		if !ok {
			dat[in] = in
			orderedDat = append(orderedDat, in)
		}
	}

	for _, out := range orderedDat {
		outChan.Write(out)
	}
	outChan.Close()
}
