# go-commit

Conventional Commits from the terminal - written in Go.

## Motivation

-   I've been using [VS Code Conventional Commits][vscode_conventional_commits] for sometime and I wanted to have that experience on my terminal as well.

    Fun fact: I wrote this `go` version before rewriting this as [commit.sh][commit_sh] (shell script). I'm just moving this to my public repository.

### Note

-   This utility covers just the happy path ([commit.sh][commit_sh] covers a few more scenarios than this one)
-   This does not -

    -   use any external configuration file (The change list and emoji list are hardcoded)
    -   have any fuzzy search option
    -   clear the terminal after every choice selection

## Install instructions

-   Clone the repository

    ```sh
    git clone https://github.com/krish-r/go-commit.git
    ```

-   Build the binary

    ```sh
    go build -o bin/go-commit .
    ```

-   Copy it somewhere in $PATH, and make it executable. (for ex. `~/.local/bin`)

    ```sh
    cp ./bin/go-commit ~/.local/bin/go-commit && chmod u+x ~/.local/bin/go-commit
    ```

## Usage

-   Run the binary inside a git repository

    ```sh
    go-commit
    ```

## Uninstall instructions

-   Simply remove the binary from your path.

    ```sh
    rm -i $(which go-commit)
    ```

## Credits

-   [VS Code Conventional Commits][vscode_conventional_commits] - for the inspiration & description
-   [gitmoji][gitmoji] - for the emoji guides & description

## Demo

![Demo][demo_gif]

[commit_sh]: https://github.com/krish-r/commit.sh
[demo_gif]: https://user-images.githubusercontent.com/54745129/209949329-a48bd230-85e4-4983-9080-a122cb3937be.gif
[vscode_conventional_commits]: https://github.com/vivaxy/vscode-conventional-commits
[gitmoji]: https://github.com/carloscuesta/gitmoji
