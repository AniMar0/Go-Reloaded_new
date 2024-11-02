# 🚀 Text Completion and Auto-Correction Tool



## 🌟 Objectives


This project aims to create a **powerful tool** for text completion, editing, and auto-correction.
The tool will help users automate tedious text modifications.

 ## 🔥 Features

<img src="https://raw.githubusercontent.com/MicaelliMedeiros/micaellimedeiros/master/image/computer-illustration.png" alt="Computer Illustration" min-width="400px" max-width="400px" width="400px" align="right">
 
 
The tool accepts two arguments:
1. **Input file**: A file with text that needs modifications.
2. **Output file**: A file where the modified text will be saved.

> ⚠️ Make sure your input text is properly formatted for the tool to work its magic.

Here's what the tool does:

### 1. Hexadecimal to Decimal Conversion
Converts `(hex)` numbers to decimal.
- **Example**: `"1E (hex)"` becomes `"30"`

### 2. Binary to Decimal Conversion
Converts `(bin)` numbers to decimal.
- **Example**: `"10 (bin)"` becomes `"2"`

### 3. Uppercase, Lowercase, Capitalization
Modifies text case based on instructions like `(up)`, `(low)`, and `(cap)`. You can also specify multiple words to convert using a number.
- **Example**: `"go (up)"` -> `"GO"`
- **Example with multiple words**: `"This is exciting (up, 2)"` -> `"THIS IS exciting"`

### 4. Punctuation Handling
Ensures punctuation is correctly placed:
- **Example**: `"Hello , world !"` becomes `"Hello, world!"`

Special cases for `...` or `!?` are preserved as needed.
  
### 5. Single Quotes Handling
Ensures words or phrases between single quotes are correctly formatted.
- **Example**: `"I am ' amazing '"` becomes `"I am 'amazing'"`

### 6. Article Correction
Changes "a" to "an" when followed by a word starting with a vowel or "h".
- **Example**: `"A amazing day"` -> `"An amazing day"`

---

## 🛠 Usage

Run the tool using:

```bash
cd cmd
```

```bash
go run main.go input.txt output.txt
```

## 🧪 Testing
It's essential to test your code thoroughly. Here's how you can run the unit tests:

bash
Copy code

```bash
go test
```
> 💡 Unit testing ensures your tool works as expected under various scenarios. It’s highly recommended to write tests for every major feature.

---

## 💻 Tech Stack

**🦄 Languages**:   Go (Golang), Markdown<br> 
**💼 Tools**:   VSCode, Go testing framework, Git<br> 
**📦 Packages**:   strconv, strings<br>

## 👨‍💻 About Me
<p align="left"> I'm a passionate developer who enjoys building text processing tools and solving complex problems through code. Currently, I'm honing my Go skills and working on various backend projects. I also love collaborating on open-source projects. </p> <p align="left"> 💌 Feel free to reach out: ⤵️ </p> 
<p align="left"> 
  <a href="mailto:mr.zakariakahlaoui@gmail.com" title="Gmail"> <img src="https://img.shields.io/badge/-Gmail-FF0000?style=flat-square&labelColor=FF0000&logo=gmail&logoColor=white&link=YOUR-GMAIL-LINK" alt="Gmail"/></a> 
  <a href="https://www.linkedin.com/in/zakaria-kahlaoui/" title="LinkedIn"> <img src="https://img.shields.io/badge/-Linkedin-0e76a8?style=flat-square&logo=Linkedin&logoColor=white&link=YOUR-LINKEDIN-LINK" alt="LinkedIn"/></a></a> 
  <a href="https://www.instagram.com/zakariia__48/" title="Instagram"> <img src="https://img.shields.io/badge/-Instagram-DF0174?style=flat-square&labelColor=DF0174&logo=instagram&logoColor=white&link=YOUR-INSTAGRAM-LINK" alt="Instagram"/></a> </p>

## 📜 License
This project is licensed under the MIT License.

## 🌟 Contributing
Want to help improve this tool? Contributions are welcome! Feel free to fork the repository and submit a pull request. Ensure your code follows good Go practices and is well-tested.

## 📝 Acknowledgements
Special thanks to my peers and @kenji20thh for their insights and reviews. Thanks to the Go community for providing excellent documentation and support.
