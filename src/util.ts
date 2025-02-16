import * as vscode from "vscode";

export function checkGoFileOpened(showErrorInMessageBox: boolean): boolean {
  const ret = isGoFileOpened();
  if (!ret) {
    error(
      ".go file not opened in current active text editor.",
      showErrorInMessageBox
    );
  }
  return ret;
}

export function isGoFileOpened(): boolean {
  return vscode.window.activeTextEditor?.document.languageId === "go";
}

export function getCurrentFileName() {
  return vscode.window.activeTextEditor?.document.fileName;
}

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
