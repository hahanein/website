package tarpit

import (
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	sentences := make([]string, 0, 6)

	for range 6 {
		x := words[rng.Intn(len(words))]
		y := words[rng.Intn(len(words))]
		z := words[rng.Intn(len(words))]
		sentence := fmt.Sprintf("The %s of %s was %s.", x, y, z)
		sentences = append(sentences, sentence)
	}

	content := strings.Join(sentences, " ")
	return content
}

func generateDeterministicPath(rng *rand.Rand) string {
	pathLen := rng.Intn(8) + 1
	pathParts := make([]string, pathLen)

	for i := range pathLen {
		pathParts[i] = words[rng.Intn(len(words))]
	}

	return "/" + strings.Join(pathParts, "-")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("If-None-Match") != "" {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	time.Sleep(4 * time.Second)

	seed := hashPath(r.URL.Path)
	content := generateDeterministicContent(seed)

	rng1 := rand.New(rand.NewSource(int64(seed) + 1))
	link1 := generateDeterministicPath(rng1)

	rng2 := rand.New(rand.NewSource(int64(seed) + 2))
	link2 := generateDeterministicPath(rng2)

	response := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<title>%s</title>
</head>
<body>
<p>%s</p>
<ul>
	<li><a href="%s">%s</a></li>
	<li><a href="%s">%s</a></li>
</ul>
</body>
</html>
	`,
		content, content, link1, link1, link2, link2)

	h := sha256.New()
	h.Write([]byte(response))
	etag := fmt.Sprintf("\"%x\"", h.Sum(nil))

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("ETag", etag)

	if _, err := fmt.Fprint(w, response); err != nil {
		log.Printf("handling request failed: %v", err)
	}
}
