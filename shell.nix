{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/73ad5f9e147c0d2a2061f1d4bd91e05078dc0b58.tar.gz") {} }:

pkgs.mkShell {
    name = "commentify-env";
    buildInputs = [
        pkgs.nodejs
        pkgs.go_1_18
        pkgs.yarn
        pkgs.go-task
    ];
    shellHook = ''
        echo "You're inside commentify dev environment"
    '';
}
