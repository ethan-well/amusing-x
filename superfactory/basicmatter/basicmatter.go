package basicmatter

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	Comment = "#"
	SectionS = "["
	sectionE = "]"
	KVSep = " "
	tagSep = ":"
)

type ConfigContainer struct {
	data map[string]*Section
}

type Section struct {
	// config key:value map
	data map[string]*ConfigValueItem
}

type ConfigValueItem struct {
	Key string
	Value string
	Line int64
}

func New() *ConfigContainer {
	return &ConfigContainer{}
}

func (c *ConfigContainer) Unmarshal(configObject interface{}, fileName string) error {
	var configFile string
	flag.StringVar(&configFile, "cf", "test.conf", "-cf is expected")
	flag.Parse()

	if len(fileName) != 0{
		configFile = fileName
	}

	err :=  c.parseFile(configFile)
	if err != nil {
		return err
	}

	err = c.parse(configObject)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigContainer) parse(configObject  interface{}) error {
	if c.data == nil {
		return nil
	}

	if reflect.TypeOf(configObject).Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("config obj is must be a ptr"))
	}

	sv := reflect.ValueOf(configObject).Elem()
	st := reflect.TypeOf(configObject).Elem()

	for i := 0; i < sv.NumField(); i ++ {
		tag := st.Field(i).Tag.Get("config")
		tagArr := strings.Split(tag, tagSep)
		if len(tagArr) != 2 {
			return errors.New(fmt.Sprintf("config object tag must split by %s", tagSep))
		}

		configSection := tagArr[0]
		configName := tagArr[1]

		if c.data == nil ||
			c.data[configSection] == nil ||
			c.data[configSection].data == nil ||
			c.data[configSection].data[configName] == nil {

			continue
		}

		if !sv.Field(i).CanSet() {
			return errors.New(fmt.Sprintf("struct filed: %s can not set", st.Field(i).Name))
		}

		configValueItem := c.data[configSection].data[configName]
		switch sv.Field(i).Kind() {
		case reflect.String:
			sv.Field(i).SetString(configValueItem.Value)
		case reflect.Bool:
			cv, err := strconv.ParseBool(configValueItem.Value)
			if err != nil {
				return errors.New(fmt.Sprintf("a bool-value is expected at config file line: %d", configValueItem.Line))
			}
			sv.Field(i).SetBool(cv)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if len(configValueItem.Value) == 0 {
				configValueItem.Value = "0"
			}
			cv, err := strconv.ParseInt(configValueItem.Value, 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("a int-value is expected at config file line: %d", configValueItem.Line))
			}
			sv.Field(i).SetInt(cv)
		case reflect.Uint:
			if len(configValueItem.Value) == 0 {
				configValueItem.Value = "0"
			}
			cv, err := strconv.ParseUint(configValueItem.Value, 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("a int-value is expected at config file line: %d", configValueItem.Line))
			}
			sv.Field(i).SetUint(cv)
		case reflect.Float32, reflect.Float64:
			if len(configValueItem.Value) == 0 {
				configValueItem.Value = "0"
			}
			cv, err := strconv.ParseFloat(configValueItem.Value, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("a int-value is expected at config file line: %d", configValueItem.Line))
			}
			sv.Field(i).SetFloat(cv)
		}
	}

	return nil
}

// 解析文件，将解析结果存放到 conf.data 种
func (c *ConfigContainer) parseFile(file string) error {
	fi, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fi.Close()

	var line int64
	rd := bufio.NewScanner(fi)

	var sectionStr string

	for rd.Scan() {
		line++

		row := rd.Text()

		row = strings.TrimSpace(row)
		// 不解析空行和注释
		if len(row) == 0 ||
			strings.HasPrefix(row, Comment) {
			continue
		}

		// 判断是不是 section
		if strings.HasPrefix(row, SectionS) {
			if !strings.HasSuffix(row, sectionE) {
				return errors.New(fmt.Sprintf("no end section: %s, at: %d", sectionE, line))
			}

			sectionStr = row[1 : len(row)-1]
			section, ok := c.data[sectionStr]
			if !ok {
				section = &Section{data: map[string]*ConfigValueItem{}}
			}

			if c.data == nil {
				c.data = map[string]*Section{}
			}

			c.data[sectionStr] = section

			continue
		}

		if len(sectionStr) == 0 {
			return errors.New(fmt.Sprintf("section lost at line: %d", line))
		}

		// 分析行的 key value
		kvArr := strings.Split(row, KVSep)
		if len(kvArr) != 2 {
			return errors.New(fmt.Sprintf("config line: %d is expected only two column", line))
		}

		c.data[sectionStr].data[kvArr[0]] = &ConfigValueItem{
			Key:   kvArr[0],
			Value: kvArr[1],
			Line:  line,
		}
	}

	return nil
}