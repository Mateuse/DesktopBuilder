package constants

const (
	ALL_COLUMNS       = "*"
	COMPONENTS_TABLE  = "components"
	DEFAULT_PAGE_SIZE = 50
)

var (
	COMPONENTS_SELECT_COLUMNS = []string{"id", "category", "brand", "model", "sku", "upc", "specs", "created_at"}
)

type LimitAndOffset struct {
	Limit  int
	Offset int
}
