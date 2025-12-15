package stealth
import "github.com/go-rod/rod/lib/proto"


import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

type Point struct {
	X float64
	Y float64
}

// MoveMouseBezier moves the mouse in a human-like curved path
func MoveMouseBezier(page *rod.Page, from, to Point) {
	cp1 := randomControlPoint(from, to)
	cp2 := randomControlPoint(from, to)

	steps := rand.Intn(25) + 30 // 30–55 steps
	duration := time.Duration(rand.Intn(400)+300) * time.Millisecond

	for i := 0; i <= steps; i++ {
    	t := float64(i) / float64(steps)
    	p := cubicBezier(from, cp1, cp2, to, t)

    	_ = page.Mouse.MoveTo(proto.Point{
    		X: p.X,
    		Y: p.Y,
    	})

    	time.Sleep(duration / time.Duration(steps))
    }


}

func cubicBezier(p0, p1, p2, p3 Point, t float64) Point {
	u := 1 - t
	tt := t * t
	uu := u * u
	uuu := uu * u
	ttt := tt * t

	x := uuu*p0.X +
		3*uu*t*p1.X +
		3*u*tt*p2.X +
		ttt*p3.X

	y := uuu*p0.Y +
		3*uu*t*p1.Y +
		3*u*tt*p2.Y +
		ttt*p3.Y

	// micro jitter
	x += rand.Float64()*1.5 - 0.75
	y += rand.Float64()*1.5 - 0.75

	return Point{X: x, Y: y}
}

func randomControlPoint(a, b Point) Point {
	dx := b.X - a.X
	dy := b.Y - a.Y

	dist := math.Hypot(dx, dy)
	offset := dist * (rand.Float64()*0.3 + 0.2)

	angle := math.Atan2(dy, dx) + (rand.Float64()-0.5)
	return Point{
		X: a.X + math.Cos(angle)*offset,
		Y: a.Y + math.Sin(angle)*offset,
	}
}
