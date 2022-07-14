// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import * as child_process from "child_process";
import path = require("path");

const decorations: vscode.TextEditorDecorationType[] = [];

// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // The command has been defined in the package.json file
  // Now provide the implementation of the command with registerCommand
  // The commandId parameter must match the command field in package.json

  // fold error cases command
  let disposable = vscode.commands.registerCommand(
    "hide-error-cases.foldErrorCases",
    () => {
      fold(context, true);
    }
  );
  context.subscriptions.push(disposable);

  // make error cases transparent command
  disposable = vscode.commands.registerCommand(
    "hide-error-cases.makeErrorCasesTransparent",
    () => {
      makeTransparent(context, 0.5, true);
    }
  );
  context.subscriptions.push(disposable);

  // reset error cases transparency command
  disposable = vscode.commands.registerCommand(
    "hide-error-cases.resetErrorCasesTransparency",
    () => {
      decorations.map((d) => d.dispose());
      decorations.splice(0, decorations.length);
    }
  );
  context.subscriptions.push(disposable);
}

// this method is called when your extension is deactivated
export function deactivate() {}

async function fold(
  context: vscode.ExtensionContext,
  showErrorInMessageBox: boolean,
  parseResult?: ParseResult
) {
  if (!parseResult) {
    parseResult = await parse(context);
    if (parseResult.status === "failure") {
      parseError(parseResult, showErrorInMessageBox);
      return;
    }
  }
  const selectionLines = parseResult.errorCodeLocations.reduce<number[]>(
    (lines, location) => {
      lines.push(location.start.line);
      return lines;
    },
    []
  );
  vscode.commands.executeCommand("editor.fold", {
    level: 1,
    selectionLines,
  });
}

async function makeTransparent(
  context: vscode.ExtensionContext,
  opacity: number,
  showErrorInMessageBox: boolean,
  parseResult?: ParseResult
) {
  if (!parseResult) {
    parseResult = await parse(context);
    if (parseResult.status === "failure") {
      parseError(parseResult, showErrorInMessageBox);
      return;
    }
  }
  if (!vscode.window.activeTextEditor) {
    error("There are no active text editor.", showErrorInMessageBox);
    return;
  }
  const decoration = vscode.window.createTextEditorDecorationType({
    opacity: opacity.toString(),
  });
  decorations.push(decoration);
  vscode.window.activeTextEditor.setDecorations(
    decoration,
    parseResult.errorCodeLocations.map<vscode.Range>(
      (loc) =>
        new vscode.Range(
          loc.start.line - 1,
          loc.start.column - 1,
          loc.end.line - 1,
          loc.end.column - 1
        )
    )
  );
}

interface ParseResult {
  status: "success" | "failure";
  failureMessage?: string;
  errorCodeLocations: {
    start: Position;
    end: Position;
  }[];
}

interface Position {
  line: number;
  column: number;
}

// never rejected, parse result is known by status
function parse(context: vscode.ExtensionContext): Promise<ParseResult> {
  return new Promise((resolve) => {
    if (!vscode.window.activeTextEditor) {
      resolve(errorResult("There are no active text editor."));
      return;
    }
    const src = vscode.window.activeTextEditor.document.getText();
    const parserDir = context.asAbsolutePath(path.join("out", "parser"));
    const childStdin = child_process.exec(
      `go run .`,
      { cwd: parserDir },
      (error, stdout, stderr) => {
        if (error || stderr) {
          resolve(
            errorResult(error ? error.message : stderr ? stderr : undefined)
          );
          return;
        }
        resolve(JSON.parse(stdout) as ParseResult);
      }
    ).stdin;
    if (!childStdin) {
      resolve(errorResult("could not get stdin of child process"));
      return;
    }
    childStdin.write(src);
    childStdin.end();
  });
}

function errorResult(msg?: string): ParseResult {
  return {
    status: "failure",
    failureMessage: msg,
    errorCodeLocations: [],
  };
}

function parseError(parseResult: ParseResult, showInMessageBox = false) {
  error(
    `Failed to parse .go file ${
      parseResult.failureMessage
        ? `(detail:${parseResult.failureMessage}})`
        : ""
    }.`,
    showInMessageBox
  );
}

function error(msg: string, showInMessageBox = false) {
  const header = "Hide Error Cases: ";
  if (showInMessageBox) {
    vscode.window.showErrorMessage(header + msg);
  } else {
    console.error(header + msg);
  }
}
