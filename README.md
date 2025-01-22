# File Organizer

A simple Go tool to organize files in a directory by categorizing them into subfolders based on file types.

## How It Works

1. **Input**: Provide a directory path as a command-line argument.
2. **Categorize**: Files are grouped based on extensions:
   - **Images**: `.jpg`, `.jpeg`, `.png`, `.gif`
   - **Videos**: `.mov`, `.mkv`, `.mp4`
   - **Documents**: `.txt`, `.doc`, `.pdf`
   - **Others**: Files with unrecognized extensions.
3. **Folders**: Subfolders are created for each category.
4. **Move**: Files are moved into their respective folders.

## Usage

1. Save the code to `main.go`.
2. Run:
   ```bash
   go run main.go [Dir Path]
