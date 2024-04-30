## Encrypter CLI Tool
### About
The Encrypter CLI tool, developed in Golang, enables file encryption and decryption using symmetric encryption techniques. For encryption or decryption operations, the tool utilizes the crypto package, a built-in package in Go.

![Diagram](utils/Encrypter_cli.jpg)

### Usage

#### Encryption of file

```sh
# Encrypts a file via CLI tool
./EncrypterCLI encrypt sample_file.txt
```

or

```sh
# Encrypts a file via go run cmd
go run main.go encrypt sample_file.txt
```

#### Decryption of file

```sh
# Decrypts a file via CLI tool
./EncrypterCLI decrypt sample_file.txt
```

or

```sh
# Decrypts a file via go run cmd
go run main.go decrypt sample_file.txt
```

### Implementation Workflow

#### 1. Encryption

check_for_source_file.txt => open the source file => read the plain text from source file

create empty nonce => randomize the nonce => 

password based derivation function (password to encrypt the file, byte length, SHA-1, iterations,randomize the nonce )

A function in which is a key is generated from the password, can be used as an encryption key or as a hash value.

derived key => AES Cipher (Advanced encryption standard) => cipher block => galois counter mode (GCM generates a tag which is appended to the ciphertext. This tag is used to verify the integerity of the data upon decryption. 
GCM uses AES to encrypt the plain text data. It operates in a counter mode and generates a stream of encrypted blocks.)

plain text from the file => aesgcm.SEAL => cipher text => create a source file for encrypted data => write cipher text to new source file 



