// See README on how to create a certificate value
variable "vault_value" {
  description = "Value of the certificate"
  type = string
  default     = "your certificate here" # Replace with your actual certificate value
}

// See README on how to create a certificate archive value
variable "vault_certificate_archive" {
  description = "Value of the certificate archive"
  type = string
  default     = "your certificate archive here" # Replace with your actual certificate value
}

// See README on how to create a file data value
variable "vault_file_data" {
  description = "Value of the file data"
  type = string
  default     = "your file data here" # Replace with your actual file data value
}