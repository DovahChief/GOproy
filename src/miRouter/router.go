package miRouter

import (
	_ "net/url"
	_ "path"
	_ "regexp" //exp regular para el router
	"sync"
	_ "sync/atomic"
)

// MuxServ publica
type MuxServ struct {
	mu   sync.RWMutex
	m    map[string]muxEntry
	host bool //dice si el patron contiene al host
}

func matchPath(patron, path string) bool {
	if len(patron) == 0 {
		return false
	}

	n := len(patron)

	if patron[n-1] != '/' {
		match, _ := regexp.MatchString(patron, string(path[0:n]))
		return match
	}
	fullMatch, _ := regexp.MatchString(patron, string(path[0:n]))
	return len(path) >= n && fullMatch
}

func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	var n = 0
	for k, v := range mux.m {
		if !matchPath(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
		}
	}
	return
}
