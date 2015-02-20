package registration

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/kitekey"
	"github.com/koding/klient/protocol"
)

// registerWithPassword registers to the given kontrolURL. The user
func WithPassword(kontrolURL, username string) error {
	var err error

	// Open up a prompt if the username is not passed via a flag
	if username == "" {
		username, err = ask("Username:")
		if err != nil {
			return err
		}
		// User can just press enter to use the default on the prompt
		if username == "" {
			return errors.New("Username can not be empty.")
		}
	}

	k := kite.New("klient", protocol.Version)
	k.Config.Environment = protocol.Environment
	k.Config.Region = protocol.Region
	k.Config.Transport = config.XHRPolling
	k.Config.Username = username

	if _, err := kitekey.Read(); err == nil {
		result, err := ask("An existing ~/.kite/kite.key detected. Type 'yes' to override and continue:")
		if err != nil {
			return err
		}

		if result != "yes" {
			return errors.New("aborting registration")
		}
	}

	kontrol := k.NewClient(kontrolURL)
	if err := kontrol.Dial(); err != nil {
		return err
	}

	result, err := kontrol.TellWithTimeout("registerMachine", 5*time.Minute, username)
	if err != nil {
		return err
	}

	if err := kitekey.Write(result.MustString()); err != nil {
		return err
	}

	fmt.Println("Registered successfully")
	return nil
}

// ask asks for an input from standard input and returns the result back. It is
// extracted from mitcellh/cli to be used as a standalone function.
func ask(query string) (string, error) {
	if _, err := fmt.Fprint(os.Stdout, query+" "); err != nil {
		return "", err
	}

	// Register for interrupts so that we can catch it and immediately
	// return...
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	// Ask for input in a go-routine so that we can ignore it.
	errCh := make(chan error, 1)
	lineCh := make(chan string, 1)
	go func() {
		r := bufio.NewReader(os.Stdin)
		line, err := r.ReadString('\n')
		if err != nil {
			errCh <- err
			return
		}

		lineCh <- strings.TrimRight(line, "\r\n")
	}()

	select {
	case err := <-errCh:
		return "", err
	case line := <-lineCh:
		return line, nil
	case <-sigCh:
		// Print a newline so that any further output starts properly
		// on a new line.
		fmt.Fprintln(os.Stdout)

		return "", errors.New("interrupted")
	}
}
