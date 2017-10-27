package loopers

import (
	"image"
	"log"
	"testing"

	"github.com/jmigpin/editor/drawutil2"
	"golang.org/x/image/math/fixed"
)

var loremStr2 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

var testStr3 = "a\na\na\na\na\na\na\na\na\na\na\na\na\na\n"

var testStr4 = "aaaaa\nbbbbb\nccccc\n"
var testStr5 = `
aaaaa
	bbbbb
	ccccc
`

var testStr6 = "abcde abcde abcde abcde abcde"
var testStr7 = `
abcde
abcde
abcde
abcde
abcde
abcde
abcde
abcde
abcde
abcde
abcde`

func TestPosData1(t *testing.T) {
	f1 := drawutil2.GetTestFace()
	f2 := drawutil2.NewFaceRunes(f1)
	f3 := drawutil2.NewFaceCache(f2)
	face := f3

	bounds := image.Rect(0, 0, 1000, 1000)
	max := fixed.P(bounds.Dx(), bounds.Dy())

	start := &EmbedLooper{}
	strl := NewStringLooper(face, testStr7)
	linel := NewLineLooper(strl, max.Y)
	wlinel := NewWrapLine2Looper(strl, linel, max.X)
	pdl := NewPosDataLooper()
	pdl.Strl = strl
	pdl.Keepers = []PosDataKeeper{strl, wlinel}
	pdl.Jump = 5

	pdl.SetOuterLooper(start)
	strl.SetOuterLooper(pdl)
	linel.SetOuterLooper(strl)
	wlinel.SetOuterLooper(linel)

	wlinel.Loop(func() bool { return true })

	log.Printf("pdl has %v points", len(pdl.data))

	log.Printf("ri %v", strl.Ri)

	//for _, d := range pdl.data {
	//log.Printf("pdl data ri %v %v", d.ri, d.penBoundsMaxY)
	//}

	p := fixed.P(10, 0)
	//pdl.RestorePosDataCloseToPoint(&p)
	pd, ok := pdl.PosDataCloseToPoint(&p)
	if ok {
		log.Printf("restoring %+v", pd)
		pdl.restore(pd)
	}
	i := pdl.GetIndex(&p, wlinel)

	log.Printf("ri %v", strl.Ri)
	log.Printf("i %v", i)

	//------

	//// update

	//bounds = image.Rect(0, 0, 200, 20)
	//max = fixed.P(bounds.Dx(), bounds.Dy())

	//wlinel.MaxX = max.X
	//pdl.Update()

	//log.Printf("pen max %v", max)
	//log.Printf("pen %v", strl.Pen)
	//for _, d := range pdl.data {
	//	log.Printf("%v", spew.Sdump(d))
	//}
}