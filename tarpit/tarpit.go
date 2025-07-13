package tarpit

import (
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

var words = []string{
	"wind", "time", "road", "change",
	"train", "love", "death", "freedom",
	"truth", "lies", "soul", "dream",
	"night", "light", "darkness", "pain",
	"song", "voice", "river", "sky",
	"fire", "rain", "eye", "hand",
	"heart", "mind", "stone", "blood",
	"flesh", "bone", "door", "wheel",
	"dance", "mask", "gun", "cross",
	"tongue", "shadow", "mirror", "dust",
	"chain", "grave", "cage", "bell",
	"storm", "sin", "grace", "angel",
}

func hashPath(path string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(path))
	return h.Sum64()
}

func generateDeterministicContent(seed uint64) string {
	rng := rand.New(rand.NewSource(int64(seed)))
	selectedWords := make([]string, 0, 16)
	for range 16 {
		selectedWords = append(selectedWords, words[rng.Intn(len(words))])
	}

	content := strings.Join(selectedWords, " ")
	return content
}

func generateDeterministicPath(rng *rand.Rand) string {
	pathLen := rng.Intn(8) + 1
	pathParts := make([]string, pathLen)

	for i := range pathLen {
		pathParts[i] = words[rng.Intn(len(words))]
	}

	return "/" + strings.Join(pathParts, "/")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	seed := hashPath(r.URL.Path)
	content := generateDeterministicContent(seed)

	rng1 := rand.New(rand.NewSource(int64(seed) + 1))
	link1 := generateDeterministicPath(rng1)
	escapedLink1 := url.QueryEscape(link1)

	rng2 := rand.New(rand.NewSource(int64(seed) + 2))
	link2 := generateDeterministicPath(rng2)
	escapedLink2 := url.QueryEscape(link2)

	response := fmt.Sprintf("%s\n\n<a href=\"%s\">%s</a>\n<a href=\"%s\">%s</a>\n",
		content, link1, escapedLink1, link2, escapedLink2)

	h := sha256.New()
	h.Write([]byte(response))
	etag := fmt.Sprintf("\"%x\"", h.Sum(nil))

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("ETag", etag)

	if _, err := fmt.Fprint(w, response); err != nil {
		log.Printf("handling request failed: %v", err)
	}
}
