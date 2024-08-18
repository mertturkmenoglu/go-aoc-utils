package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func ReadLines(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func Sum[T Number](arr []T) T {
	var s T = 0

	for _, v := range arr {
		s += v
	}

	return s
}

func Prod[T Number](arr []T) T {
	var s T = 1

	for _, v := range arr {
		s *= v
	}

	return s
}

func IsFloatString(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsIntString(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func GCD(a int, b int) int {
	tmp := b

	for b != 0 {
		b = a % b
		a = tmp
		tmp = b
	}

	return a
}

func LCM(a int, b int) float32 {
	return float32(a*b) / float32(GCD(a, b))
}

type Pos[T comparable] struct {
	Fst T
	Snd T
}

func (p *Pos[T]) IsEqual(other Pos[T]) bool {
	return p.Fst == other.Fst && p.Snd == other.Snd
}

func (p *Pos[T]) Destruct() (T, T) {
	return p.Fst, p.Snd
}

// export type TGrid<T> = T[][];

var HVAdjMap []Pos[int] = []Pos[int]{
	{Fst: -1, Snd: 0},
	{Fst: 1, Snd: 0},
	{Fst: 0, Snd: -1},
	{Fst: 0, Snd: 1},
}

var AdjMap []Pos[int] = []Pos[int]{
	{Fst: -1, Snd: 0},
	{Fst: 1, Snd: 0},
	{Fst: 0, Snd: -1},
	{Fst: 0, Snd: 1},
	{Fst: -1, Snd: -1},
	{Fst: -1, Snd: 1},
	{Fst: 1, Snd: -1},
	{Fst: 1, Snd: 1},
}

type Grid[T any] struct {
	Mtr [][]T
}

func NewGrid[T any](list [][]T) Grid[T] {
	return Grid[T]{
		Mtr: list,
	}
}

func CreateGrid[T any](rows int, cols int, fill T) Grid[T] {
	m := make([]([]T), rows)

	for i := 0; i < rows; i++ {
		m[i] = make([]T, cols)

		for j := 0; j < cols; j++ {
			m[i][j] = fill
		}
	}

	return Grid[T]{
		Mtr: m,
	}
}

func (g *Grid[T]) GetRow(i int) []T {
	return g.Mtr[i]
}

func (g *Grid[T]) GetCol(i int) []T {
	rows := len(g.Mtr)
	arr := make([]T, rows)

	for j := 0; j < rows; j++ {
		arr[j] = g.Mtr[j][i]
	}

	return arr
}

func (g *Grid[T]) At(row int, col int) T {
	return g.Mtr[row][col]
}

func (g *Grid[T]) AtPos(pos Pos[int]) T {
	return g.Mtr[pos.Fst][pos.Snd]
}

func (g *Grid[T]) Set(row int, col int, v T) {
	g.Mtr[row][col] = v
}

func (g *Grid[T]) SetPos(pos Pos[int], v T) {
	g.Mtr[pos.Fst][pos.Snd] = v
}

func (g *Grid[T]) Dims() (int, int) {
	return len(g.Mtr), len(g.Mtr[0])
}

func (g *Grid[T]) IsPosInGrid(pos Pos[int]) bool {
	row, col := pos.Destruct()
	rowCount, colCount := g.Dims()

	if row < 0 || row >= rowCount {
		return false
	}

	if col < 0 || col >= colCount {
		return false
	}

	return true
}

func (g *Grid[T]) IsValueInGrid(v T, comp func(T, T) bool) bool {
	for _, row := range g.Mtr {
		for _, cell := range row {
			if comp(cell, v) {
				return true
			}
		}
	}

	return false
}

func (g *Grid[T]) GetPosOfValue(v T, comp func(T, T) bool) *Pos[int] {
	for i, row := range g.Mtr {
		for j, cell := range row {
			if comp(cell, v) {
				return &Pos[int]{Fst: i, Snd: j}
			}
		}
	}
	return nil
}

//   hash(): string {
//     return this.toString();
//   }

//   toString(): string {
//     return JSON.stringify(this.mtr);
//   }

//   equal(other: Grid<T>): boolean {
//     return JSON.stringify(this) === JSON.stringify(other);
//   }
