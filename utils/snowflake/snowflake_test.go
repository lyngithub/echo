package snowflake

import (
	"fmt"
	"testing"
)

func TestGetResp(t *testing.T) {

	for i := 0; i < 101; i++ {
		fmt.Println(GetResp())
	}
}
