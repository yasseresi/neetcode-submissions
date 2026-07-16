func groupAnagrams(strs []string) [][]string {
    result := [][]string{}
    visited := make([]bool, len(strs))

    for i, str := range strs {
        // Skip if this word is already in another group
        if visited[i] {
            continue
        }

        group := []string{str}
        visited[i] = true

        // Find all anagrams of strs[i]
        for j := i + 1; j < len(strs); j++ {
            if visited[j] {
                continue
            }

            if Anagram(str, strs[j]) {
                group = append(group, strs[j])
                visited[j] = true
            }
        }

        result = append(result, group)
    }

    return result
}

func Anagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    counts := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        counts[s[i]]++
    }

    for i := 0; i < len(t); i++ {
        counts[t[i]]--
    }

    for _, v := range counts {
        if v != 0 {
            return false
        }
    }

    return true
}