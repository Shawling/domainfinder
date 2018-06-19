package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

//解析json 时，字段需要是首字母大写，这样子 json 库才能正确获取到这些字段
type synonymsSearchReasult struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(para string) ([]string, error) {
	var syns []string

	resp, err := http.Get("http://words.bighugelabs.com/api/2/" +
		b.APIKey + "/" + para + "/json")
	defer resp.Body.Close()
	if err != nil {
		return syns, fmt.Errorf("bighuge: Failed when looking for synonyms for %s: %s", para, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return syns, fmt.Errorf("bighuge: Failed with status: %s", resp.Status)
	}
	var data synonymsSearchReasult
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}
