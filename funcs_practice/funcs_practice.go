package funcs_practice

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

type option func(*Item)

func NewItem(opts ...option) *Item {
	// инициализируем типовыми значениями
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}
	// применяем опции в том порядке, в котором они были заявлены
	for _, opt := range opts {
		opt(i)
	}
	return i
}

func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}

func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func Area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	}
	return nil, false
}
