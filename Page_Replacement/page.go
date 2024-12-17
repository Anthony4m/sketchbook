package Page_Replacement

import (
	"github.com/google/uuid"
)

type Page struct {
	id        string
	value     string
	reference int
}

func NewPage(value string) *Page {
	return &Page{
		id:        uuid.New().String(),
		value:     value,
		reference: 1,
	}
}

func (p *Page) GetValue(id string) string {
	if id == p.id {
		return p.value
	}
	return ""
}

func (p *Page) GetValueWithPage(page Page) string {
	return page.value
}
