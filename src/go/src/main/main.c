#include "_cgo_export.h"
#include <R.h>
#include <Rinternals.h>

SEXP theproxy() {
  return Rf_ScalarInteger( runProxy() );
}
