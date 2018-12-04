package libtc

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func execute(name string, params ...string) ([]string, error) {

	fmt.Printf("  Exec: %s %s\n", name, strings.Join(params, " "))
	var cmd = exec.Command(name, params...)
	var out, eout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &eout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("  Error: %s %s\n", err.Error(), eout.String())
		return nil, err
	}

	return strings.Split(out.String(), "\n"), nil
}

func findString(s string, num int, base int) int {
	var strs = strings.Split(s, ":")
	if num >= len(strs) {
		return 0
	}
	res, _ := strconv.ParseInt(strs[num], base, 0)
	return int(res)
}

func s2i(s string) int {
	var i, e = strconv.Atoi(s)
	if e == nil {
		return i
	}
	return 0
}

func hex2uint(s string) uint32 {
	var i, e = strconv.ParseUint(s, 16, 0)
	if e == nil {
		return uint32(i)
	}
	return 0
}

func i2hex(i int) string {
	return strconv.FormatInt(int64(i), 16)
}

func isIdent(s string) bool {
	var idre = regexp.MustCompile("^[[:alpha:]][[:word:]]*$")
	return idre.Match([]byte(s))
}

func parseOptions(params []string) []Option {
	var opts = []Option{}
	var key string
	var value []string
	for _, param := range params {
		if len(param) == 0 {
			continue
		}
		if isIdent(param) {
			if key != "" {
				if len(value) != 0 {
					opts = append(opts, Option{key, strings.Join(value, " ")})
					key, value = param, []string{}
					continue
				}
			} else {
				key = param
				continue
			}
		}
		value = append(value, param)
	}
	if key != "" && len(value) != 0 {
		opts = append(opts, Option{key, strings.Join(value, " ")})
	}
	return opts
}

func split(s string) []string {
	var res = regexp.MustCompile("[[:space:]]+").Split(s, -1)
	var final []string
	for _, r := range res {
		if len(r) != 0 {
			final = append(final, r)
		}
	}
	return final
}

func opt2params(opts []Option) []string {
	var res []string
	for _, opt := range opts {
		var params = split(opt.Value)
		res = append(res, opt.Name)
		res = append(res, params...)
	}
	return res
}

func act2params(acts []Action) []string {
	var res []string
	for _, act := range acts {
		res = append(res, ACTION_PARAM, act.Type)
		res = append(res, act.Options...)
	}
	return res
}

func normalizeMac(mac string) string {
	var res = mac
	res = strings.ToLower(res)
	res = regexp.MustCompile("[[:^xdigit:]]").ReplaceAllString(res, "")
	if len(res) < 12 {
		res = strings.Repeat("0", 12-len(res)) + res
	} else {
		res = res[:12]
	}
	res = regexp.MustCompile("(..)").ReplaceAllString(res, "$1:")
	return res[:17]
}

func parseMac(mac string) []int {
	var octets = strings.Split(normalizeMac(mac), ":")
	var res []int
	for _, o := range octets {
		var oo, err = strconv.ParseInt(o, 16, 32)
		if err != nil {
			continue
		}
		res = append(res, int(oo))
	}
	return res
}

func hex(val uint32, size int) string {
	switch size {
	case 8:
		return fmt.Sprintf("%#02x", val)
	case 16:
		return fmt.Sprintf("%#04x", val)
	default:
		return fmt.Sprintf("%#08x", val)
	}
}

func mask(length, offset uint) uint32 {
	var res int64 = 1<<length - 1
	return uint32(res) << offset
}

func parseIp(ip string) []int {
	var octets = strings.Split(ip, ".")
	var res []int
	for _, o := range octets {
		var oo, err = strconv.ParseInt(o, 10, 32)
		if err != nil {
			continue
		}
		res = append(res, int(oo))
	}
	for len(res) < 4 {
		res = append(res, 0)
	}
	return res
}
