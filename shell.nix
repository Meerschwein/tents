let
  pkgs = import <nixpkgs> {system = "x86_64-linux";};
in
  pkgs.mkShell {
    packages = with pkgs; [
      delve
      go
      gotools
      gofumpt
      gopls

      clingo

      typst
      typst-fmt
      typst-lsp
    ];
  }
