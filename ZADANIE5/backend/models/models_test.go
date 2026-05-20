package models

import (
	"testing"
	"time"
)



func TestProduct_PoprawnyProdukt(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 2878.99}
	if err := p.Validate(); err != nil {
		t.Errorf("oczekiwano nil, dostano: %v", err)
	}
}

func TestProduct_PustaNazwa(t *testing.T) {
	p := Product{ID: 1, Name: "", Price: 2878.99}
	if err := p.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustej nazwy")
	}
}

func TestProduct_CenaUjemna(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: -100.0}
	if err := p.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ujemnej ceny")
	}
}

func TestProduct_CenaZero(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 0}
	if err := p.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ceny zero")
	}
}

func TestProduct_IDZero(t *testing.T) {
	p := Product{ID: 0, Name: "PlayStation 5", Price: 2878.99}
	if err := p.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ID zero")
	}
}

func TestProduct_IDUjemne(t *testing.T) {
	p := Product{ID: -1, Name: "PlayStation 5", Price: 2878.99}
	if err := p.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ujemnego ID")
	}
}

func TestProduct_NazwaNiePusta(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 2878.99}
	if p.Name == "" {
		t.Error("nazwa nie powinna być pusta")
	}
}

func TestProduct_CenaWiększaOdZera(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 2878.99}
	if p.Price <= 0 {
		t.Errorf("cena powinna być większa od zera, dostano: %f", p.Price)
	}
}

func TestProduct_IDWiększeOdZera(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 2878.99}
	if p.ID <= 0 {
		t.Errorf("ID powinno być większe od zera, dostano: %d", p.ID)
	}
}

func TestProduct_OpisOpcjonalny(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 2878.99, Description: ""}
	if err := p.Validate(); err != nil {
		t.Errorf("opis jest opcjonalny, nie powinno być błędu: %v", err)
	}
}

func TestProduct_ImageURLOpcjonalny(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 2878.99, ImageURL: ""}
	if err := p.Validate(); err != nil {
		t.Errorf("image_url jest opcjonalny, nie powinno być błędu: %v", err)
	}
}

func TestProduct_BardzoMałaCena(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 0.01}
	if err := p.Validate(); err != nil {
		t.Errorf("cena 0.01 powinna być poprawna: %v", err)
	}
}



func TestOrderItem_PoprawnyItem(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 2, Price: 2878.99}
	if err := item.Validate(); err != nil {
		t.Errorf("oczekiwano nil, dostano: %v", err)
	}
}

func TestOrderItem_PustaNazwa(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "", Quantity: 2, Price: 2878.99}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustej nazwy")
	}
}

func TestOrderItem_CenaUjemna(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 2, Price: -50.0}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ujemnej ceny")
	}
}

func TestOrderItem_CenaZero(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 2, Price: 0}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ceny zero")
	}
}

func TestOrderItem_IlośćZero(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 0, Price: 2878.99}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ilości zero")
	}
}

func TestOrderItem_IlośćUjemna(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: -1, Price: 2878.99}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ujemnej ilości")
	}
}

func TestOrderItem_ProductIDZero(t *testing.T) {
	item := OrderItem{ProductID: 0, Name: "PlayStation 5", Quantity: 1, Price: 2878.99}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ProductID zero")
	}
}

func TestOrderItem_ProductIDUjemne(t *testing.T) {
	item := OrderItem{ProductID: -1, Name: "PlayStation 5", Quantity: 1, Price: 2878.99}
	if err := item.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ujemnego ProductID")
	}
}

func TestOrderItem_ObliczWartość(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 2, Price: 100.0}
	wartosc := item.Price * float64(item.Quantity)
	if wartosc != 200.0 {
		t.Errorf("oczekiwano 200.0, dostano: %f", wartosc)
	}
}

func TestOrderItem_IlośćJeden(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 1, Price: 2878.99}
	if err := item.Validate(); err != nil {
		t.Errorf("ilość 1 powinna być poprawna: %v", err)
	}
}



func TestCustomer_PoprawnyKlient(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1"}
	if err := c.Validate(); err != nil {
		t.Errorf("oczekiwano nil, dostano: %v", err)
	}
}

func TestCustomer_PustaNazwa(t *testing.T) {
	c := Customer{Name: "", Email: "jan@test.pl", Address: "ul. Testowa 1"}
	if err := c.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustej nazwy")
	}
}

func TestCustomer_PustyEmail(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "", Address: "ul. Testowa 1"}
	if err := c.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustego emaila")
	}
}

func TestCustomer_PustyAdres(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@test.pl", Address: ""}
	if err := c.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustego adresu")
	}
}

func TestCustomer_NazwaNiePusta(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1"}
	if c.Name == "" {
		t.Error("nazwa klienta nie powinna być pusta")
	}
}

func TestCustomer_EmailNiePusty(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1"}
	if c.Email == "" {
		t.Error("email klienta nie powinien być pusty")
	}
}

func TestCustomer_AdresNiePusty(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1"}
	if c.Address == "" {
		t.Error("adres klienta nie powinien być pusty")
	}
}

func TestCustomer_WszystkiePustePola(t *testing.T) {
	c := Customer{}
	if err := c.Validate(); err == nil {
		t.Error("oczekiwano błędu dla wszystkich pustych pól")
	}
}


func TestOrder_PoprawneZamówienie(t *testing.T) {
	o := Order{
		ID:    1,
		Total: 99.99,
		Items: []OrderItem{{ProductID: 1, Name: "Test", Quantity: 1, Price: 99.99}},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	if err := o.Validate(); err != nil {
		t.Errorf("oczekiwano nil, dostano: %v", err)
	}
}

func TestOrder_PusteItems(t *testing.T) {
	o := Order{
		Total: 99.99,
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustej listy pozycji")
	}
}

func TestOrder_TotalZero(t *testing.T) {
	o := Order{
		Total: 0,
		Items: []OrderItem{{ProductID: 1, Name: "Test", Quantity: 1, Price: 99.99}},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err == nil {
		t.Error("oczekiwano błędu dla sumy zero")
	}
}

func TestOrder_TotalUjemny(t *testing.T) {
	o := Order{
		Total: -50.0,
		Items: []OrderItem{{ProductID: 1, Name: "Test", Quantity: 1, Price: 99.99}},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err == nil {
		t.Error("oczekiwano błędu dla ujemnej sumy")
	}
}

func TestOrder_PustyKlient(t *testing.T) {
	o := Order{
		Total:    99.99,
		Items:    []OrderItem{{ProductID: 1, Name: "Test", Quantity: 1, Price: 99.99}},
		Customer: Customer{},
	}
	if err := o.Validate(); err == nil {
		t.Error("oczekiwano błędu dla pustego klienta")
	}
}

func TestOrder_KilkaItems(t *testing.T) {
	o := Order{
		Total: 199.98,
		Items: []OrderItem{
			{ProductID: 1, Name: "Test1", Quantity: 1, Price: 99.99},
			{ProductID: 2, Name: "Test2", Quantity: 1, Price: 99.99},
		},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err != nil {
		t.Errorf("oczekiwano nil, dostano: %v", err)
	}
}

func TestOrder_StatusNiePusty(t *testing.T) {
	o := Order{Status: "pending"}
	if o.Status == "" {
		t.Error("status nie powinien być pusty")
	}
}

func TestOrder_CreatedAtNiePusty(t *testing.T) {
	o := Order{CreatedAt: time.Now()}
	if o.CreatedAt.IsZero() {
		t.Error("created_at nie powinien być zerowy")
	}
}



func TestProduct_BardzoWysokaCena(t *testing.T) {
	p := Product{ID: 1, Name: "PlayStation 5", Price: 999999.99}
	if err := p.Validate(); err != nil {
		t.Errorf("bardzo wysoka cena powinna być poprawna: %v", err)
	}
}

func TestProduct_NazwaZeZnakamiSpecjalnymi(t *testing.T) {
	p := Product{ID: 1, Name: "PS5 & Xbox #1!", Price: 2878.99}
	if err := p.Validate(); err != nil {
		t.Errorf("znaki specjalne w nazwie powinny być poprawne: %v", err)
	}
}

func TestProduct_BardzoWielkieID(t *testing.T) {
	p := Product{ID: 999999, Name: "PlayStation 5", Price: 2878.99}
	if err := p.Validate(); err != nil {
		t.Errorf("bardzo duże ID powinno być poprawne: %v", err)
	}
}



func TestOrderItem_BardzoWielkaIlość(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 9999, Price: 2878.99}
	if err := item.Validate(); err != nil {
		t.Errorf("duża ilość powinna być poprawna: %v", err)
	}
}

func TestOrderItem_BardzoMałaCena(t *testing.T) {
	item := OrderItem{ProductID: 1, Name: "PlayStation 5", Quantity: 1, Price: 0.01}
	if err := item.Validate(); err != nil {
		t.Errorf("cena 0.01 powinna być poprawna: %v", err)
	}
}


func TestCustomer_DługaNazwa(t *testing.T) {
	c := Customer{Name: "Jan Aleksander Kowalski-Nowak", Email: "jan@test.pl", Address: "ul. Testowa 1"}
	if err := c.Validate(); err != nil {
		t.Errorf("długa nazwa powinna być poprawna: %v", err)
	}
}


func TestCustomer_DługiAdres(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Bardzo Długa Nazwa Ulicy 123/45, 00-000 Warszawa"}
	if err := c.Validate(); err != nil {
		t.Errorf("długi adres powinien być poprawny: %v", err)
	}
}



func TestCustomer_EmailZSubdomeną(t *testing.T) {
	c := Customer{Name: "Jan Kowalski", Email: "jan@mail.test.pl", Address: "ul. Testowa 1"}
	if err := c.Validate(); err != nil {
		t.Errorf("email z subdomeną powinien być poprawny: %v", err)
	}
}



func TestOrder_JedenItem(t *testing.T) {
	o := Order{
		Total: 99.99,
		Items: []OrderItem{{ProductID: 1, Name: "Test", Quantity: 1, Price: 99.99}},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err != nil {
		t.Errorf("zamówienie z jednym itemem powinno być poprawne: %v", err)
	}
}

func TestOrder_WieleItems(t *testing.T) {
	o := Order{
		Total: 499.95,
		Items: []OrderItem{
			{ProductID: 1, Name: "Test1", Quantity: 1, Price: 99.99},
			{ProductID: 2, Name: "Test2", Quantity: 1, Price: 99.99},
			{ProductID: 3, Name: "Test3", Quantity: 1, Price: 99.99},
			{ProductID: 4, Name: "Test4", Quantity: 1, Price: 99.99},
			{ProductID: 5, Name: "Test5", Quantity: 1, Price: 99.99},
		},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err != nil {
		t.Errorf("zamówienie z wieloma itemami powinno być poprawne: %v", err)
	}
}

func TestOrder_BardzoWysokiTotal(t *testing.T) {
	o := Order{
		Total: 999999.99,
		Items: []OrderItem{{ProductID: 1, Name: "Test", Quantity: 1, Price: 999999.99}},
		Customer: Customer{
			Name: "Jan Kowalski", Email: "jan@test.pl", Address: "ul. Testowa 1",
		},
	}
	if err := o.Validate(); err != nil {
		t.Errorf("bardzo wysoki total powinien być poprawny: %v", err)
	}
}

func TestOrder_MaID(t *testing.T) {
	o := Order{ID: 1, Total: 99.99}
	if o.ID <= 0 {
		t.Errorf("ID powinno być większe od zera, dostano: %d", o.ID)
	}
}