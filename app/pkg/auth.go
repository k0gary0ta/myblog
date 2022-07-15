package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

func GenerateJWTToken(id string, password string) (string, error) {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("error opening env file: ", envErr)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"sub":  id,
		"name": password,
		"iat":  time.Now().Unix(),                     // Token issue date
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	signed, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ValidateJWTToken(token string) (bool, error) {
	parsed, err := jwt.Parse(token, func(parsed *jwt.Token) (interface{}, error) {
		if _, ok := parsed.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsed.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok {

		// Validate token expiration date
		if !parsed.Valid {
			fmt.Printf("exp: %v", int64(claims["exp"].(float64)))
			return false, fmt.Errorf("token expiration error")
		}

		// Validate username and password
		if claims["sub"] == os.Getenv("ADMIN_USER_NAME") && claims["name"] == os.Getenv("ADMIN_USER_PASS") {
			return true, nil
		}
	}
	return false, fmt.Errorf("not a valid token")
}

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	},
})

var globalSessions *session.Manager

type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // return current sessionID
}

type Provider interface {
	SessionInit(sid string) (Session, error) // セッションの初期化
	// Session変数を返す
	// 存在しなければ新たにSession変数を作成して返す
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error // Session変数を破棄する
	SessionGC(maxLifeTime int64)     // 期限の切れたSessionを破棄
}

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ := manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxlifetime),
		}
		http.SetCookie(w, &cookie)
		return session
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ := manager.provider.SessionRead(sid)
		return session
	}
}
