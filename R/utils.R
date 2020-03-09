
noop_print <- function(input){
  if (length(input) > 0) {
    message(paste(input, collapse = "\n"))
  }
}

scoped_silence <- function(frame = rlang::caller_env()) {
  rlang::scoped_options(
    .frame = frame,
    goproxy_disable_warnings = TRUE
  )
}

warn_once <- function(msg, id = msg) {
  if (rlang::is_true(rlang::peek_option("goproxy_disable_warnings"))) {
    return(invisible(NULL))
  }
  
  if (rlang::env_has(warn_env, id)) {
    return(invisible(NULL))
  }
  
  has_color <- function() rlang::is_installed("crayon") && crayon::has_color()
  silver <- function(x) if (has_color()) crayon::silver(x) else x
  
  msg <- paste0(
    msg,
    "\n",
    silver("This warning is displayed once per session.")
  )
  
  rlang::env_poke(warn_env, id, TRUE)
  
  rlang::signal(msg, .subclass = "warning")
}
warn_env <- new.env(parent = emptyenv())
