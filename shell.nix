{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/32356ce11b8cc5cc421b68138ae8c730cc8ad4a2.tar.gz") {} }:

pkgs.mkShell {
    name = "commentify-env";
    buildInputs = [
        pkgs.nodejs
        pkgs.go_1_17
        pkgs.yarn
        pkgs.go-task
        pkgs.postgresql_14
    ];
    shellHook = ''
        export PGDATA=$HOME/postgres/data
        export PGHOST=$PGDATA
        export PGLOG=$HOME/postgres/postgres.log

        mkdir -p $HOME/postgres

        if [ ! -d $PGDATA ]; then
            echo 'Initializing postgresql database...'
            initdb -D $PGDATA --no-locale --encoding=UTF-8
        fi

        if ! pg_ctl status; then
            echo 'starting postgresql database...'
            pg_ctl start -l $PGLOG -o "--unix_socket_directories='$PGHOST'"
        fi
    '';
}