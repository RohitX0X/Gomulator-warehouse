// eventhub_connect.go

package src

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const operatorstring string = "operator"

// PushToEventHub sends data to an Azure Event Hub.
func transformData(id string, event_map *map[string]*[]byte, aseq []int) error {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf(string(r1.Int31n(10)))

	for i := 0; i < len(aseq); i++ {
		v := *event_map
		iter := strconv.Itoa(i)
		data := v[string(iter)]
		data_str := string(*data)

		go PushToEventHub(data_str)
	}
	return nil
}
