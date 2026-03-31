package reader

import (
	"bufio"
	"os"
	"strings"
)

type Hosts struct {
	Entries map[string]string
}

func ParseHosts() (*Hosts, error) {
	f, err := os.Open("/etc/hosts")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	hosts := &Hosts{
		Entries: make(map[string]string),
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.Split(line, "#")[0]
		fields := strings.Fields(line)

		if fields[0] == "0.0.0.0" {
			hosts.Entries[fields[1]] = fields[0]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hosts, nil
}
