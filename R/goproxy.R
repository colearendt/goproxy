#' Run a reverse proxy
#' 
#' @param url The url (with protocol) to proxy
#' @param port The port on which to listen
#'
#' @useDynLib goproxy
#' @export
run_proxy <- function(url, port) {
  .Call("theproxy", port, url, PACKAGE = "goproxy")
}
