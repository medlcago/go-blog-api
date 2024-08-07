package pagination

const (
	defaultLimit  = 20
	maxLimit      = 100
	defaultOffset = 0
	maxOffset     = 10000
)

type LimitOffsetPaginator interface {
	GetLimit() int
	GetOffset() int
}

type LimitOffset struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (l *LimitOffset) SetDefault() {
	l.Limit = defaultLimit
	l.Offset = defaultOffset
}

func (l *LimitOffset) GetLimit() int {
	if l.Limit <= 0 || l.Limit > maxLimit {
		return defaultLimit
	}
	return l.Limit
}

func (l *LimitOffset) GetOffset() int {
	if l.Offset < 0 {
		return defaultOffset
	} else if l.Offset > maxOffset {
		return maxOffset
	}
	return l.Offset
}
