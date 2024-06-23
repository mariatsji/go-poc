let nixpkgs = import (builtins.fetchTarball {
      url = "https://github.com/NixOS/nixpkgs/archive/refs/tags/23.11.tar.gz";
  }) {};

in nixpkgs.mkShell {
  name = "go-poc";
  buildInputs = with nixpkgs; [ go gopls ];
}

