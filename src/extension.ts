// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import { fold } from "./fold";
import { makeTransparent, resetTransparency } from "./transparent";

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
      resetTransparency();
    }
  );
  context.subscriptions.push(disposable);
}

// this method is called when your extension is deactivated
export function deactivate() {
  resetTransparency();
}
