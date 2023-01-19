package middleware

import (
	"net/http"
	"time"
	
)
type rate struct {
    limit     int
    interval  int
    count     int
    last      time.Time
}
var RateLimiter = func(limit int, interval int) func(http.HandlerFunc) http.HandlerFunc {
    //store the number of requests and the last request time
    var store = make(map[string]*rate)
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            clientIP := r.RemoteAddr
            if store[clientIP] == nil {
                store[clientIP] = &rate{
                    limit:  limit,
                    interval: interval,
                    count:  1,
                    last:   time.Now(),
                }
            } else {
                now := time.Now()
                diff := now.Sub(store[clientIP].last)
                if int(diff.Seconds()) < store[clientIP].interval {
                    store[clientIP].count++
                } else {
                    store[clientIP].count = 1
                }
                store[clientIP].last = now
                if store[clientIP].count > store[clientIP].limit {
                    http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
                    return
                }
            }
            next(w, r)
        }
    }
}
