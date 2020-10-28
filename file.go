package gob

import (
	"fmt"

	"github.com/oxtoacart/god"
)

func CallMe() string {
	return fmt.Sprintf("b:%s", god.CallMe())
}
