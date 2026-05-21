package misc

// user agents are provided by https://github.com/MichaelJorky/top-user-agents-latest/blob/main/combined/random-generator-list.txt

import (
	_ "embed"
	"math/rand/v2"
	"strings"
)

//go:embed useragents.txt
var useragents string

// GenerateUserAgen provides a random usre agent from the useragents.txt
func GenerateUserAgent() (userAgent string) {
	useragentsArr := strings.Split(useragents, "\n")
	// i dont think we should return an err because there's no chance it can occur
	//#nosec G404
	return useragentsArr[rand.IntN(len(useragentsArr))]
}
