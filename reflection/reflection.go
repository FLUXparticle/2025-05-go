package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	X     int
	Y     float64
	Debug bool
	Name  string
}

func ParseConfig(lines []string, out any) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("out must be pointer to struct")
	}

	structVal := v.Elem()
	structType := structVal.Type()

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		field, ok := structType.FieldByNameFunc(func(name string) bool {
			return strings.EqualFold(name, key) // ignore case
		})
		if !ok {
			continue
		}

		fv := structVal.FieldByIndex(field.Index)
		if !fv.CanSet() {
			continue
		}

		switch fv.Kind() {
		case reflect.Int:
			n, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			fv.SetInt(int64(n))

		case reflect.Float64:
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fv.SetFloat(f)

		case reflect.Bool:
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fv.SetBool(b)

		case reflect.String:
			fv.SetString(value)
		}
	}

	return nil
}

func main() {
	file, err := os.Open("config.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}

	var cfg Config
	err = ParseConfig(lines, &cfg)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cfg)
}
