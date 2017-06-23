package main

type Cooker interface {
	MakeSushi()
	MakeNimono()
	StirFry()
}

func CookFish(ingredient interface{}) {
	o := &cook.Oliver{}
	switch ingredient.(type) {
	case fish.Mackerel:
		o.MakeSushi()
	case fish.Tuna:
		o.MakeNimono()
	default:
		o.StirFry()
	}
}
func main() {

}
