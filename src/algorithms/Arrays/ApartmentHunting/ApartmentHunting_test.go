package apartmenthunting

import (
	"fmt"
	"testing"
)

// go test -v -run TestApartmentHunting ./.. -count=1
func TestApartmentHunting(t *testing.T) {
	b := []Block{
		Block{"gym": false, "school": true, "store": false},
		Block{"gym": true, "school": false, "store": false},
		Block{"gym": true, "school": true, "store": false},
		Block{"gym": false, "school": true, "store": false},
		Block{"gym": false, "school": true, "store": true},
	}

	reqs := []string{"gym", "school", "store"}

	reqStatus := map[string]int{"gym": 0, "school": 0, "store": 0}

	// res := ApartmentHunting(b, reqs)
	res := updateReqStatus(Block{"gym": true, "school": false, "store": true}, &reqStatus, true, reqs)
	res2 := updateReqStatus(Block{"gym": true, "school": false, "store": true}, &reqStatus, false, reqs)
	res2 = updateReqStatus(Block{"gym": true, "school": true, "store": true}, &reqStatus, false, reqs)
	fmt.Println("Res: ", res, b)
	fmt.Println("Res: ", res2, b)
	fmt.Println("********************")
	fmt.Println(ApartmentHunting(b, reqs))
}
