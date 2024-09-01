package prober

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	maxn int = 2048
)

type (
	Node struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		Stat int    `json:"stat"`
	}

	Pipe struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	Prober struct {
		Name  string `json:"name"`
		Nodes []Node `json:"nodes"`
		Pipes []Pipe `json:"pipes"`

		Graph [maxn][]int
		Nodid map[string]int
	}
)

var (
	mark [maxn]int
)

func findSink(u int, pb *Prober) bool {
	if pb.Nodes[u].Type == -1 {
		return true
	}
	mark[u] = 1
	for _, v := range pb.Graph[u] {
		if mark[v] == 0 && pb.Nodes[v].Stat != 0 {
			if findSink(v, pb) {
				return true
			}
		}
	}
	return false
}

func checkCap(pb *Prober) bool {
	for i, v := range pb.Nodes {
		if v.Type == 1 && v.Stat == 1 {
			for u := range mark {
				mark[u] = 0
			}
			if !findSink(i, pb) {
				return false
			}
		}
	}
	return true
}

func findCycle(u int, pb *Prober) bool {
	for _, v := range pb.Graph[u] {
		if mark[v] == mark[u] {
			return true
		}
		if mark[v] == 0 {
			mark[v] = mark[u]
			if findCycle(v, pb) {
				return true
			}
		}
	}
	return false
}

func checkDAG(pb *Prober) bool {
	for i := range mark {
		mark[i] = 0
	}
	for i := range pb.Nodes {
		if mark[i] == 0 {
			mark[i] = i + 1
			if findCycle(i, pb) {
				return false
			}
		}
	}
	return true
}

func setup(pb *Prober) error {
	pb.Nodid = make(map[string]int)
	for i, v := range pb.Nodes {
		if _, ok := pb.Nodid[v.Name]; ok {
			return fmt.Errorf("prober.go:setup: Found duplicate names for node %s!", v.Name)
		}
		pb.Nodid[v.Name] = i
	}
	for _, v := range pb.Pipes {
		pb.Graph[pb.Nodid[v.From]] = append(pb.Graph[pb.Nodid[v.From]], pb.Nodid[v.To])
	}
	return nil
}

func (pb *Prober) Check() bool {
	return checkDAG(pb) && checkCap(pb)
}

func (pb *Prober) Init(fileName string) error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, pb); err != nil {
		return err
	}
	if err := setup(pb); err != nil {
		return err
	}
	if !pb.Check() {
		return fmt.Errorf("prober.go:Init: %s is not a valid prober!", pb.Name)
	}
	return nil
}
