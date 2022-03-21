package auth

import (
	"encoding/base64"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/argon2"
)

const (
	AuthSalt         = "##QWERTYUIOPasdfghjklZxCvBnM1230984756##"
	AuthIter         = 3
	AuthMem          = 64 * 1024
	AuthThreads      = uint8(2)
	AuthSecretLength = uint32(32)
)

var (
	AuthSaltB64 = base64.RawStdEncoding.EncodeToString([]byte(AuthSalt))
)

func GenerateSecret(p string) (ehash string) {
	// pll stores the level of paralellism to allow. This shuold be equal to or less than the number of CPUs
	pll := uint8(runtime.NumCPU())
	if AuthThreads < pll {
		pll = AuthThreads
	}
	// Base64 encode the salt and hashed password.
	// Generate the has
	hash := argon2.IDKey([]byte(p), []byte(AuthSalt), AuthIter, AuthMem, pll, AuthSecretLength)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	ehash = fmt.Sprintf("$argon2id$v=%d$%s", argon2.Version, hashB64)

	return ehash
}

// Verify checks to see if the provided password p can be cryptographically
// matched to the stored secret s
func Verify(p string, s string) bool {
	// pll stores the level of paralellism to allow. This shuold be equal to or less than the number of CPUs
	hash := GenerateSecret(p)

	if string(hash) != s {
		return false
	} else {
		return true
	}
}

func WithJWT(hf func(c *gin.Context), role *string) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		bt := ctx.GetHeader("Authorization")
		if len(bt) == 0 {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		log.Println("Bearer Token:", bt)
		if role != nil {
			if !strings.Contains(bt, *role) {
				ctx.JSON(401, gin.H{"error": "Your role doesn't have access"})
				return
			}
		}
		hf(ctx)
	}
}
