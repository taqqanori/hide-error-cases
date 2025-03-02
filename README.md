# Hide Error Cases (for Go)

**Tired of seeing `if err != nil { return err }`? This may help!**

When you read Go codes, you may want to concentrate on normal cases rather than error cases (`if err!= nil { return err }`s).
This extension may help you in such situation, by folding error case codes or/and making error case codes transparent.

Ideally, this feature should be officially provided by [vscode-go](https://github.com/golang/vscode-go), so, if you like my extension, let's up-vote [this issue (golang/vscode-go#2311)](https://github.com/golang/vscode-go/issues/2311)!

![Hide Error Cases Screenshot](images/screen-shot.png)

This extension recognizes codes fulfilling followings as error case codes.

- if, else-if, else blocks
- Ending with return statement
- Returning a object which is not static `nil` as a type that matches regexp defined in `go.hideErrorCases.errorTypeRegexp`

## Features

### Fold Error Cases

`Ctrl+Shift+P` and find & execute "Fold Error Cases" command, then error case codes in .go file opened in current editor will be folded.
To unfold error case codes, `Ctrl+Shift+P` and execute "Unfold All" command (built-in command of VSCode).
You can automatically fold error case codes every time you open .go files by configuring `go.hideErrorCases.autoFold` (`Extensions` > `Hide Error Cases (Go)` > `Auto Fold` on Settings UI).

### Make Error Cases Transparent

`Ctrl+Shift+P` and find & execute "Make Error Cases Transparent" command, then error case codes in .go file opened in current editor will be transparent.
You can automatically make error case codes transparent every time you open .go files by configuring `go.hideErrorCases.autoMakeTransparent` (`Extensions` > `Hide Error Cases (Go)` > `Auto Make Transparent` on Settings UI).

### Auto-Fold/Auto-Make-transparent on save

You can perform auto-fold/auto-make-transparent feature above every time you **save** .go files by configuring `go.hideErrorCases.hideOnSave` (`Extensions` > `Hide Error Cases (Go)` > `Hide On Save` on Settings UI).

## Requirements

- `go` command is available on `PATH`.
- `go` language extension is installed in VSCode.

## Extension Settings

This extension contributes the following settings:

- `go.hideErrorCases.autoFold`: enable/disable auto-fold feature
- `go.hideErrorCases.autoMakeTransparent`: enable/disable auto-make-transparent feature
- `go.hideErrorCases.hideOnSave`: enable/disable auto-fold-on-save/auto-make-transparent-on-save feature
- `go.hideErrorCases.errorCasesOpacity`: configures opacity of error case codes
- `go.hideErrorCases.errorTypeRegexp`: configures which type should be recognized as error type

### 1.0.7

Added hide-on-save feature
Bug fix: folds wrong lines when opened .go file has changed quickly.

### 1.0.6

Revised README, fixed minor bugs.

### 1.0.5

Bug fix: not working for else-if and else, fails to fold nested if block.

### 1.0.4

Bug fix: could not recognize types with selector (like `somepackage.MyError`) as error types.

### 1.0.3

Introduced `go.hideErrorCases.errorTypeRegexp` setting, for working with any custom error types.

### 1.0.1~1.0.2

Just revised README, no functionality change.

### 1.0.0

Initial release of Hide Error Cases
