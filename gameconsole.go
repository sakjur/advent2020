package advent2020

import (
	"fmt"
	"strconv"
)

type Op struct {
	Operator string
	Argument int
}

func GameConsoleDetectLoop(ops []Op) (int, bool) {
	acc := 0
	visited := make([]bool, len(ops))

	var incr int
	for i := 0; i < len(ops); i += incr {
		incr = 1

		if i < 0 || visited[i] {
			return acc, false
		}
		visited[i] = true

		switch op := ops[i]; op.Operator {
		case "acc":
			acc += op.Argument
		case "jmp":
			incr = op.Argument
		case "nop":
			// noop
		}
	}
	return acc, true
}

func GameConsoleAutoPatcher(ops []Op) (int, bool) {
	for i, op := range ops {
		if op.Operator != "nop" && op.Operator != "jmp" {
			continue
		}

		cpy := make([]Op, len(ops))
		copy(cpy, ops)

		newOp := "nop"
		if op.Operator == "nop" {
			newOp = "jmp"
		}

		cpy[i] = Op{
			Operator: newOp,
			Argument: op.Argument,
		}
		acc, ok := GameConsoleDetectLoop(cpy)
		if ok {
			return acc, true
		}
	}
	return 0, false
}

func GameConsoleOps(c <-chan string) ([]Op, error) {
	ops := []Op{}

	for rawOp := range c {
		if len(rawOp) < 6 {
			return nil, fmt.Errorf("expected instruction to be at least 6 chars long: %s", rawOp)
		}

		arg, err := strconv.Atoi(rawOp[4:])
		if err != nil {
			return nil, err
		}

		ops = append(ops, Op{
			Operator: rawOp[:3],
			Argument: arg,
		})
	}

	return ops, nil
}
