{
  description = "Backend API Challenge";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    go-nixpkgs.url = "github:nixos/nixpkgs/63143ac2c9186be6d9da6035fa22620018c85932";
  };

  outputs = { self, nixpkgs, ... } @inputs:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.x86_64-linux.default =
        pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            inputs.go-nixpkgs.legacyPackages.${system}.go
            go-migrate
            docker-compose
            sqlc
          ];
        };
    };
}
