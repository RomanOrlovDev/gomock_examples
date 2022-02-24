package item

//go:generate mockgen -source=item.go -destination=./item_mock.go -package=item

type Isnterface interface {
	Elements() []FacadeItem
}

type FacadeItem interface {
	Id() int
	GetName() string
	GetNameWithValidation() (string, error)
}

type Items struct {
	elements []FacadeItem
}

type AnotherItems struct {
	elements []AnotherItem
}

type AnotherItem struct {
	id   int
	name string
}

func NewAnotherItem(id int, name string) *AnotherItem {
	return &AnotherItem{
		id:   id,
		name: name,
	}
}

func (a AnotherItems) AnotherElements() []AnotherItem {
	return a.elements
}

func (a *AnotherItems) AddAnotherItem(ai *AnotherItem) {
	a.elements = append(a.elements, *ai)
	t := 1
	_ = t
}

func NewAnotherItems() AnotherItems {
	return AnotherItems{
		elements: []AnotherItem{},
	}
}

type Item struct {
	id   int
	name string
}

func NewItems() *Items {
	return &Items{
		elements: []FacadeItem{},
	}
}

func NewItem(params ...interface{}) *Item {
	return &Item{
		id:   params[0].(int),
		name: params[1].(string),
	}
}

func (is *Items) Elements() []FacadeItem {
	return is.elements
}

func (i *Item) Id() int {
	return i.id
}

func (i *Item) GetName() string {
	return i.name
}

func (is *Items) Add(i FacadeItem) {
	is.elements = append(is.elements, i)
}

func (i *Item) GetNameWithValidation() (string, error) {
	return "", nil
	//return "", fmt.Errorf("validation error occurred")
}
