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

	if err := save(hosts); err != nil {
		return err
	}

	return nil
}

func Unblock(hosts *reader.Hosts, domain string) error {
	if _, exists := hosts.Entries[domain]; !exists {
		return errors.New("domain not found")
	}

	delete(hosts.Entries, domain)

	if err := save(hosts); err != nil {
		return err
	}

	return nil
}

func save(hosts *reader.Hosts) error {
	tmpFile, err := os.CreateTemp("", "hostctl-*")
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	hostsFile, err := os.Open("/etc/hosts")
	if err != nil {
		return err
	}
	defer hostsFile.Close()
	scanner := bufio.NewScanner(hostsFile)

	inBlockedSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "# blocked websites" {
			inBlockedSection = true
		}
		if !inBlockedSection {
			tmpFile.Write([]byte(line))
			tmpFile.Write([]byte("\n"))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	tmpFile.Write([]byte("# blocked websites\n"))

	for k, v := range hosts.Entries {
		tmpFile.Write([]byte(v))
		tmpFile.Write([]byte("\t"))
		tmpFile.Write([]byte(k))
		tmpFile.Write([]byte("\n"))
	}

	if err := os.Rename(tmpFile.Name(), "/etc/hosts"); err != nil {
		return err
	}

	return nil
}
