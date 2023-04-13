package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	part     string
	verify   bool
	pub      string
	secret   string
	showTime bool
)

// init params by flags
func init() {
	pflag.StringVarP(&part, "part", "p", "", "indicate part Number of jwt token,default(2)")
	pflag.BoolVarP(&verify, "verify", "v", false, "indicate  check whether,default(false)")
	pflag.BoolVarP(&showTime, "time", "t", false, "print time human readable ")
	pflag.StringVarP(&pub, "pub", "", "", "indicate signature pub key")
	pflag.StringVarP(&secret, "type", "s", "", "indicate signature secret")
	pflag.Usage = func() {
		fmt.Println(`Usage: jwt [pvs] [--pub]
Describe: 
	this is a simple jwt tool with checking token and print json indent format
Examples:
  # parse part NUMBER of jwt token  default 1&2
  echo "token" | jwt -p 1
  # verify token when the signature method is RSA256 that must be given pubKey file path
  echo "token" | jwt -v --pub sa.pub
  # verify token when the signature method is HSA256 that must be given a secert
  echo "token" | jwt -v -s 1234
Options:`)
		pflag.PrintDefaults()
	}
	pflag.Parse()
}

func main() {
	var (
		token string
	)
	sc := bufio.NewScanner(os.Stdin)

	for {
		if ok := sc.Scan(); ok {
			token = sc.Text()
			if strings.TrimSpace(sc.Text()) == "" {
				fmt.Println("the jwt command read data from os.stdin example: echo Token | jwt ")
				continue
			} else {
				break
			}
		}
	}

	if verify {
		verifySignature(token)
		os.Exit(0)
	}
	parseJWTToken(token)
}

func parseJWTToken(token string) {
	segments := splitSegment(token)
	switch part {
	case "1":
		toJsonAndPrint(segments[0])
	case "2":
		toJsonAndPrint(segments[1])
	default:
		toJsonAndPrint(segments[0], segments[1])
	}
}

func toJsonAndPrint(segments ...string) {
	for _, seg := range segments {
		dataMap := getMapBySegment(seg)
		if showTime {
			for _, key := range []string{"exp", "iat", "nbf"} {
				if value, ok := dataMap[key]; ok {
					dataMap[key] = time.Unix(int64(value.(float64)), 0).Format("2006/01/02 15:04:05")
				}
			}
		}
		data, err := json.MarshalIndent(dataMap, "", "  ")
		checkError(err)
		fmt.Println(string(data))
	}
}

func getSignatureMethod(segment string) string {
	dataMap := getMapBySegment(segment)
	return dataMap["alg"].(string)
}

func getMapBySegment(segment string) map[string]interface{} {
	base, err := base64.RawStdEncoding.DecodeString(segment)
	checkError(err)
	dataMap := make(map[string]interface{}, 0)
	checkError(json.Unmarshal(base, &dataMap))
	return dataMap
}

func verifySignature(token string) {
	alg := getSignatureMethod(splitSegment(token)[0])

	var err error
	switch alg {
	case "RS256":
		if strings.TrimSpace(pub) == "" {
			err = fmt.Errorf("error: not indicate pub key file by --pub")
			goto STOP
		}
		err = verifyToken(token, func(token *jwt.Token) (interface{}, error) {
			pubBytes, err := ioutil.ReadFile(pub)
			checkError(err)
			return jwt.ParseRSAPublicKeyFromPEM(pubBytes)
		})
	case "HS256":
		if strings.TrimSpace(secret) == "" {
			err = fmt.Errorf("error: not indicate secret by -s")
			goto STOP
		}
		err = verifyToken(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	}
STOP:
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Verify success")
	}
}

func verifyToken(tokenString string, keyFunc jwt.Keyfunc) (err error) {
	if _, err = jwt.Parse(tokenString, keyFunc); err != nil {
		return err
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func splitSegment(token string) []string {
	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		fmt.Println("token format err,please check token")
		os.Exit(1)
	}
	return segments
}
