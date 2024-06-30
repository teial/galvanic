{ pkgs, lib, config, inputs, ... }:
let
  pkgs-unstable = import inputs.nixpkgs-unstable { system = pkgs.stdenv.system; };
in {
    env.GOEXPERIMENT = "rangefunc,aliastypeparams";
    packages = [
        pkgs-unstable.goimports-reviser
        pkgs-unstable.gopls
        pkgs-unstable.golines
        pkgs-unstable.go_1_23
    ];
}
