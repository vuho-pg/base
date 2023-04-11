package pagination

import (
	"gorm.io/gorm"
	"strings"
)

const DefaultLimit = 10
const DefaultPage = 1

type order struct {
	Name  string
	Order string
}

type Applier interface {
	Apply(db *gorm.DB) *gorm.DB
	SetTotal(total int64)
}

type Page struct {
	Limit     int    `json:"limit"`
	Total     int    `json:"total"`
	Page      int    `json:"page"`
	TotalPage int    `json:"total_page"`
	OrderBy   string `json:"order_by"`
	orders    []order
}

func (p *Page) Correct() {
	if p.Page < DefaultPage {
		p.Page = DefaultPage
	}
	if p.Limit <= 0 {
		p.Limit = DefaultLimit
	}
	if p.OrderBy != "" {
		orders := strings.Split(p.OrderBy, ",")
		for _, orderString := range orders {
			if orderString == "" {
				continue
			}
			order := order{}
			if strings.HasPrefix(orderString, "-") {
				order.Name = strings.TrimPrefix(orderString, "-")
				order.Order = "DESC"
			} else {
				order.Name = orderString
				order.Order = "ASC"
			}
			p.orders = append(p.orders, order)
		}
	}
}

func (p *Page) Apply(db *gorm.DB) *gorm.DB {
	p.Correct()
	db = db.Limit(p.Limit).Offset((p.Page - 1) * p.Limit)
	for _, order := range p.orders {
		db = db.Order(order.Name + " " + order.Order)
	}
	return db
}

func (p *Page) SetTotal(total int64) {
	p.Correct()
	p.Total = int(total)
	p.TotalPage = p.Total / p.Limit
	if p.Total%p.Limit != 0 {
		p.TotalPage++
	}
}
