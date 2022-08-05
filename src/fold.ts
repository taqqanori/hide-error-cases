import * as vscode from "vscode";
import { parse, ParseResult } from "./parse";
import { checkGoFileOpened, parseError } from "./util";

export async function fold(
  context: vscode.ExtensionContext,
  showErrorInMessageBox: boolean,
  parseResult?: ParseResult
) {
  if (!checkGoFileOpened(showErrorInMessageBox)) {
    return;
  }
  if (!parseResult) {
    parseResult = await parse(context);
    if (parseResult.status === "failure") {
      parseError(parseResult.failureMessage, showErrorInMessageBox);
      return;
    }
  }
  if (parseResult.errorCodeLocations.length === 0) {
    return;
  }
  const selectionLines = parseResult.errorCodeLocations.reduce<number[]>(
    (lines, location) => {
      lines.push(location.blockStartLine - 1);
      return lines;
    },
    []
  );
  vscode.commands.executeCommand("editor.fold", {
    levels: 1,
    selectionLines,
    direction: "down",
  });
}
