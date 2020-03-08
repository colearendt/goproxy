#' Doubles an integer using go
#'
#' @useDynLib goproxy
#' @export
run_proxy <- function(port, url) {
  .Call("theproxy", port, url, PACKAGE = "goproxy")
}
