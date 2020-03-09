#' Run a reverse proxy directly
#' 
#' WARNING: this is currently not able to be interrupted. Use `run_proxy`
#' instead.
#' 
#' @param url The url (with protocol) to proxy
#' @param port The port on which to listen
#' 
#' @export
run_proxy_raw <- function(url, port = "9845") {
  warn_once("`run_proxy_raw` is currently not interruptible. Please use `run_proxy()` instead")
  
  .Call("theproxy", port, url, PACKAGE = "goproxy")
}

#' Run a reverse proxy
#' 
#' Starts a reverse proxy in a subprocess (for process cleanup),
#' but monitors the process for stdout / stderr output.
#' 
#' @param url The url (with protocol) to reverse proxy
#' @param port The port on which to listen
#' 
#' @export
run_proxy <- function(url, port = "9845", timeout = 20) {
  
  message(glue::glue("Starting GoProxy on http://127.0.0.1:{port}"))
  
  server_proc <- callr::r_bg(
    function(url, port) {options("goproxy_disable_warnings" = TRUE); goproxy::run_proxy_raw(url, port)},
    args = list(
      url = url,
      port = port
    )
  )
  
  # ensure process cleans up after itself
  on.exit(server_proc$kill_tree(), add = TRUE)
  
  # max 20 seconds wait
  ready <- FALSE
  message("Waiting for proxy to start...")
  for (i in 1:(timeout * 10)) {
    Sys.sleep(0.1)
    
    if (!server_proc$is_alive()) {
      stop(noop_print(server_proc$read_error_lines()))
    }
    
    try_to_connect <- try(httr::GET(glue::glue("http://127.0.0.1:{port}")), silent = TRUE)
    if (!inherits(try_to_connect, "try-error")) {
      ready <- TRUE
      break
    }
  }
  
  
  if (!ready) stop(glue::glue("ERROR: GoProxy did not start after {timeout} seconds"))
  
  message("Started! Opening browser")
  
  if (rstudioapi::isAvailable()) rstudioapi::viewer(glue::glue("http://127.0.0.1:{port}"))
  
  # poll and watch
  while(server_proc$is_alive()) {
    Sys.sleep(0.1)
    
    noop_print(server_proc$read_output_lines())
    noop_print(server_proc$read_error_lines())
  }
  invisible()
}
