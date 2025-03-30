package cell

import (
	"kwdb/internal/helper"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var storage map[string]Cell
var randoms [1_000_000]string
var bs []byte

func BenchmarkPrepare(b *testing.B) {

	bs = []byte(generateRandomString())

	for i := 0; i < 1000000; i++ {
		rs := generateRandomString()
		randoms[i] = rs
	}

	b.Log(len(randoms))
	b.Log(helper.MemStatInfo())
}
func BenchmarkCellMem(b *testing.B) {
	storage = make(map[string]Cell)

	b.Log("s-" + helper.MemStatInfo())

	for i := 0; i < 1000000; i++ {
		storage[strconv.Itoa(i)] = Cell{84600, generateRandomString()}
	}

	b.Log("e-" + helper.MemStatInfo())
	storage = nil
}

func generateRandomString() string {
	// Набор символов для генерации
	return "лцыова жох012 гухлфыовд ьалыьа х9г120=397(*?:(?*?(*) одувыождл ожфыдвло жфдылов ждфлыо вждлохз29г312зл3уь фьовжшопждфрнывжд оырпжыв ажыовражщхйщцу фьывт жфрехщйфшыво жфдлтапзйщшцрух шфыжвдфт ыжр2х3812хушо ждьфытвждфолры вхжшо1х 2шувржд ыфтвж1л2хуо218унрдлфыовж  1293г=091 ожйдлво жэлоувжэлох023г1о 0хоыжвлфоыжвлов х1шгув0 х9гхзфловжлфовжэл хэ12093гдлпорждьыоаждлгйзйоудлао .я"
	characters := "abcdefghi s;ldfkj ;sdlkfj ;sld oiy12309- 8y019u23 hj; sdhnf;l yh-0128 3y[0dfjdfiyh 012epfdj fk  hf[u1023 yeednf kf j;sldk jf;sdlkf j;sdlkf j;saldkf js;aldkf j;saldkf;saldkfj ;alsdkf jal;skdjf ;saldkfj jklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST ;lsjdf ksdjf ;kashdjf ;;dfa shd;fkasj d;flkasdj f;lsakdj f;lsadkjf ;alsdkf j;sadlkf j;sadlkfj ;saldkfj ;saldkfUVWXYZ0123456789"
	return characters
}

func randILen() int {
	rand.Seed(time.Now().UnixNano())

	// Генерация числа от 10 до 1000
	return rand.Intn(991) + 10 // 991 = 1000 - 10 + 1
}
