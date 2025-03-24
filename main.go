package main

import (
	"errors"
	"fmt"
	"slices"
)


type board []uint8


func parse(str string) (board, error) {
    if (len(str) != 81) {
        return nil, errors.New("invalid input length")
    }
    var b board = make([]uint8, 81)
    for i, v := range str {
        if v < '0' || v > '9' {
            return nil, errors.New("invalid char")
        }
        b[i] = uint8(v - '0')
    }
    if !b.valid() {
        return nil, errors.New("board not valid")
    }
    return b, nil
}

func itorc(i uint8) (uint8, uint8) {
    return i / 9, i % 9
}

func rctoi(r, c uint8) uint8 {
    return r * 9 + c
}


func (b board) rowContains(r, v uint8) bool {
    return slices.Contains(b[rctoi(r, 0) : rctoi(r, 9)], v)
}

func (b board) colContains(c, v uint8) bool {
    for ; c < 81; c += 9 {
        if b[c] == v {
            return true
        }
    }
    return false
}

func (b board) boxContains(i, v uint8) bool {
    bi := i / 3 * 3 * 9 + i % 3 * 3
    for j := range uint8(9) {
        if b[bi + j / 3 * 9 + j % 3] == v {
            return true
        }
    }
    return false
}


func (b board) moveValid(i, v uint8) bool {
    r, c := itorc(i)
    bi := r / 3 * 3 + c / 3
    return !b.rowContains(r, v) && !b.colContains(c, v) && !b.boxContains(bi, v)
}

func (b board) getRow(r uint8) []uint8 {
    return slices.Clone(b[rctoi(r, 0) : rctoi(r, 9)])
}

func (b board) getCol(c uint8) []uint8 {
    col := make([]uint8, 9)
    for i := range uint8(9) {
        col[i] = b[rctoi(i, c)]
    }
    return col
}

func (b board) getBox(i uint8) []uint8 {
    box := make([]uint8, 9)
    r := i / 3 * 3
    c := i % 3 * 3
    for i := range uint8(9) {
        box[i] = b[rctoi(r + i / 3, c + i % 3)]
    }
    return box
}

func spliceValid(s []uint8) bool {
    has := make([]bool, 9)
    for _, v := range s {
        if v == 0 {
            continue
        }
        if has[v - 1] {
            return false
        }
        has[v - 1] = true
    }
    return true
}
func (b board) valid() bool {
    for i := range uint8(9) {
        if !spliceValid(b.getRow(i)) ||
           !spliceValid(b.getCol(i)) ||
           !spliceValid(b.getBox(i)) {
            return false
        }
    }
    return true
}


func (b board) backtrack() {
    next := func(i uint8) uint8 {
        for {
            if b[i] == 0 {
                return i
            }
            i++
            if i == 81 {
                return 81
            }
        }
    }
    stack := make([]uint8, 81)
    stackIdx := uint8(1)
    i := next(0)
    stack[0] = i
    for {
        foundValidTry := false
        for try := b[i] + 1; try <= 9; try++ {
            if b.moveValid(i, try) {
                b[i] = try
                foundValidTry = true
                stack[stackIdx] = i
                stackIdx++
                break
            }
        }
        if foundValidTry {
            i = next(i)
            if i == 81 {
                return
            }
        } else {
            b[i] = 0
            stackIdx--
            i = stack[stackIdx]
        }
    }
}


func (b board) solve() {
    b.backtrack()
}


func (b board) print() {
    fmt.Println("valid:", b.valid())
    for i, v := range b {
        if i != 0 && i % 9 == 0 {
            fmt.Println()
        }
        fmt.Print(v)
    }
    fmt.Println()
}





func main() {

    var str string
    fmt.Scanln(&str)
    board, err := parse(str)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    board.solve()

    board.print()

}



