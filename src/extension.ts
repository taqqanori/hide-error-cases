// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import { parse } from "./parse";
import { fold } from "./fold";
import { makeTransparent, resetTransparency } from "./transparent";
import { getCurrentFileName, isGoFileOpened } from "./util";

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
      makeTransparent(context, true);
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

  // setup listeners enabled by config
  setupConfigurables(context);
  vscode.workspace.onDidChangeConfiguration((event) => {
    setupConfigurables(context);
  });

  // perform auto fold/make-transparent
  autoFoldAndMakeTransparent(context);
}

const configurableDisposables: vscode.Disposable[] = [];
function setupConfigurables(context: vscode.ExtensionContext) {
  configurableDisposables.forEach((c) => c.dispose());
  configurableDisposables.splice(0);
  if (autoFoldEnabled() || autoTransparentEnabled()) {
    const autoHide = vscode.window.onDidChangeActiveTextEditor(async () => {
      autoFoldAndMakeTransparent(context);
    });
    configurableDisposables.push(autoHide);
    context.subscriptions.push(autoHide);
    if (hideOnSaveEnabled()) {
      resetTransparency();
      const hideOnSave = vscode.workspace.onDidSaveTextDocument(async () => {
        autoFoldAndMakeTransparent(context);
      });
      configurableDisposables.push(hideOnSave);
      context.subscriptions.push(hideOnSave);
    }
  }
}

async function autoFoldAndMakeTransparent(context: vscode.ExtensionContext) {
  if (!isGoFileOpened() || (!autoFoldEnabled() && !autoTransparentEnabled())) {
    return;
  }
  const targetFileName = getCurrentFileName();
  const parseResult = await parse(context);
  if (autoFoldEnabled()) {
    fold(context, false, parseResult, targetFileName);
  }
  if (autoTransparentEnabled()) {
    makeTransparent(context, false, parseResult, targetFileName);
  }
}

function autoFoldEnabled(): boolean {
  return vscode.workspace
    .getConfiguration("go")
    .get("hideErrorCases.autoFold", false);
}

function autoTransparentEnabled(): boolean {
  return vscode.workspace
    .getConfiguration("go")
    .get("hideErrorCases.autoMakeTransparent", true);
}

function hideOnSaveEnabled(): boolean {
  return vscode.workspace
    .getConfiguration("go")
    .get("hideErrorCases.hideOnSave", false);
}

// this method is called when your extension is deactivated
export function deactivate() {
  resetTransparency();
}
