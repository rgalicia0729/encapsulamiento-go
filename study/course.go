package study

type course struct {
	name    string
	price   float64
	isFree  bool
	classes map[uint]string
}

func NewCourse(name string) *course {
	return &course{
		name:   name,
		isFree: true,
	}
}

func (c *course) Name() string {
	return c.name
}

func (c *course) Price() float64 {
	return c.price
}

func (c *course) IsFree() bool {
	return c.isFree
}

func (c *course) Classes() map[uint]string {
	return c.classes
}

func (c *course) SetPrice(price float64) {
	c.price = price

	if price == 0 {
		c.isFree = true
	} else {
		c.isFree = false
	}
}

func (c *course) SetClasses(classes []string) {
	newClasses := make(map[uint]string)

	for i, class := range classes {
		newClasses[uint(i+1)] = class
	}

	c.classes = newClasses
}
