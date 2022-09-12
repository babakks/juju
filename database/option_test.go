// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package database

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"

	"github.com/juju/juju/agent"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
)

type optionSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&optionSuite{})

func (s *optionSuite) TestEnsureDataDir(c *gc.C) {
	subDir := strconv.Itoa(rand.Intn(10))

	cfg := dummyAgentConfig{dataDir: "/tmp/" + subDir}

	f := NewOptionFactory(cfg, dqlitePort, nil)

	expected := fmt.Sprintf("/tmp/%s/%s", subDir, dqliteDataDir)
	s.AddCleanup(func(*gc.C) { _ = os.RemoveAll(expected) })

	// Call twice to check both the creation and extant scenarios.
	_, err := f.EnsureDataDir()
	c.Assert(err, jc.ErrorIsNil)

	dir, err := f.EnsureDataDir()
	c.Assert(err, jc.ErrorIsNil)
	c.Check(dir, gc.Equals, expected)

	_, err = os.Stat(expected)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *optionSuite) TestWithAddressOption(c *gc.C) {
	f := NewOptionFactory(nil, dqlitePort, func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPAddr{IP: net.ParseIP("10.0.0.5")},
			&net.IPAddr{IP: net.ParseIP("127.0.0.1")},
		}, nil
	})

	// We can not actually test the realisation of this option,
	// as the options type from the go-dqlite is not exported.
	// We are also unable to test by creating a new dqlite app,
	// because it fails to bind to the contrived address.
	// The best we can do is check that we selected an address
	// based on the absence of an error.
	_, err := f.WithAddressOption()
	c.Assert(err, jc.ErrorIsNil)
}

type dummyAgentConfig struct {
	agent.Config
	dataDir string
}

// DataDir implements agent.AgentConfig.
func (cfg dummyAgentConfig) DataDir() string {
	return cfg.dataDir
}
