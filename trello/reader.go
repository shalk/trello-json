package trello

import (
	"encoding/json"
	"io/ioutil"
)

type myCard struct {
	id     string `trello:"id"`
	idList string `trello:"idList"`
	name   string `trello:"name"`
}
type myList struct {
	id   string `trello:"id"`
	name string `trello:"name"`
}

func Read(filename string) map[string][]string {
	info := make(map[string][]string)
	byt, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	if err = json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	cards := dat["cards"].([]interface{})
	lists := dat["lists"].([]interface{})

	for _, list := range lists {
		listMap := list.(map[string]interface{})
		info[listMap["id"].(string)] = make([]string, 0)
		info[listMap["id"].(string)] = append(info[listMap["id"].(string)], listMap["name"].(string))
	}

	for _, card := range cards {
		map1 := card.(map[string]interface{})
		card1 := &myCard{map1["id"].(string), map1["idList"].(string), map1["name"].(string)}
		v, ok := info[card1.idList]
		if ok {
			v = append(v, card1.name)
			info[card1.idList] = v
		} else {
			info[card1.idList] = make([]string, 0)
		}
	}
	return info
}


