package reno

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// readFile reads a file and streams its content into a buffer
// The return data is the contents in a slice of bytes
//
// Caller:
//
//	TrackFiles() -> []FileHash
func readFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("An error occured in reading file")
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 1024) // read the file into buffer 1024 bytes at a time
	var fileContent []byte
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}
		fileContent = append(fileContent, buffer[:n]...)
	}

	return fileContent
}

// hashContent takes the file content provided by func readFile() and hashes it
// After hashing it, it gets converted to hexCode using encoding/hex and returns it as string
func hashContent(fileContent []byte) string {
	hasher := sha1.New()
	hasher.Write(fileContent) // hashing the file content
	hashedContent := hasher.Sum(nil)

	hexedHash := hex.EncodeToString(hashedContent) // converts sha1Hash to hex for readability
	return hexedHash
}

// FileHash is mainly for debugging purposes and to represent a higher level view of the data passed aroud
// Users:
//
//	TrackFiles() -> []FileHash
//	ComputeFinalHash(fileHashes []FileHash) -> string
type FileHash struct {
	fileName string
	fileHash string
}

// TrackFiles uses a  stdlib, filePath.WalkDir that recursively gets all the files residing in views and locales folders.
//
// It calls on the func readFile() and func hashContent() and stores it inside the FileHash type definition.
//
// It then returns a slice/list of all these files and their hashes ready to be used by func ComputeFinalHash()
//
// Calls:
//
//	filePath.WalkDir(root string, fn fs.WalkDirFunc) -> error
//	readFile(filePath) -> []bytes
//	hashContent(fileContent) -> string
//
// Callers:
//
//	main()
func TrackFiles() []FileHash {
	var hashes = []FileHash{}

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {

		if !d.IsDir() && (strings.Contains(path, "views") || strings.Contains(path, "locales")) {
			fileContent := readFile(path)
			hexedHash := hashContent(fileContent)
			fileHash := FileHash{
				fileName: path,
				fileHash: hexedHash,
			}

			hashes = append(hashes, fileHash)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return hashes
}

// ComputeFinalHash takes all filehashes for the views and locales folder content and computes a combined hash.
//
// This is to make sure we are relying on state and content of these files at a particular instance in time and to
//
// increase entropy and randomness. This way, new files and deleted files are also taken care of.
//
// Calls:
//
//	hashContent(content []byte) -> string
//
// Callers:
//
//	main()
//
// Returns:
//
//	The final hash, hexEncoded as a string
func ComputeFinalHash(fileHashes []FileHash) string {
	// Here, we are just going to extract the hashes from the each item in the FileHash{}
	var hashes = []string{}
	for _, file := range fileHashes {
		hashes = append(hashes, file.fileHash)
	}

	combinedHash := strings.Join(hashes, "")
	finalHash := hashContent([]byte(combinedHash))

	return finalHash
}
