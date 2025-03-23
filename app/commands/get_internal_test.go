package commands

import (
	"reflect"
	"testing"
)

//пакет не меняем

func TestCommandSelector(t *testing.T) {

	args := new(Arguments)
	args.CmdName = "SET"
	args.Key = "1"
	args.Value = "1"
	args.TTL = 100

	cmd := selectCommand(args)

	switch cmd.(type) {
	case *SetCommand:
	default:
		t.Errorf("Ошибка выбора команды, выбрана не та команда %v : %s", reflect.TypeOf(cmd), args.CmdName)
	}

	args.CmdName = CommandGet
	cmd = selectCommand(args)

	switch cmd.(type) {
	case *GetCommand:
	default:
		t.Errorf("Ошибка выбора команды, выбрана не та команда %v : %s", reflect.TypeOf(cmd), args.CmdName)
	}

	args.CmdName = CommandSet
	cmd = selectCommand(args)

	switch cmd.(type) {
	case *SetCommand:
	default:
		t.Errorf("Ошибка выбора команды, выбрана не та команда %v : %s", reflect.TypeOf(cmd), args.CmdName)
	}

	args.CmdName = CommandDelete
	cmd = selectCommand(args)

	switch cmd.(type) {
	case *DeleteCommand:
	default:
		t.Errorf("Ошибка выбора команды, выбрана не та команда %v : %s", reflect.TypeOf(cmd), args.CmdName)
	}

	args.CmdName = CommandInfo
	cmd = selectCommand(args)

	switch cmd.(type) {
	case *InfoCommand:
	default:
		t.Errorf("Ошибка выбора команды, выбрана не та команда %v : %s", reflect.TypeOf(cmd), args.CmdName)
	}

	args.CmdName = CommandRestore
	cmd = selectCommand(args)

	switch cmd.(type) {
	case *RestoreCommand:
	default:
		t.Errorf("Ошибка выбора команды, выбрана не та команда %v : %s", reflect.TypeOf(cmd), args.CmdName)
	}
}

func TestArgsParser_Key(t *testing.T) {

	msg := "GET k=1"
	args, err := NewArgsFromString(msg)

	if err != nil {
		t.Error(err)
	}

	if args.Key != "1" {
		t.Errorf("Ошибка установки аргумента key, передано k=%s : получено %s", "1", args.Key)
	}

	msg = "GET key=1"
	args, err = NewArgsFromString(msg)

	if err != nil {
		t.Error(err)
	}

	if args.Key != "1" {
		t.Errorf("Ошибка установки аргумента key, передано key=%s : получено %s", "1", args.Key)
	}
}

func TestArgsParser_Value(t *testing.T) {

	msg := "SET k=1 v=1"
	args, err := NewArgsFromString(msg)

	if err != nil {
		t.Error(err)
	}

	if args.Value != "1" {
		t.Errorf("Ошибка установки аргумента value, передано v=%s : получено %s", "1", args.Value)
	}

	msg = "SET k=1 value=1"
	args, err = NewArgsFromString(msg)

	if err != nil {
		t.Error(err)
	}

	if args.Value != "1" {
		t.Errorf("Ошибка установки аргумента value, передано value=%s : получено %s", "1", args.Value)
	}
}

func TestArgsParser_TTL(t *testing.T) {

	msg := "SET k=1 v=1 ttl=3600"
	args, err := NewArgsFromString(msg)

	if err != nil {
		t.Error(err)
	}

	if args.TTL != 3600 {
		t.Errorf("Ошибка установки аргумента TTL, передано ttl=%v : получено %v", "3600", args.TTL)
	}
}

func TestArgsParser_CMD(t *testing.T) {

	msg := "SET k=1 v=1 ttl=3600"
	args, err := NewArgsFromString(msg)

	if err != nil {
		t.Error(err)
	}

	if args.CmdName != "SET" {
		t.Errorf("Ошибка установки аргумента CmdName, передано name %v : получено %v", "SET", args.CmdName)
	}
}
