{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    dart
    flutter
  ];
  shellHook = "export GPG_TTY=$(tty)";
}
