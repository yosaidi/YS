package reload

import (
	reload "reload/modifications"

	"strings"
)

func ProccessText(str string) string {
	strsplit := strings.Fields(str)
	i := 0
	for i < len(strsplit) {
		reload.Base(&strsplit)
		reload.Case(&strsplit, "(up", strings.ToUpper)
		reload.Case(&strsplit, "(cap", reload.Capitalize)
		reload.Case(&strsplit, "(low", strings.ToLower)
		reload.VowelCheck(&strsplit)
		reload.Ponctuation(&strsplit)
		reload.AdjustSingleQuotes(&strsplit)
		reload.CompleteSingleQuotes(&strsplit)
		i++
	}
	reload.CleanText(&strsplit, "(up")
	reload.CleanText(&strsplit, "(cap")
	reload.CleanText(&strsplit, "(low")
	reload.CleanText(&strsplit,"(hex")
	reload.CleanText(&strsplit,"(bin")

	return (strings.Join(strsplit, " "))

}
