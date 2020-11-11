package xor

import (
	"container/heap"
	"matasano/cryptopals/util"
	"math"
	"unicode"
)

func Crack(ciphertext []byte) [][]byte {
	keys := make([][]byte, 5)
	sizes := guessKeySizes(ciphertext, 5)
	for index, size := range sizes {
		possibleKeyword := make([]byte, size.Length)
		blocks := util.Blockify(ciphertext, size.Length)
		for i := 0; i < size.Length; i++ {
			possibleKeyword[i] = crackKey(blocks[i])
		}
		keys[index] = possibleKeyword
	}
	return keys
}

func crackKey(ciphertext []byte) byte {
	type crack struct {
		Text  string
		Key   byte
		Score float64
	}
	var best crack
	for key := 0; key < 256; key++ {
		text := caesarCipher(ciphertext, byte(key))
		if score := scorePlaintext(text); score > best.Score {
			best = crack{string(text), byte(key), score}
		}
	}
	return best.Key
}

func scorePlaintext(plaintext []byte) float64 {
	characterFrequencies := map[rune]float64{
		'a': .08167, 'b': .01492, 'c': .02782, 'd': .04253,
		'e': .12702, 'f': .02228, 'g': .02015, 'h': .06094,
		'i': .06094, 'j': .00153, 'k': .00772, 'l': .04025,
		'm': .02406, 'n': .06749, 'o': .07507, 'p': .01929,
		'q': .00095, 'r': .05987, 's': .06327, 't': .09056,
		'u': .02758, 'v': .00978, 'w': .02360, 'x': .00150,
		'y': .01974, 'z': .00074, ' ': .13000,
	}
	var sum float64
	for _, character := range plaintext {
		sum += characterFrequencies[unicode.ToLower(rune(character))]
	}
	return sum
}

type keySize struct {
	Length int
	Score  float64
}

type keySizeHeap []keySize

func (h keySizeHeap) Len() int            { return len(h) }
func (h keySizeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h keySizeHeap) Less(i, j int) bool  { return h[i].Score < h[j].Score }
func (h *keySizeHeap) Push(x interface{}) { *h = append(*h, x.(keySize)) }
func (h *keySizeHeap) Pop() interface{} {
	temp := *h
	top := len(temp)
	x := temp[top-1]
	*h = temp[:top-1]
	return x
}

func guessKeySizes(ciphertext []byte, head int) keySizeHeap {
	minNormalHeap := keySizeHeap{}
	heap.Init(&minNormalHeap)
	for size := 2; size < 42; size++ {
		chunks := util.Chunkify(ciphertext, size)
		if len(chunks) < 2 {
			break
		}
		heap.Push(
			&minNormalHeap,
			keySize{size, averageDistanceOfChunks(chunks) / float64(size)},
		)
	}
	bestGuesses := make(keySizeHeap, head)
	for i := 0; i < head; i++ {
		bestGuesses[i] = heap.Pop(&minNormalHeap).(keySize)
	}
	return bestGuesses
}

func averageDistanceOfChunks(chunks [][]byte) float64 {
	var sum float64
	var length float64
	for index := 0; index < len(chunks)-1; index++ {
		diff := distance(chunks[index], chunks[index+1])
		if diff == -1 {
			break
		}
		sum += float64(diff)
		length += 1.0
	}
	averageDistance := sum / length
	if math.IsNaN(averageDistance) {
		return math.MaxFloat64
	}
	return averageDistance
}

func distance(left, right []byte) int {
	if len(left) != len(right) {
		return -1
	}
	distance := 0
	for index := range left {
		diff := left[index] ^ right[index]
		for bit := 0; bit < 8; bit++ {
			distance += int((diff >> bit) & 1)
		}
	}
	return distance
}
