package client

import (
	"context"
	"time"

	"koding/kites/kloud/klient"
	"koding/klient/machine"
	"koding/klient/machine/index"
	"koding/klient/os"

	"github.com/koding/kite"
)

// KiteBuilder implements Builder interface. It creates Kite clients that use
// kite query string as their source address.
type KiteBuilder struct {
	pool *klient.KlientPool
}

// NewKiteBuilder creates a new KiteBuilder instance.
func NewKiteBuilder(k *kite.Kite) *KiteBuilder {
	return &KiteBuilder{
		pool: klient.NewPool(k),
	}
}

// Ping uses kite network that stores kite's query string to ping the machine.
func (kb *KiteBuilder) Ping(dynAddr DynamicAddrFunc) (machine.Status, machine.Addr, error) {
	addr, err := dynAddr("kite")
	if err != nil {
		return machine.Status{}, machine.Addr{}, err
	}

	if _, err := kb.pool.Get(addr.Value); err != nil {
		return machine.Status{}, machine.Addr{}, err
	}

	return machine.Status{
		State: machine.StateConnected,
		Since: time.Now(),
	}, addr, nil
}

// Build builds new kite client that will connect to machine using provided
// address.
func (kb *KiteBuilder) Build(ctx context.Context, addr machine.Addr) Client {
	return &kiteClient{
		ctx:  ctx,
		addr: addr.Value,
		pool: kb.pool,
	}
}

// kiteClient implements Client interface it uses klient from internal pool.
type kiteClient struct {
	ctx  context.Context
	addr string
	pool *klient.KlientPool
}

// CurrentUser returns remote machine current username.
func (kc *kiteClient) CurrentUser() (string, error) {
	return kc.get().CurrentUser()
}

// Abs returns absolute representation of given path.
func (kc *kiteClient) Abs(path string) (string, bool, bool, error) {
	return kc.get().Abs(path)
}

// SSHAddKeys adds SSH public keys to user's authorized_keys file.
func (kc *kiteClient) SSHAddKeys(username string, keys ...string) error {
	return kc.get().SSHAddKeys(username, keys...)
}

// MountHeadIndex returns the number and the overall size of files in a
// given remote directory.
func (kc *kiteClient) MountHeadIndex(path string) (string, int, int64, error) {
	return kc.get().MountHeadIndex(path)
}

// MountGetIndex returns an index that describes the current state of remote
// directory.
func (kc *kiteClient) MountGetIndex(path string) (*index.Index, error) {
	return kc.get().MountGetIndex(path)
}

// Exec runs a command on a remote machine.
func (kc *kiteClient) Exec(req *os.ExecRequest) (*os.ExecResponse, error) {
	return kc.get().Exec(req)
}

// Kill terminates previously started command on a remote machine.
func (kc *kiteClient) Kill(req *os.KillRequest) (*os.KillResponse, error) {
	return kc.get().Kill(req)
}

// Context returns client's Context.
func (kc *kiteClient) Context() context.Context {
	return kc.get().Context()
}

// Get gets klient from the pool, if an error occurs, disconnected client will
// be returned.
func (kc *kiteClient) get() Client {
	k, err := kc.pool.Get(kc.addr)
	if err != nil {
		return NewDisconnected(kc.ctx)
	}

	k.SetContext(kc.ctx)
	return k
}
