# ASCII-TERMINAL

---

## **Description**

ASCII-TERMINAL is a command-line tool that converts user-provided text into visually appealing ASCII ART. It offers multiple banner styles and text formatting options, allowing users to create and manipulate ASCII text efficiently. The tool supports alignment, colorization, and reverse text functionality, making it a versatile choice for terminal-based ASCII ART generation.

---

## **Features**

- **Multiple Banner Styles**: Supports `standard`, `shadow`, and `thinkertoy` ASCII fonts.
- **Text Alignment**: Align text to the left, right, center, or justify it.
- **Color Customization**: Apply various colors to ASCII text output.
- **Reverse Art**: Convert ASCII text back into readable characters.
- **File Exporting**: Save ASCII output to a text file for future use.
- **Terminal-Friendly**: Optimized for terminal-based interaction with efficient processing.

---

## **Usage**

### **How to Run**

1. Clone the repository:
   ```bash
   git clone https://github.com/AliHJMM/ASCII-TERMINAL
   ```
2. Navigate to the project directory:
   ```bash
   cd ASCII-TERMINAL
   ```
3. Run the application:
   ```bash
   go run main.go "Hello, World!"
   ```

### **Command Options**

#### **Basic Usage**

```bash
 go run main.go "Hello, World!"
```

#### **Apply Banner Styles**

```bash
 go run main.go "Hello" standard
 go run main.go "Hello" shadow
 go run main.go "Hello" thinkertoy
```

#### **Colorize Text**

```bash
 go run main.go --color=red "Hello"
 go run main.go --color=blue "ASCII Art"
```

#### **Align Text**

```bash
 go run main.go --align=center "Hello"
 go run main.go --align=right "ASCII"
 go run main.go --align=justify "Hello there!"
```

#### **Reverse ASCII Art**

```bash
 go run main.go --reverse=AuditExamples/example01.txt
```

#### **Save Output to File**

```bash
 go run main.go --output=output.txt "Hello, World!"
```

---

## **Technical Implementation**

- **Go-based Processing**: The backend is built in Go, ensuring fast and reliable execution.
- **ASCII Banners**: ASCII templates are stored as text files and processed dynamically.
- **Custom Formatting**: Implements text alignment, coloring, and reverse ASCII recognition.
- **Efficient File Handling**: Users can save and retrieve ASCII text outputs from files.

---

## **Project Structure**

```
├── ascii-terminal/
│   ├── AuditExamples/    # Sample ASCII art files
│   ├── Banners/          # ASCII font styles
│   ├── functions/        # Go functions for processing text
│   ├── test/             # Test scripts
│   ├── main.go           # Main entry point
│   ├── README.md         # Documentation
```

---

## **Authors**

- [Ali Hasan](https://github.com/AliHJMM)
- [Habib Mansoor](https://github.com/7abib04)
