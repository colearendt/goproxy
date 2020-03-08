#include "_cgo_export.h"
#include <R.h>
#include <Rinternals.h>
#include <stdlib.h>
#include <string.h>

//#define SHORT_VEC_LENGTH(xx) (((VECSXP) (xx))->vecsxp.length)

SEXP theproxy(SEXP port, SEXP url) {
  SEXP sport = STRING_ELT(port, 0);
  GoString gosport = { (char*) CHAR(sport), Rf_xlength(sport) };
  
  SEXP surl = STRING_ELT(url, 0);
  GoString gosurl = { (char*) CHAR(surl), Rf_xlength(surl) };
  return Rf_ScalarInteger( runProxy(gosport, gosurl) );
}
