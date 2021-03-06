// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package engine

import (
	"fmt"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"runtime/debug"
)

type Prereqs struct {
	TemporarySession bool
	Device           bool
}

type Engine2 interface {
	Run(libkb.MetaContext) error
	Prereqs() Prereqs
	libkb.UIConsumer
}

type UIDelegateWanter interface {
	WantDelegate(libkb.UIKind) bool
}

func requiresUI(c libkb.UIConsumer, kind libkb.UIKind) bool {
	for _, ui := range c.RequiredUIs() {
		if ui == kind {
			return true
		}
	}
	for _, sub := range c.SubConsumers() {
		if requiresUI(sub, kind) {
			return true
		}
	}
	return false
}

func isLoggedInWithUIDAndError(m libkb.MetaContext) (ret bool, uid keybase1.UID, err error) {
	ret, uid, err = libkb.BootstrapActiveDeviceWithMetaContext(m)
	return ret, uid, err
}
func isLoggedIn(m libkb.MetaContext) (ret bool, uid keybase1.UID) {
	ret, uid, _ = libkb.BootstrapActiveDeviceWithMetaContext(m)
	return ret, uid
}

func isLoggedInWithError(m libkb.MetaContext) (ret bool, err error) {
	ret, _, err = isLoggedInWithUIDAndError(m)
	return ret, err
}

func runPrereqs(m libkb.MetaContext, e Engine2) error {
	prq := e.Prereqs()

	if prq.TemporarySession {
		err := m.G().AssertTemporarySession(m.LoginContext())
		if err != nil {
			return err
		}
	}

	if prq.Device {
		ok, err := isLoggedInWithError(m)
		if err != nil {
			return err
		}
		if !ok {
			return libkb.DeviceRequiredError{}
		}
	}

	return nil

}

func RunEngine2(m libkb.MetaContext, e Engine2) (err error) {
	defer m.CTrace(fmt.Sprintf("RunEngine(%s)", e.Name()), func() error { return err })()

	if m, err = delegateUIs(m, e); err != nil {
		return err
	}
	if err = check(m, e); err != nil {
		return err
	}
	if err = runPrereqs(m, e); err != nil {
		return err
	}

	err = e.Run(m)
	return err
}

func delegateUIs(m libkb.MetaContext, e Engine2) (libkb.MetaContext, error) {
	if m.G().UIRouter == nil {
		return m, nil
	}

	// currently, only doing this for SecretUI, but in future,
	// perhaps should iterate over all registered UIs in UIRouter.
	if requiresUI(e, libkb.SecretUIKind) {
		sessionID := m.UIs().SessionID
		if ui, err := m.G().UIRouter.GetSecretUI(sessionID); err != nil {
			return m, err
		} else if ui != nil {
			m.CDebugf("using delegated secret UI for engine %q (session id = %d)", e.Name(), sessionID)
			m = m.WithSecretUI(ui)
		}
	}

	if wantsDelegateUI(e, libkb.IdentifyUIKind) {
		m.CDebugf("IdentifyUI wanted for engine %q", e.Name())
		if ui, err := m.G().UIRouter.GetIdentifyUI(); err != nil {
			return m, err
		} else if ui != nil {
			m.CDebugf("using delegated identify UI for engine %q", e.Name())
			m = m.WithDelegatedIdentifyUI(ui)
		}
	}

	return m, nil
}

func wantsDelegateUI(e Engine2, kind libkb.UIKind) bool {
	if !requiresUI(e, kind) {
		return false
	}
	if i, ok := e.(UIDelegateWanter); ok {
		return i.WantDelegate(kind)
	}
	return false
}

func check(m libkb.MetaContext, c libkb.UIConsumer) error {
	if err := checkUI(m, c); err != nil {
		return err
	}

	for _, sub := range c.SubConsumers() {
		if err := check(m, sub); err != nil {
			if _, ok := err.(CheckError); ok {
				return err
			}
			return CheckError{fmt.Sprintf("%s: %s", sub.Name(), err)}
		}
	}

	return nil
}

func checkUI(m libkb.MetaContext, c libkb.UIConsumer) error {
	for _, ui := range c.RequiredUIs() {
		if !m.UIs().HasUI(ui) {
			return CheckError{fmt.Sprintf("%s: requires ui %q\n\n%s", c.Name(), ui, string(debug.Stack()))}
		}
	}
	return nil
}
