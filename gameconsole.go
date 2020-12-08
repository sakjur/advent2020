package advent2020

import (
	"fmt"
	"math"
	"strconv"
)

type Op struct {
	Operator string
	Argument int
}

func GameConsoleDetectLoop(ops []Op) (int, bool) {
	acc := 0
	visited := make([]bool, len(ops))

	incr := 1
	for i := 0; i < len(ops); i += incr {
		incr = 1

		if i < 0 {
			return math.MinInt64, false
		}

		instr := ops[i]
		if visited[i] {
			return acc, false
		}

		switch instr.Operator {
		case "nop":
			// noop
		case "acc":
			acc += instr.Argument
		case "jmp":
			// subtracting one because loop increments by one
			incr = instr.Argument
		}

		visited[i] = true
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

		switch op.Operator {
		case "nop":
			cpy[i] = Op{
				Operator: "jmp",
				Argument: op.Argument,
			}
		case "jmp":
			cpy[i] = Op{
				Operator: "nop",
				Argument: op.Argument,
			}
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
		op := rawOp[:3]
		rawArg := rawOp[4:]

		if rawArg[0] == '+' {
			rawArg = rawArg[1:]
		}
		arg, err := strconv.Atoi(rawArg)
		if err != nil {
			return nil, err
		}

		ops = append(ops, Op{
			Operator: op,
			Argument: arg,
		})
	}

	return ops, nil
}
