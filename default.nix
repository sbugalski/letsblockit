{ pkgs ? import <nixpkgs> { } }:
pkgs.buildGoModule.override
{
  go = pkgs.go_1_17;
}
{
  pname = "letsblockit";
  version = "1.0";
  vendorSha256 = "sha256-8SWHnnIAgYjTZegbvoQ1K4oIvdXaxQXYEyHlU6Jdi7Q=";
  src = ./.;
  doCheck = false;
}
