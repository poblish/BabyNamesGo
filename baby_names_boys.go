package main

import "fmt"
import "unicode"
import "sort"

func main() {
    foreNames := sort.StringSlice{"Rohan","Nathaniel","Anthony","Chris","Jonathan","Lemur","Harry","Percy","Peregrine","James","Jamie","Sidney","Gabriel","Leyton","Curtley","Jarvis"}
    middleNames := sort.StringSlice{"Rohan","Nathaniel","Tony","Chris","Jonathan","Lemur","Harry","Percy","Peregrine","James","Jamie","Sidney","Gabriel","Leyton","Curtley","Jarvis"}
    count := 0

    sort.Sort(foreNames)
    sort.Sort(middleNames)

    for _, name1 := range foreNames {
        count++
        fmt.Println(count, "...", name1, "Jordan-Regan")

        syllables1 := syllableCount(name1)

        for _, name2 := range middleNames {
            if name1 != name2 {
                syllables2 := syllableCount(name2)

                if name1[0] == name2[0] || (syllables1 == 1 && syllables2 == 1) || (syllables1 == 1 && syllables2 >= 3) || (syllables1 >= 3 && syllables2 >= 3) {
                    continue
                }

                count++
                fmt.Println(count, "...", name1, name2, "Jordan-Regan")
            }
        }
    }
}

func isVowel(inChar rune) bool {
    return (inChar == 'a' || inChar == 'e' || inChar == 'i' || inChar == 'o' || inChar == 'u')
}

func syllableCount(inStr string) int {
    if inStr == "Anthony" || inStr == "Gabriel" {
        return 3
    }

    if inStr == "James" {
        return 1
    }

    var runes_1 = []rune(inStr)

    syllables := 0
    lastChar := '.'
    wasVowel := false

    for i := 0; i < len(runes_1); i++ {
        c := unicode.ToLower(runes_1[i])

        if wasVowel && ((lastChar == 'u' && c == 'a') || (lastChar == 'i' && c == 'a')) {
            syllables++
        } else if isVowel(c) || c == 'y' {
            if !wasVowel && (!(c == 'e') || i < len(runes_1)-1) {
                syllables++
                wasVowel = true
            }
        } else {
            wasVowel = false
        }

        lastChar = c
    }

    if syllables == 0 {
        return 1
    }

    return syllables
}
