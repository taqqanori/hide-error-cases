import * as vscode from "vscode";
import { parse, ParseResult } from "./parse";
import { parseError } from "./util";

export async function fold(
  context: vscode.ExtensionContext,
  showErrorInMessageBox: boolean,
  parseResult?: ParseResult
) {
  if (!parseResult) {
    parseResult = await parse(context);
    if (parseResult.status === "failure") {
      parseError(parseResult.failureMessage, showErrorInMessageBox);
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
