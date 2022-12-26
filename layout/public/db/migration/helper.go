package migration

import "gorm.io/gorm"

type Helper struct {
	m   gorm.Migrator
	err error
}

func With(db *gorm.DB) *Helper {
	return &Helper{m: db.Migrator()}
}

func (h *Helper) Error() error {
	return h.err
}

func (h *Helper) CreateTable(tables ...interface{}) *Helper {
	if h.err != nil {
		return h
	}
	for _, v := range tables {
		if err := h.m.CreateTable(v); err != nil {
			h.err = err
			break
		}
	}
	return h
}

func (h *Helper) DropTable(tables ...interface{}) *Helper {
	if h.err != nil {
		return h
	}
	for _, v := range tables {
		if err := h.m.DropTable(v); err != nil {
			h.err = err
			break
		}
	}
	return h
}

func (h *Helper) AddColumn(table interface{}, columns ...string) *Helper {
	if h.err != nil {
		return h
	}
	for _, col := range columns {
		if err := h.m.AddColumn(table, col); err != nil {
			h.err = err
			break
		}
	}
	return h
}

func (h *Helper) DropColumn(table interface{}, columns ...string) *Helper {
	if h.err != nil {
		return h
	}
	for _, col := range columns {
		if err := h.m.DropColumn(table, col); err != nil {
			h.err = err
			break
		}
	}
	return h
}
