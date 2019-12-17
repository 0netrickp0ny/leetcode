package main

import (
    "fmt"
    "sort"
    "strings"
)

func sortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
    words := make(map[string][]string)
    for _, str := range strs {
        sortedStr := sortString(str)
        words[sortedStr] = append(words[sortedStr], str)
    }
    ans := make([][]string, len(words))
    ind := 0
    for _, word := range words {
        ans[ind] = append(ans[ind], word...)
        ind += 1
    }
    return ans
}
