// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import * as child_process from "child_process";
import path = require("path");
import { stderr } from "process";

// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // The command has been defined in the package.json file
  // Now provide the implementation of the command with registerCommand
  // The commandId parameter must match the command field in package.json
  let disposable = vscode.commands.registerCommand(
    "hide-error-cases.foldErrorCases",
    async () => {
      const parseResult = await parse(context);
      // The code you place here will be executed every time your command is executed
      // Display a message box to the user
      vscode.window.showInformationMessage("foldErrorCases");
    }
  );
  context.subscriptions.push(disposable);
  disposable = vscode.commands.registerCommand(
    "hide-error-cases.makeErrorCasesTransparent",
    () => {
      // The code you place here will be executed every time your command is executed
      // Display a message box to the user
      vscode.window.showInformationMessage("makeErrorCasesTransparent");
    }
  );
  context.subscriptions.push(disposable);
}

// this method is called when your extension is deactivated
export function deactivate() {}

interface ParseResult {
  status: "success" | "failure";
  failureMessage?: string;
  errorCodeLocations: {
    startLine: number;
    endLine: number;
  }[];
}

// never rejected, parse result is know by status
function parse(context: vscode.ExtensionContext): Promise<ParseResult> {
  return new Promise((resolve) => {
    const src = vscode.window.activeTextEditor?.document.getText();
    const parserDir = context.asAbsolutePath(path.join("out", "parser"));
    const childStdin = child_process.exec(
      `go run .`,
      { cwd: parserDir },
      (error, stdout, stderr) => {
        if (error || stderr) {
          resolve(
            errorResult(error ? error.message : stderr ? stderr : undefined)
          );
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
