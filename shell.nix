{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/32356ce11b8cc5cc421b68138ae8c730cc8ad4a2.tar.gz") {} }:

pkgs.mkShell {
    name = "commentify-env";
    buildInputs = [
        pkgs.nodejs
        pkgs.go_1_17
        pkgs.yarn
        pkgs.go-task
    ];
    shellHook = ''
        echo "You're inside commentify dev environment"
    '';
}