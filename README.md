# go_reloaded

![image](https://kruschecompany.com/wp-content/uploads/2020/08/Cover-image-for-blog-post-introducing-the-programming-language-Golang.png)

# `Text Manipulation`

This go program has full functionality in manipulating text in different ways. Some of them include; to convert numbers from hexadecimal or binary to decimal, change the case of words, or handle punctuation formatting, this tool has got you covered.

## Features

- [**Case Modification**](https://learn.zone01kisumu.ke/git/mdudi/go-reloaded/src/branch/master/applytranformations.go): Convert words to uppercase, lowercase, or capitalize them.
- [**Number Conversion**](https://learn.zone01kisumu.ke/git/mdudi/go-reloaded/src/branch/master/hextoint.go): Convert hexadecimal or binary numbers to decimal.
- [**Dynamic Modifications**](https://learn.zone01kisumu.ke/git/mdudi/go-reloaded/src/branch/master/applytranformations.go): Support for modifying a specified number of words before the command.
- [**Punctuation Formatting**](https://learn.zone01kisumu.ke/git/mdudi/go-reloaded/src/branch/master/punctuations.go): Handles punctuation placement according to specified rules.
- [**A/An Conversion**](https://learn.zone01kisumu.ke/git/mdudi/go-reloaded/src/branch/master/changearticles.go): Automatically changes "a" to "an" based on the following word startin in a lowercase vowel or `h`.

## Usage

```bash
$ go run . sample.txt result.txt
```
- Some of the command options include
 - `(up)`: Convert the preceding word to uppercase.
 - `(low)`: Convert the preceding word to lowercase.
 - `(cap)`: Capitalize the preceding word.
 - `(hex)`: Convert the preceding hexadecimal number to decimal.
 - `(bin)`: Convert the preceding binary number to decimal.
 - `(up, n)`: Convert the preceding 'n' words to uppercase.
 - `(low, n)`: Convert the preceding 'n' words to lowercase.
 - `(cap, n)`: Capitalize the preceding 'n' words.
Also handling punctuations

## Examples

```bash
Input: "I have 0x10 (hex) apples."
Output: "I have 16 apples."

Input: "This is so exciting (up, 2)!!"
Output: "This is SO EXCITING!!"

Input: "I was sitting over there ,and then BAMM !!"
Output: "I was sitting over there, and then BAMM!!"

Input: "' It is a beautiful day! '"
Output: "It is a beautiful day!"

Input: "There it is. A unicorn!"
Output: "There it is. An unicorn!"
```

## Getting Started

Thank you for your interest in contributing to our project! Follow the steps below to clone the repository and start working on it:

### Prerequisites

Make sure you have the following prerequisites installed on your system:

- Git: [Download & Install Git](https://git-scm.com/downloads)
- Go Programming Language: [Download & Install Go](https://golang.org/doc/install)

### Clone the Repository

1. Open your terminal or command prompt.
2. Navigate to the directory where you want to store the project.
3. Run the following command to clone the repository:

```bash
   git clone https://github.com/your-username/your-project.git
```
## Install Dependencies

 1. Navigate into the cloned repository directory:
```bash
cd your-project
```
 2. If your project has any dependencies, install them using `go get` or your preferred package manage

## Make Changes

- Now you're ready to make changes to the project:

    1. Open the project directory in your preferred code editor.
    2. Make the necessary changes or additions to the codebase.
## Test

- Before committing your changes, make sure to test the project to ensure everything is working as expected:
```bash
go test ./...
```
This command will run all the tests in your project.
## Commit Changes

Once you're satisfied with your changes, commit them to your local repository:
```bash
git add .
git commit -m "Your descriptive commit message here"
```
## Push Changes

If you're ready to contribute your changes, push them to your forked repository:
```bash
git push origin main
```
Replace main with the name of the branch you're working on.
## Create a Pull Request

Finally, create a pull request to merge your changes into the main project repository:

 1. Visit the repository on GitHub.
 2. Click on the "Pull Requests" tab.
 3. Click on the "New Pull Request" button.
 4. Follow the prompts to create a new pull request, providing a descriptive title and description for your changes.

That's it! Your pull request will be reviewed by me, and once approved,I will merge your changes into the main project.


## Contribution Guidlines

Your contribution to this project is greatly valued! Here's how you can contribute:

- `Bug Fixes`: If you notice any bugs or unexpected behavior, please open an issue or submit a pull request with your fix.

- `Feature Requests`: Have an idea for a new feature or enhancement? Share it with us by opening an issue.

- `Documentation`: Improvements to the documentation, including this README, are always welcome.

- `Testing`: Help us ensure the reliability of the program by writing and running tests.
- `Code Optimization`: If you have ideas for optimizing the codebase, feel free to contribute your improvements.

Feel free to reach out to us if you have any questions or need assistance with the contribution process.

## Used Packages

- `os`: Package for operating system functionality.
- `strconv`: Package for converting strings to other types.
- `strings`: Package for manipulating strings.
- `fmt`: Package for formatted I/O operations.
Let's make this project even better together! 🚀🎉🔧