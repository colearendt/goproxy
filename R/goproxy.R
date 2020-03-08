#' Doubles an integer using go
#'
#' @useDynLib goproxy
#' @export
run_proxy <- function(port) {
  .Call("theproxy", port, PACKAGE = "goproxy")
}
