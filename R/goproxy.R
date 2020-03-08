#' Doubles an integer using go
#'
#' @useDynLib goproxy
#' @export
run_proxy <- function() {
  .Call("theproxy", PACKAGE = "goproxy")
}
