//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// This file is machine generated by the command:
//   mksyscall.pl serial_darwin.go
// The generated stub is modified to make it compile under the "serial" package

package serial // import "github.com/ds-forks/bugst-go-serial"

import "golang.org/x/sys/unix"

func ioctl(fd int, req uint64, data uintptr) (err error) {
	_, _, e1 := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(data))
	if e1 != 0 {
		err = e1
	}
	return
}
