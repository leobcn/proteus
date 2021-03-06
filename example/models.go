package example

import "time"

//go:generate proteus -p github.com/src-d/proteus/example -f $GOPATH/src/github.com/src-d/proteus/example/protos

//proteus:generate
type Product struct {
	Model

	Name  string
	Price float64

	Tags       Tags
	CategoryID int64
	// Category will not be generated because we explicitly said so.
	Category Category `proteus:"-"`
}

//proteus:generate
type Category struct {
	Model

	Name    string
	Type    Type
	Color   Color
	Options CategoryOptions
}

type Tags []string

// Type will be transformed into an enum.
//proteus:generate
type Type byte

const (
	Public Type = iota
	Private
	Custom
)

// Color does not have proteus:generate in a comment, so all fields of
// type Color will be treated as just string, not as an enum.
type Color string

const (
	Blue   Color = "blue"
	Red    Color = "red"
	Yellow Color = "yellow"
)

// CategoryOptions is not marked for generations, but it is used in another
// structs, so it will be generated because of it.
type CategoryOptions struct {
	ShowPrices bool
	CanBuy     bool
}

// Model is not marked for generation, so it won't be generated.
type Model struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// User will not be generated.
type User struct {
	Model

	Username string
	Password string
	Email    string
}

type MyTime struct {
	Time time.Time
	Name string
}

//proteus:generate
func RandomNumber(mean, std float64) float64 {
	// Related documentation: https://xkcd.com/221/
	return 4*std + mean // 4 was chosen using the XKCD RNG
}

//proteus:generate
func RandomCategory() CategoryOptions {
	return CategoryOptions{
		ShowPrices: RandomBool(),
		CanBuy:     RandomBool(),
	}
}

//proteus:generate
func GetAlphaTime() MyTime {
	return MyTime{Time: time.Unix(0, 0), Name: "alpha"}
}

//proteus:generate
func GetOmegaTime() (*MyTime, error) {
	t, err := time.Parse("Jan 2, 2006 at 3:04pm", "Dec 12, 2012 at 10:30am")
	if err != nil {
		return nil, err
	}

	return &MyTime{Time: t, Name: "omega"}, nil
}

//proteus:generate
func GetPhone() *Product {
	return &Product{
		Name:       "MiPhone",
		Price:      733.33,
		Tags:       Tags{"cool", "mi", "phone"},
		CategoryID: 1,
		Category:   Category{},
	}
}

func RandomBool() bool {
	return true // Truly random. Selected by flipping a coin... once.
}
