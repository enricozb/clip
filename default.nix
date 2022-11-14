{ pkgs }:

pkgs.buildGoModule rec {
  pname = "clip";
  version = "0.0.2";

  src = pkgs.lib.cleanSource ./.;
  vendorSha256 = "1y8flix4wd24f4qh56vlfizgj40vpzsrxfaq06cr3bh9d74waccr";

  subPackages = [ "cmd/clip" ];

  meta = with pkgs.lib; {
    description = "Clip is a synchronizing extensible clipboard";
    homepage = "https://github.com/enricozb/clip";
    license = licenses.unlicense;
    maintainers = with maintainers; [ enricozb ];
    platforms = platforms.linux;
  };
}
