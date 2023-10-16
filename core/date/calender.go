package date

import (
	"fmt"
	"time"
)

func YearAndQuarter(t time.Time) string {
	return fmt.Sprintf("%d%d", Year(t), Quarter(t))
}
