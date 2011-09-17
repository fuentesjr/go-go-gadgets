# Copyright 2009 The Go Authors. All rights reserved. 1
# Use of this source code is governed by a BSD-style 2
# license that can be found in the LICENSE file. 3

include $(GOROOT)/src/Make.inc
TARG=<binary-name>
GOFILES=\
  <src-filename>.go\

include $(GOROOT)/src/Make.cmd
