%module goLib
%{
   #include "add.h"
%}
%insert(cgo_comment_typedefs)
%{
#cgo LDFLAGS: -L/. -ladd
#include "add.h"
%}
%include "add.h"