package turing

import (
	"bufio"
	"errors"
	"log"
	"strings"
)

//图灵机的构成组件：纸带（磁盘）、读写头、转换规则（代码）、寄存器
type TuringMachine struct {
	tapes        []byte
	head         int
	transitRules map[TransitRuleKey]TransitRuleValue
	register     string //current status
}

type TransitRuleKey struct {
	currentState string
	input        byte
}

type TransitRuleValue struct {
	newState   string
	valToWrite byte
	headMove   byte
}

const InitialStatus string = "q0"
const AcceptStatus string = "qAccept"
const ErrorStatus string = "qReject"

func NewMachine(transitRules map[TransitRuleKey]TransitRuleValue) *TuringMachine {
	machine := TuringMachine{
		tapes:        make([]byte, 2048),
		head:         0,
		transitRules: transitRules,
		register:     InitialStatus,
	}
	return &machine
}

func NewFromReader(scanner *bufio.Scanner) (*TuringMachine, error) {
	transitRules := map[TransitRuleKey]TransitRuleValue{}

	i := 0
	var previous string
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			continue
		}

		if (i & 1) == 1 {
			rukeKey := parseRuleKey(previous)
			ruleValue := parseRuleValue(text)
			transitRules[rukeKey] = ruleValue
		}

		previous = text
		i++
	}

	return NewMachine(transitRules), nil
}

func (m *TuringMachine) Execute(input []byte) (status string, output []byte) {
	//Init tapes
	for index := range m.tapes {
		m.tapes[index] = '_'
	}
	copy(m.tapes, input)
	//State transition
	for {
		//Read a value in tapes by current head
		input := m.tapes[m.head]
		//Find transition rule according to input
		ruleValue, err := m.findRuleValue(input)
		if err != nil {
			log.Fatalln(err)
			return ErrorStatus, m.tapes
		}

		//Advance machine status
		m.register = ruleValue.newState
		//Write value to tapes
		m.tapes[m.head] = ruleValue.valToWrite
		//Move head
		switch ruleValue.headMove {
		case '<':
			m.head--
		case '>':
			m.head++
		default:
		}

		//Stop or continue
		if ruleValue.newState == ErrorStatus || ruleValue.newState == AcceptStatus {
			return ruleValue.newState, m.tapes
		}
	}
}

func (m *TuringMachine) Reset() {
	newTuring := NewMachine(m.transitRules)
	m.tapes = newTuring.tapes
	m.head = newTuring.head
	m.register = newTuring.register
}

func (m *TuringMachine) findRuleValue(input byte) (TransitRuleValue, error) {
	ruleKey := TransitRuleKey{
		currentState: m.register,
		input:        input,
	}

	ruleValue, ok := m.transitRules[ruleKey]
	if !ok {
		return TransitRuleValue{}, errors.New("not found rule")
	}
	return ruleValue, nil
}

func parseRuleKey(ruleKeyStr string) TransitRuleKey {
	strs := strings.Split(ruleKeyStr, ",")
	currentState := strs[0]
	input := strs[1][0]

	return TransitRuleKey{
		currentState: currentState,
		input:        input,
	}
}

func parseRuleValue(ruleKeyStr string) TransitRuleValue {
	strs := strings.Split(ruleKeyStr, ",")
	newState := strs[0]
	valToWrite := strs[1][0]
	move := strs[2][0]

	return TransitRuleValue{
		newState:   newState,
		valToWrite: valToWrite,
		headMove:   move,
	}
}

//TODO: GO中map的key如果是struct类型，是如何比对的？
//TODO: 怎么判断key不存在
//TODO: map在并发情况下怎么上锁
