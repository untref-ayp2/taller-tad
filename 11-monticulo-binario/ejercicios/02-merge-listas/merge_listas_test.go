package merge_listas

import (
	"reflect"
	"testing"

	"github.com/untref-ayp2/data-structures/list"
)

func intCmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func buildList(vals ...int) list.List[int] {
	l := list.NewSinglyLinkedList[int]()
	for _, v := range vals {
		l.Append(v)
	}
	return l
}

func checkResult(t *testing.T, result list.List[int], expected []int) {
	t.Helper()
	vals := result.Values()
	if !reflect.DeepEqual(vals, expected) {
		t.Errorf("Esperaba %v, obtuve %v", expected, vals)
	}
}

func TestMergeKListasTresListas(t *testing.T) {
	l1 := buildList(1, 4, 7)
	l2 := buildList(2, 5, 8)
	l3 := buildList(3, 6, 9)

	result := MergeKListas([]list.List[int]{l1, l2, l3}, intCmp)
	checkResult(t, result, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestMergeKListasListasVacias(t *testing.T) {
	l1 := buildList()
	l2 := buildList(1, 3)
	l3 := buildList()

	result := MergeKListas([]list.List[int]{l1, l2, l3}, intCmp)
	checkResult(t, result, []int{1, 3})
}

func TestMergeKListasUnaLista(t *testing.T) {
	l1 := buildList(10, 20, 30)
	result := MergeKListas([]list.List[int]{l1}, intCmp)
	checkResult(t, result, []int{10, 20, 30})
}

func TestMergeKListasTodasVacias(t *testing.T) {
	result := MergeKListas([]list.List[int]{buildList(), buildList()}, intCmp)
	if !result.IsEmpty() {
		t.Error("Resultado debería estar vacío cuando todas las listas están vacías")
	}
}

func TestMergeKListasDuplicados(t *testing.T) {
	l1 := buildList(1, 3, 5)
	l2 := buildList(1, 3, 5)

	result := MergeKListas([]list.List[int]{l1, l2}, intCmp)
	checkResult(t, result, []int{1, 1, 3, 3, 5, 5})
}

func TestMergeKListasUnElementoPorLista(t *testing.T) {
	l1 := buildList(10)
	l2 := buildList(5)
	l3 := buildList(20)

	result := MergeKListas([]list.List[int]{l1, l2, l3}, intCmp)
	checkResult(t, result, []int{5, 10, 20})
}
