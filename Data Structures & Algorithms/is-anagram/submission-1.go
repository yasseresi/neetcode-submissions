import "maps"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sMap := make(map[byte]int)
    tMap := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        sMap[s[i]]++
        tMap[t[i]]++
    }

    return maps.Equal(sMap, tMap)
}