// Copyright 2021 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// System: FreeBSD

package sysutil

import "github.com/tredoe/osutil/executil"

const (
	filePkg = "pkg"
	pathPkg = "/usr/sbin/pkg"
)

// ManagerPkg is the interface to handle the FreeBSD package manager,
// called 'package' or 'pkg'.
type ManagerPkg struct {
	pathExec string
	cmd      *executil.Command
}

// NewManagerPkg returns the Pkg package manager.
func NewManagerPkg() ManagerPkg {
	return ManagerPkg{
		pathExec: pathPkg,
		cmd: excmd.Command("", "").
			BadExitCodes([]int{1}),
	}
}

func (m ManagerPkg) setExecPath(p string) { m.pathExec = p }

func (m ManagerPkg) ExecPath() string { return m.pathExec }

func (m ManagerPkg) PackageType() string { return Pkg.String() }

func (m ManagerPkg) Install(name ...string) error {
	args := []string{pathPkg, "install", "-y"}

	_, err := m.cmd.Command(sudo, append(args, name...)...).Run()
	return err
}

func (m ManagerPkg) Remove(name ...string) error {
	args := []string{pathPkg, "delete", "-y"}

	_, err := m.cmd.Command(sudo, append(args, name...)...).Run()
	return err
}

func (m ManagerPkg) Purge(name ...string) error {
	return m.Remove(name...)
}

func (m ManagerPkg) Update() error {
	_, err := m.cmd.Command(sudo, pathPkg, "update").Run()
	return err
}

func (m ManagerPkg) Upgrade() error {
	_, err := m.cmd.Command(sudo, pathPkg, "upgrade", "-y").Run()
	return err
}

func (m ManagerPkg) Clean() error {
	_, err := m.cmd.Command(sudo, pathPkg, "autoremove", "-y").Run()
	if err != nil {
		return err
	}
	_, err = m.cmd.Command(sudo, pathPkg, "clean", "-y").Run()
	return err
}

func (m ManagerPkg) AddRepo(alias string, url ...string) error {
	panic("unimplemented")
}

func (m ManagerPkg) RemoveRepo(r string) error {
	panic("unimplemented")
}
