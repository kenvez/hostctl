package blocker

import (
	"bufio"
	"errors"
	"os"

	"github.com/kenvez/hostctl/internal/reader"
)

func Block(hosts *reader.Hosts, domain string) error {
	if _, exists := hosts.Entries[domain]; exists {
		return errors.New("domain already blocked")
	}

	hosts.Entries[domain] = "0.0.0.0"

	return nil
}

func Unblock(hosts *reader.Hosts, domain string) error {
	if _, exists := hosts.Entries[domain]; !exists {
		return errors.New("domain not found")
	}

	delete(hosts.Entries, domain)

	return nil
}

func Save(hosts *reader.Hosts) error {
	f, err := os.Create("hosts")
	if err != nil {
		return err
	}
	defer f.Close()

	hostsFile, err := os.Open("/etc/hosts")
	if err != nil {
		return err
	}
	defer hostsFile.Close()
	scanner := bufio.NewScanner(hostsFile)

	for scanner.Scan() {
		f.Write([]byte(scanner.Text()))
		f.Write([]byte("\n"))
	}

	f.Write([]byte("# blocked websites\n"))

	for k, v := range hosts.Entries {
		f.Write([]byte(v))
		f.Write([]byte("\t"))
		f.Write([]byte(k))
		f.Write([]byte("\n"))
	}

	return nil
}
