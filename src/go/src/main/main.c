#include "_cgo_export.h"
#include <R.h>
#include <Rinternals.h>
#include <stdlib.h>
#include <string.h>

//#define SHORT_VEC_LENGTH(xx) (((VECSXP) (xx))->vecsxp.length)

SEXP theproxy(SEXP port) {
  SEXP sx = STRING_ELT(port, 0);
  
  GoString h = { (char*) CHAR(sx), Rf_xlength(sx) };
  return Rf_ScalarInteger( runProxy(h) );
}
