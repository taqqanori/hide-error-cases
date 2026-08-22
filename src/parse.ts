import * as vscode from "vscode";
import * as childProcess from "child_process";
import path = require("path");
import { Shescape } from "shescape";
import * as os from "node:os";

export interface ParseResult {
  status: "success" | "failure";
  failureMessage?: string;
  errorCodeLocations: {
    start: Position;
    end: Position;
    blockStartLine: number;
  }[];
}

export interface Position {
  line: number;
  column: number;
}

// never rejected, parse result is known by status
export function parse(context: vscode.ExtensionContext): Promise<ParseResult> {
  return new Promise((resolve) => {
    if (!vscode.window.activeTextEditor) {
      resolve(errorResult("There are no active text editor."));
      return;
    }
    const src = vscode.window.activeTextEditor.document.getText();
    const errorTypeRegexp = vscode.workspace
      .getConfiguration("go")
      .get("hideErrorCases.errorTypeRegexp", "(E|e)rror");
    const parserDir = context.asAbsolutePath(path.join("out", "parser"));
    const childStdin = childProcess.exec(
      `go run . ${quote(errorTypeRegexp)}`,
      { cwd: parserDir },
      (error, stdout, stderr) => {
        if (error || stderr) {
          resolve(
            errorResult(error ? error.message : stderr ? stderr : undefined),
          );
          return;
        }
        resolve(JSON.parse(stdout) as ParseResult);
      },
    ).stdin;
    if (!childStdin) {
      resolve(errorResult("could not get stdin of child process"));
      return;
    }
    childStdin.write(src);
    childStdin.end();
  });
}

// https://github.com/ericcornelissen/shescape/blob/HEAD/docs/use-cases.md
function quote(s: string) {
  if (os.platform() === "win32") {
    // A shell is required on Windows to run scripts.
    const shescape = new Shescape({ shell: true });
    return shescape.quote(s);
  } else {
    // No shell is preferred because it's safer.
    const shescape = new Shescape({ shell: false });
    return shescape.escape(s);
  }
}

function errorResult(msg?: string): ParseResult {
  return {
    status: "failure",
    failureMessage: msg,
    errorCodeLocations: [],
  };
}
