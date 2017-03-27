package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"koding/klient/machine/mount/sync"
	"koding/klientctl/ctlcli"
	"koding/klientctl/endpoint/machine"

	"github.com/codegangsta/cli"
	"github.com/dustin/go-humanize"
	"github.com/koding/logging"
)

// MachineListCommand returns list of remote machines belonging to the user or
// that can be accessed by her.
func MachineListCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	// List command doesn't support identifiers.
	idents, err := getIdentifiers(c)
	if err != nil {
		return 1, err
	}
	if err := identifiersLimit(idents, "machine", 0, 0); err != nil {
		return 1, err
	}

	opts := &machine.ListOptions{
		Log: log.New("machine:list"),
	}

	infos, err := machine.List(opts)
	if err != nil {
		return 1, err
	}

	if c.Bool("json") {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "\t")
		enc.Encode(infos)
		return 0, nil
	}

	tabListFormatter(os.Stdout, infos)
	return 0, nil
}

// MachineSSHCommand allows to SSH into remote machine.
func MachineSSHCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	// SSH command must have only one identifier.
	idents, err := getIdentifiers(c)
	if err != nil {
		return 1, err
	}
	if err := identifiersLimit(idents, "machine", 1, 1); err != nil {
		return 1, err
	}

	opts := &machine.SSHOptions{
		Identifier: idents[0],
		Username:   c.String("username"),
		Log:        log.New("machine:ssh"),
	}

	if err := machine.SSH(opts); err != nil {
		return 1, err
	}

	return 0, nil
}

// MachineMountCommand allows to create mount between remote and local machines.
func MachineMountCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	defer fixDescription("Mount remote folder to a local directory.")()
	// Mount command has two identifiers - remote machine:path and local path.
	idents, err := getIdentifiers(c)
	if err != nil {
		return 1, err
	}
	if len(idents) == 0 {
		return 0, cli.ShowSubcommandHelp(c)
	}
	if err := identifiersLimit(idents, "argument", 2, 2); err != nil {
		return 1, err
	}
	ident, remotePath, path, err := mountAddress(idents)
	if err != nil {
		return 1, err
	}

	opts := &machine.MountOptions{
		Identifier: ident,
		Path:       path,
		RemotePath: remotePath,
		Log:        log.New("machine:mount"),
	}

	if err := machine.Mount(opts); err != nil {
		return 1, err
	}

	return 0, nil
}

// MachineListMountCommand lists available mounts.
func MachineListMountCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	// Mount list command doesn't need identifiers.
	idents, err := getIdentifiers(c)
	if err != nil {
		return 1, err
	}
	if err := identifiersLimit(idents, "mount", 0, 0); err != nil {
		return 1, err
	}

	opts := &machine.ListMountOptions{
		MountID: c.String("filter"),
		Log:     log.New("machine:mount:list"),
	}

	mounts, err := machine.ListMount(opts)
	if err != nil {
		return 1, err
	}

	if c.Bool("json") {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "\t")
		enc.Encode(mounts)
		return 0, nil
	}

	tabListMountFormatter(os.Stdout, mounts)
	return 0, nil
}

// MachineUmountCommand removes the mount.
func MachineUmountCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	// Umount command needs exactly one identifier. Either mount ID or
	// mount local path.
	idents, err := getIdentifiers(c)
	if err != nil {
		return 1, err
	}
	if err := identifiersLimit(idents, "mount", 1, 1); err != nil {
		return 1, err
	}

	opts := &machine.UmountOptions{
		Identifier: idents[0],
		Log:        log.New("machine:umount"),
	}

	if err := machine.Umount(opts); err != nil {
		return 1, err
	}

	return 0, nil
}

// MachineExecCommand runs a command in a started machine.
func MachineExecCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	if c.NArg() < 2 {
		cli.ShowCommandHelp(c, "exec")
		return 1, nil
	}

	done := make(chan int, 1)

	opts := &machine.ExecOptions{
		Cmd:  c.Args()[1],
		Args: c.Args()[2:],
		Stdout: func(line string) {
			fmt.Println(line)
		},
		Stderr: func(line string) {
			fmt.Fprintln(os.Stderr, line)
		},
		Exit: func(exit int) {
			done <- exit
			close(done)
		},
	}

	if s := c.Args()[0]; strings.HasPrefix(s, "@") {
		opts.MachineID = s[1:]
	} else {
		if filepath.IsAbs(s) {
			var err error
			if s, err = filepath.Abs(s); err != nil {
				return 1, err
			}
		}

		opts.Path = s
	}

	pid, err := machine.Exec(opts)
	if err != nil {
		return 1, err
	}

	ctlcli.CloseOnExit(ctlcli.CloseFunc(func() error {
		select {
		case <-done:
			return nil
		default:
			return machine.Kill(&machine.KillOptions{
				MachineID: opts.MachineID,
				Path:      opts.Path,
				PID:       pid,
			})
		}
	}))

	return <-done, nil
}

// MachineCpCommand copies file(s) from one machine to another.
func MachineCpCommand(c *cli.Context, log logging.Logger, _ string) (int, error) {
	idents, err := getIdentifiers(c)
	if err != nil {
		return 1, err
	}
	fmt.Println("Identifiers:", idents)
	if len(idents) == 0 {
		return 0, cli.ShowSubcommandHelp(c)
	}
	if err := identifiersLimit(idents, "argument", 2, 2); err != nil {
		return 1, err
	}
	fmt.Println("Parsed")

	opts := &machine.CpOptions{
		Download:        true,
		Identifier:      idents[0],
		SourcePath:      "source",
		DestinationPath: "~/",
		Log:             log.New("machine:cp"),
	}

	if err := machine.Cp(opts); err != nil {
		return 1, err
	}

	return 0, nil
}

// getIdentifiers extracts identifiers and validate provided arguments.
// TODO(ppknap): other CLI libraries like Cobra have this out of the box.
func getIdentifiers(c *cli.Context) (idents []string, err error) {
	unknown := make([]string, 0)
	for _, arg := range c.Args() {
		if strings.HasPrefix(arg, "-") {
			unknown = append(unknown, arg)
			continue
		}

		idents = append(idents, arg)
	}

	if len(unknown) > 0 {
		plural := ""
		if len(unknown) > 1 {
			plural = "s"
		}

		return nil, fmt.Errorf("unrecognized argument%s: %s", plural, strings.Join(unknown, ", "))
	}

	return idents, nil
}

// identifiersLimit checks if the number of identifiers is in specified limits.
// If max is -1, there are no limits for the maximum number of identifiers.
func identifiersLimit(idents []string, kind string, min, max int) error {
	l := len(idents)
	switch {
	case l > 0 && min == 0:
		return fmt.Errorf("this command does not use %s identifiers", kind)
	case l < min:
		return fmt.Errorf("required at least %d %ss", min, kind)
	case max != -1 && l > max:
		return fmt.Errorf("too many %ss: %s", kind, strings.Join(idents, ", "))
	}

	return nil
}

// mountAddress checks if provided identifiers are valid from the mount
// perspective. The identifiers should satisfy the following format:
//
//  [ID|Alias|IP]:remote_directory/path local_directory/path
//
func mountAddress(idents []string) (ident, remotePath, path string, err error) {
	if len(idents) != 2 {
		return "", "", "", fmt.Errorf("invalid number of arguments: %s", strings.Join(idents, ", "))
	}

	remote := strings.Split(idents[0], ":")
	if len(remote) != 2 {
		return "", "", "", fmt.Errorf("invalid remote address format: %s", idents[0])
	}

	if path, err = filepath.Abs(idents[1]); err != nil {
		return "", "", "", fmt.Errorf("invalid format of local path %q: %s", idents[1], err)
	}

	return remote[0], remote[1], path, nil
}

func tabListFormatter(w io.Writer, infos []*machine.Info) {
	now := time.Now()
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "ID\tALIAS\tTEAM\tSTACK\tPROVIDER\tLABEL\tOWNER\tAGE\tIP\tSTATUS\n")
	for _, info := range infos {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			info.ID,
			info.Alias,
			dashIfEmpty(info.Team),
			dashIfEmpty(info.Stack),
			dashIfEmpty(info.Provider),
			info.Label,
			info.Owner,
			machine.ShortDuration(info.CreatedAt, now),
			info.IP,
			machine.PrettyStatus(info.Status, now),
		)
	}
	tw.Flush()
}

func tabListMountFormatter(w io.Writer, mounts map[string][]sync.Info) {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	// TODO: keep the mounts list sorted.
	fmt.Fprintf(tw, "ID\tMACHINE\tMOUNT\tFILES\tQUEUED\tSYNCING\tSIZE\n")
	for alias, infos := range mounts {
		for _, info := range infos {
			sign := info.Syncing
			fmt.Fprintf(tw, "%s\t%s\t%s\t%s/%s\t%s\t%s\t%s/%s\n",
				info.ID,
				alias,
				info.Mount,
				dashIfNegative(sign, info.Count),
				dashIfNegative(sign, info.CountAll),
				dashIfNegative(sign, info.Queued),
				errorIfNegative(info.Syncing),
				dashIfNegative(sign, humanize.IBytes(uint64(info.DiskSize))),
				dashIfNegative(sign, humanize.IBytes(uint64(info.DiskSizeAll))),
			)
		}
	}
	tw.Flush()
}

func errorIfNegative(val int) string {
	if val < 0 {
		return "err"
	}

	return strconv.Itoa(val)
}

func dashIfNegative(sign int, val interface{}) string {
	if sign < 0 {
		return "-"
	}

	return fmt.Sprint(val)
}

func dashIfEmpty(val string) string {
	if val == "" {
		return "-"
	}

	return val
}
