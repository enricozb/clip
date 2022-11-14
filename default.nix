{ pkgs ? import <nixpkgs> { } }:

pkgs.buildGoModule rec {
  pname = "clip";
  version = "0.0.2";

  src = pkgs.lib.cleanSource ./.;
  vendorSha256 = "sha256-mTHFyWkJrpGZAVi5nvW/GxD5fnR0mwIxcUQ0TnqkDvk=";

  subPackages = [ "cmd/clip" ];

  meta = with pkgs.lib; {
    description = "Clip is a synchronizing extensible clipboard";
    homepage = "https://github.com/enricozb/clip";
    license = licenses.unlicense;
    maintainers = with maintainers; [ enricozb ];
    platforms = platforms.linux;
  };
}
