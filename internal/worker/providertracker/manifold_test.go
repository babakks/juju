// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package providertracker

import (
	"context"

	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/dependency"
	dependencytesting "github.com/juju/worker/v4/dependency/testing"
	"github.com/juju/worker/v4/workertest"
	gc "gopkg.in/check.v1"
	"gopkg.in/tomb.v2"

	"github.com/juju/juju/environs"
	"github.com/juju/juju/internal/servicefactory"
	storage "github.com/juju/juju/internal/storage"
)

type manifoldSuite struct {
	baseSuite
}

var _ = gc.Suite(&manifoldSuite{})

func (s *manifoldSuite) TestValidateConfig(c *gc.C) {
	defer s.setupMocks(c).Finish()

	cfg := s.getConfig()
	c.Check(cfg.Validate(), jc.ErrorIsNil)

	cfg = s.getConfig()
	cfg.ProviderServiceFactoryName = ""
	c.Check(cfg.Validate(), jc.ErrorIs, errors.NotValid)

	cfg = s.getConfig()
	cfg.NewEnviron = nil
	c.Check(cfg.Validate(), jc.ErrorIs, errors.NotValid)

	cfg = s.getConfig()
	cfg.NewWorker = nil
	c.Check(cfg.Validate(), jc.ErrorIs, errors.NotValid)

	cfg = s.getConfig()
	cfg.Logger = nil
	c.Check(cfg.Validate(), jc.ErrorIs, errors.NotValid)

	cfg = s.getConfig()
	cfg.GetProviderServiceFactory = nil
	c.Check(cfg.Validate(), jc.ErrorIs, errors.NotValid)
}

func (s *manifoldSuite) getConfig() ManifoldConfig {
	return ManifoldConfig{
		ProviderServiceFactoryName: "provider-service-factory",
		Logger:                     s.logger,
		NewEnviron: func(ctx context.Context, op environs.OpenParams) (environs.Environ, error) {
			return nil, nil
		},
		NewWorker: func(ctx context.Context, cfg Config) (worker.Worker, error) {
			return newStubWorker(), nil
		},
		GetProviderServiceFactory: func(getter dependency.Getter, name string) (ServiceFactory, error) {
			return s.serviceFactory, nil
		},
	}
}

func (s *manifoldSuite) newGetter() dependency.Getter {
	resources := map[string]any{
		"provider-service-factory": &stubProviderServiceFactory{},
	}
	return dependencytesting.StubGetter(resources)
}

var expectedInputs = []string{"provider-service-factory"}

func (s *manifoldSuite) TestInputs(c *gc.C) {
	c.Assert(Manifold(s.getConfig()).Inputs, jc.SameContents, expectedInputs)
}

func (s *manifoldSuite) TestStart(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.serviceFactory.EXPECT().Cloud().Return(s.cloudService)
	s.serviceFactory.EXPECT().Config().Return(s.configService)
	s.serviceFactory.EXPECT().Credential().Return(s.credentialService)
	s.serviceFactory.EXPECT().Model().Return(s.modelService)

	w, err := Manifold(s.getConfig()).Start(context.Background(), s.newGetter())
	c.Assert(err, jc.ErrorIsNil)
	workertest.CleanKill(c, w)
}

func (s *manifoldSuite) TestOutput(c *gc.C) {
	defer s.setupMocks(c).Finish()

	w := &trackerWorker{
		environ: s.environ,
	}

	var environ environs.Environ
	err := manifoldOutput(w, &environ)
	c.Check(err, jc.ErrorIsNil)

	var destroyer environs.CloudDestroyer
	err = manifoldOutput(w, &destroyer)
	c.Check(err, jc.ErrorIsNil)

	var registry storage.ProviderRegistry
	err = manifoldOutput(w, &registry)
	c.Check(err, jc.ErrorIsNil)

	var bob string
	err = manifoldOutput(w, &bob)
	c.Check(err, gc.ErrorMatches, `expected \*environs.Environ, \*storage.ProviderRegistry, or \*environs.CloudDestroyer, got \*string`)
}

type stubWorker struct {
	tomb tomb.Tomb
}

func newStubWorker() *stubWorker {
	w := &stubWorker{}
	w.tomb.Go(func() error {
		<-w.tomb.Dying()
		return nil
	})
	return w
}

func (w *stubWorker) Kill() {
	w.tomb.Kill(nil)
}

func (w *stubWorker) Wait() error {
	return w.tomb.Wait()
}

type stubProviderServiceFactory struct {
	servicefactory.ProviderServiceFactory
}
