package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/loadept/pirca"
)

type RequestBody struct {
	Source string `json:"source"`
	Event  string `json:"event"`
	Meta   Meta   `json:"meta"`
}

type Meta struct {
	Path    string `json:"path"`
	IP      string `json:"ip"`
	Country string `json:"country"`
	City    string `json:"city"`
}

func NewVisitsMiddleware(client *http.Client, secret, webhook string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := pirca.Ctx(r)
			next.ServeHTTP(w, r)

			if ctx.GetStatus() != http.StatusMovedPermanently && ctx.GetStatus() != http.StatusFound {
				return
			}

			ip := getClientIP(r)
			path := r.URL.Path
			go func() {
				location, err := getIPInfo(client, ip)
				if err != nil {
					return
				}
				body := RequestBody{
					Source: "loadept",
					Event:  "visit",
					Meta: Meta{
						Path:    path,
						IP:      ip,
						Country: location.Country,
						City:    location.City,
					},
				}
				sendWebhook(client, body, secret, webhook)
			}()
		})
	}
}

func sendWebhook(client *http.Client, body RequestBody, secret, webhook string) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	message := fmt.Sprintf("POST\n/notify\n%s\n%s", timestamp, jsonData)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	signature := hex.EncodeToString(mac.Sum(nil))

	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Timestamp", timestamp)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("an error occurred while sending webhook")
	}
	return nil
}

type ipLocation struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

func getIPInfo(client *http.Client, ip string) (*ipLocation, error) {
	urlPath := fmt.Sprintf("https://ipseek.space/%s", ip)

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred while querying IP address")
	}

	var ipSeekResp struct {
		Location ipLocation
	}
	if err = json.NewDecoder(resp.Body).Decode(&ipSeekResp); err != nil {
		return nil, err
	}
	return &ipSeekResp.Location, nil
}
