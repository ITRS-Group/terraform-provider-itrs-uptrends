# How to create Certificate, Certificate Archive and File Data values

## Certificate value for attribute: value

Open Git Bash and run the following commands one by one:

1. Generate a new RSA key and self-signed certificate:
    ```bash
    openssl req -newkey rsa:2048 -nodes -keyout key.key -x509 -days 365 -out certificate.pem
    ```

2. Convert the certificate to Base64:
    ```bash
    base64 -w 0 certificate.pem > certificate.b64
    ```

3. Display the Base64 encoded certificate:
    ```bash
    cat certificate.b64
    ```

4. Copy the output value and paste it into your variable in variables.tf file:
    ```hcl
    // filepath: [variables.tf](http://_vscodecontentref_/0)
    variable "vault_value" {
      description = "Value of the certificate"
      type        = string
      default     = "paste value here"
    }
    ```

## Certificate Archive value for attribute: certificate_archive.archive_data 

Open Git Bash and run the following commands one by one:

1. Generate a new RSA private key:
    ```bash
    openssl genrsa -out myprivate.key 2048
    ```

2. Generate a self-signed certificate using the newly created key:
    ```bash
    openssl req -new -x509 -key myprivate.key -out mycertificate.crt -days 365
    ```

3. Create a PKCS#12 archive using the private key and the certificate:
    > You will be asked to enter a password. The password you have to remember for the certificate_archive.password

    ```bash
    openssl pkcs12 -export -out myarchive.p12 -inkey myprivate.key -in mycertificate.crt
    ```

4. Convert the archive to Base64:
    ```bash
    base64 -w 0 myarchive.p12 > archive.b64
    ```

5. Display the Base64 encoded archive:
    ```bash
    cat archive.b64
    ```

6. Copy the output value and paste it into your variable in variables.tf file:
    ```hcl
    // filepath: [variables.tf](http://_vscodecontentref_/0)
    variable "vault_archive_data" {
      description = "Value of the certificate"
      type        = string
      default     = "paste value here"
    }
    ```

## File Data value for attribute: file.data

Open Git Bash and run the following commands one by one:

1. Navigate to the location of the file you want to upload:
    ```bash
    cd folder_name
    ```

2. Convert the file to base 64:
    ```bash
    base64 -w 0 your_file.txt > file.b64
    ```

3. Display the Base64 encoded archive:
    ```bash
    cat file.b64
    ```

4. Copy the output value and paste it into your variable in variables.tf file:
    ```hcl
    // filepath: [variables.tf](http://_vscodecontentref_/0)
    variable "vault_file_data" {
      description = "Value of the file data"
      type = string
      default     = "paste value here"
    }
    ```