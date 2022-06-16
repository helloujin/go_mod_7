package eletronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	Phone
	ButtonsCount() int
}

type Smartphone interface {
	Phone
	OS() string //  Phone operation system
}

type applePhone struct {
	brand     string
	model     string
	typephone string
	os        string
}

type androidPhone struct {
	brand     string
	model     string
	typephone string
	os        string
}

type radioPhone struct {
	brand        string
	model        string
	typephone    string
	buttonscount int
}

// реализация для Айфона
func NewApplePhone(model string) applePhone {
	return applePhone{brand: "Apple", model: model, typephone: "smartphone", os: "IOS"}
}

func (a applePhone) Brand() string {
	return a.brand
}

func (a applePhone) Model() string {
	return a.model
}

func (a applePhone) Type() string {
	return a.typephone
}

func (a applePhone) OS() string {
	return a.os
}

// Реализация для телефонов на Андройде
func NewAndroidPhone(brand, model string) androidPhone {
	return androidPhone{brand: brand, model: model, typephone: "smartphone", os: "Android"}
}
func (a androidPhone) Brand() string {
	return a.brand
}

func (a androidPhone) Model() string {
	return a.model
}

func (a androidPhone) Type() string {
	return a.typephone
}

func (a androidPhone) OS() string {
	return a.os
}

//реализация для стационара

func NewStantionPhone(brand, model string, buttonscount int) radioPhone {
	return radioPhone{brand: brand, model: model, typephone: "station", buttonscount: buttonscount}
}

func (s radioPhone) Brand() string {
	return s.brand
}

func (s radioPhone) Model() string {
	return s.model
}

func (s radioPhone) Type() string {
	return s.typephone
}

func (s radioPhone) ButtonsCount() int {
	return s.buttonscount
}
