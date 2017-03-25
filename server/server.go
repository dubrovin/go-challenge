package server

import (
	"time"
	"net/http"
	"log"
	"encoding/json"
	"sync"
	"io/ioutil"
	"net/url"
	"sort"
)

type Number struct {
	Numbers []int `json:"Numbers"`
}

type Server struct {
	ListenAddr string
	Client     *http.Client
	Numbers    []int
	wg         *sync.WaitGroup
	mu         *sync.Mutex
}

func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func NewServer(listenAddr string, timeout time.Duration) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Client: &http.Client{Timeout:timeout},
		Numbers: make([]int, 0),
		wg: &sync.WaitGroup{},
		mu: &sync.Mutex{},
	}
}

func (s *Server) Run() {
	http.HandleFunc("/numbers", serverHandler(s)) // set router
	err := http.ListenAndServe(s.ListenAddr, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (s *Server) parse(parsedURL string) {
	s.wg.Add(1)
	go func(parsedURL string) {
		defer s.wg.Done()
		resp, err := s.Client.Get(parsedURL)
		if err != nil {
			log.Println("Get error", err)
			return
		}
		defer resp.Body.Close()

		var num Number
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &num)

		if err != nil {
			log.Println("Unmarshal error", err)
			return
		}

		s.mu.Lock()
		s.Numbers = append(s.Numbers, num.Numbers...)
		s.mu.Unlock()

	}(parsedURL)
}

func serverHandler(s *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for _, ur := range r.Form["u"] {
			//check for valid url
			parsedUrl, err := url.Parse(ur)
			if err != nil {
				log.Println("URL parse error", err)
				return
			}
			//parallel parse
			s.parse(parsedUrl.String())

		}
		//wait for all urls will be parsed
		s.wg.Wait()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		s.mu.Lock()
		s.Numbers = removeDuplicates(s.Numbers)
		sort.Ints(s.Numbers)
		s.mu.Unlock()

		numbers := map[string][]int{"Numbers": s.Numbers}
		json.NewEncoder(w).Encode(numbers)

	}
}