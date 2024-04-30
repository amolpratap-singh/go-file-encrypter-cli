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

#### 1. Encryption Flow

* The program begins by checking for the presence of the source file specified by the user. If the file exists, it proceeds to open it and extract the plain text content. It then generates an empty nonce and randomizes it for further processing.

* Following that, a password-based derivation function is employed to create a key from the provided password. This function incorporates parameters such as the password itself, byte length, SHA-1 hash, iteration count, and the randomized nonce.

* Once the key is derived, it is utilized alongside the AES Cipher (Advanced Encryption Standard) to encrypt the plain text data.

* The AES Cipher operates in Galois Counter Mode (GCM), which produces a tag appended to the ciphertext. This tag plays a critical role in ensuring data integrity during decryption.

* Using AES GCM encryption (aesgcm.SEAL), the program encrypts the plain text to generate the cipher text. Subsequently, it generates a new source file to store the encrypted data and writes the cipher text into it.

#### 2. Decryption Flow

* The decryption process begins by verifying the presence of the encrypted file provided by the user. If the file exists, the program opens it and reads the cipher text content.

* Subsequently, the program extracts the 12-byte nonce added to the encryption file at the end during encryption by reading the last 12 digits of the encrypted data.

* Next, a password-based derivation function is applied. This function incorporates parameters such as the password itself, byte length, SHA-1 hash, iteration, and nonce to derive a key.

* The derived key is then utilized with the AES Cipher (Advanced Encryption Standard) to decrypt the cipher block in Galois Counter Mode (GCM).

* The program then passes the cipher text file to the decryption function without the nonce, applying the AES GCM decryption (aesgcm.Open) to retrieve the plain text.

* Finally, it creates a new file for writing the decrypted data and proceeds to write the decrypted content to this newly created file.