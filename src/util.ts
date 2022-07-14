import * as vscode from "vscode";

export function parseError(failureMessage?: string, showInMessageBox = false) {
  error(
    `Failed to parse .go file ${
      failureMessage ? `(detail:${failureMessage}})` : ""
    }.`,
    showInMessageBox
  );
}

export function error(msg: string, showInMessageBox = false) {
  const header = "Hide Error Cases: ";
  if (showInMessageBox) {
    vscode.window.showErrorMessage(header + msg);
  } else {
    console.error(header + msg);
  }
}
