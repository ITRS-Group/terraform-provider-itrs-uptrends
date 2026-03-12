
variable "operator_password" {
  description = "Value of the password for the provider"
  type = string
  ephemeral = true # This variable is ephemeral and not stored in state.
  default     = "password"
}
