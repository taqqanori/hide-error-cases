import * as vscode from "vscode";
import { parse, ParseResult } from "./parse";
import { error, checkGoFileOpened, parseError } from "./util";

const decorations: vscode.TextEditorDecorationType[] = [];

export async function makeTransparent(
  context: vscode.ExtensionContext,
  showErrorInMessageBox: boolean,
  parseResult?: ParseResult
) {
  if (!checkGoFileOpened(showErrorInMessageBox)) {
    return;
  }
  const opacity = vscode.workspace
    .getConfiguration("go")
    .get("hideErrorCases.errorCasesOpacity", 0.5);
  if (!parseResult) {
    parseResult = await parse(context);
    if (parseResult.status === "failure") {
      parseError(parseResult.failureMessage, showErrorInMessageBox);
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
          loc.end.column
        )
    )
  );
}

export function resetTransparency() {
  decorations.map((d) => d.dispose());
  decorations.splice(0, decorations.length);
}
