package roman

//known Roman Numerals symbols
var symbols = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

//define a new type ~> Roman:-)
type Roman struct{}

//used create a variable of type Roman
func NewRoman() *Roman {
	return &Roman{}
}

//converts a roman numeral to a decimal number
//s ~> the roman numeral to convert
func (r *Roman) ToDecimal(s string) int {
	var result int // will hold the decimal equivalent
	size := len(s)
	for i := 0; i < size; i++ {
		c := string(s[i]) // the current character in s
		current := symbols[c]
		if i < size-1 {
			n := string(s[i+1]) // the next character in s
			next := symbols[n]
			if current < next {
				result += next - current
				i++
			} else {
				result += current
			}
		} else {
			result += current
		}
	}
	return result
}
