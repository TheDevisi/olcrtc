package misc

import (
	_ "embed"
	"math/rand/v2"
	"strings"
)

//go:embed useragents.txt
var useragents string

// Provides a random usre agent from the useragents.txt
func GenerateUserAgent() (user_agent string) {
	useragentsArr := strings.Split(useragents, "\n")
	// i dont think we should return an err because there's no chance it can occur
	return useragentsArr[rand.IntN(len(useragentsArr))]
}
