workflow "New workflow" {
  on = "push"
  resolves = ["Setup Go for use with actions"]
}

action "Setup Go for use with actions" {
  uses = "actions/setup-go@595aed780bc79b816735c25f8ed51a5fbc96753a"
}
