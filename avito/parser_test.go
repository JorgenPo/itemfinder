package avito

import (
	"gotest.tools/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	var parser Parser

	item, err := parser.Parse("https://www.avito.ru/sankt-peterburg/knigi_i_zhurnaly/kompyuternye_seti_1554383385")

	if err != nil {
		t.Errorf("failed to parse first request: %v", err)
	}

	assert.Equal(t, item.Title, "Компьютерные сети")
	assert.Equal(t, item.Description, "Книга Э. Таненбаума, в отличном состоянии.4-е издание 2011 года.")
	assert.Equal(t, item.PhoneNumber, "8 931 XXX-XX-XX")
	assert.Equal(t, item.User.Name, "Лилит")
	assert.Equal(t, item.Address, "метро , 2 линия, Санкт-Петербург, Звёздная")
	assert.Equal(t, item.Price, 1300)
}