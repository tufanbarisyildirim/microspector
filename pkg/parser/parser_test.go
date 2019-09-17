package parser

import (
	"fmt"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest() *httptest.Server {
	serverMux := http.NewServeMux()
	server := httptest.NewServer(serverMux)

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Microspector", "Service Up")
		w.Header().Set("User-Agent", r.Header.Get("User-Agent"))
		w.Header().Set("Host", r.Header.Get("Host"))
		fmt.Fprint(w, "Microspector")

	})

	return server
}
func TestParser_Set(t *testing.T) {

	lex := Parse(`
SET {{ Domain }} "microspector.com"
SET {{ ContainsTrue }}  "microspector.com" CONTAINS "microspector"
SET {{ ContainsFalse }} "microspector.com" CONTAINS "microspectorFAIL"
SET {{ StartsWithTrue }} "microspector.com" STARTSWITH "microspector"
SET {{ StartsWithFalse }} "microspector.com" STARTSWITH "microspectorFAIL"
SET {{ DoubleDomain }} "microspector.com {{ .Domain }}"
SET {{ Hundred }} 100
SET {{ StringDigitCompare1 }} "100" LT 101
SET {{ StringDigitCompare2 }} "100" < 101
SET {{ StringDigitCompare3 }} "100" <= 101
SET {{ StringDigitCompare4 }} "100" GT 99
SET {{ StringDigitCompare5 }} "100" > 99
SET {{ StringDigitCompare6 }} "100" >= 99
SET {{ StringDigitCompare7 }} "100" <= 99
SET {{ StringDigitCompare8 }} "100" == 100
SET {{ StringDigitCompare9 }} "100" != 100
SET {{ StringDigitCompare10 }} 100 GT "99"
SET {{ StringDigitCompare11 }} {{ Hundred }} GT "99"
SET {{ StringDigitCompare12 }} {{ Hundred }} GT "999"
SET {{ WhenFalse }} FALSE WHEN "100" < "101"
SET {{ SSLRand }} "{{ openssl_rand 32 \"hex\" }}"
SET {{ SSLRandSize }} "{{ str_len .SSLRand }}"
SET {{ HashMd5 }} "{{ hash_md5 1 }}"
SET {{ Hash256 }} "{{ hash_sha256 .HashMd5 }}"
`)

	Run(lex)

	assert.Equal(t, GlobalVars["Domain"], "microspector.com")
	assert.Equal(t, GlobalVars["ContainsTrue"], true)
	assert.Equal(t, GlobalVars["ContainsFalse"], false)
	assert.Equal(t, GlobalVars["StartsWithTrue"], true)
	assert.Equal(t, GlobalVars["StartsWithFalse"], false)
	assert.Equal(t, GlobalVars["DoubleDomain"], "microspector.com microspector.com")
	assert.Equal(t, GlobalVars["StringDigitCompare1"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare2"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare3"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare4"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare5"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare6"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare7"], false)
	assert.Equal(t, GlobalVars["StringDigitCompare8"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare9"], false)
	assert.Equal(t, GlobalVars["StringDigitCompare10"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare11"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare12"], false)
	assert.Equal(t, GlobalVars["WhenFalse"], false)
	assert.Equal(t, GlobalVars["SSLRandSize"], "64")
	assert.Equal(t, len(GlobalVars["HashMd5"].(string)), 32)
	assert.Equal(t, len(GlobalVars["Hash256"].(string)), 64)
}

func TestParser_Http(t *testing.T) {
	server := setupTest()
	defer server.Close()

	GlobalVars["ServerMux"] = server.URL

	Run(Parse(`
HTTP GET {{ ServerMux }} HEADER "User-Agent:(bot)microspector.com" INTO {{ ServerResult }}
SET {{ ContentLength }} {{ ServerResult.ContentLength }}
	`))

	assert.Equal(t, GlobalVars["ServerMux"], server.URL)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).ContentLength, 12)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).StatusCode, 200)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).Headers["UserAgent"], "(bot)microspector.com")
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).Headers["Microspector"], "Service Up")

}
