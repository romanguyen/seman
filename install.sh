#!/usr/bin/env sh

set -euo pipefail

app_name="seman"
bin_dir="$HOME/.local/bin"
bin_path="$bin_dir/$app_name"
repo_root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)

mkdir -p "$bin_dir"

echo "Building $app_name..."
(cd "$repo_root" && go build -o "$bin_path.tmp" "./cmd/$app_name")
mv "$bin_path.tmp" "$bin_path"
chmod 0755 "$bin_path"
echo "Installed to $bin_path"

case ":$PATH:" in
  *":$bin_dir:"*) ;; 
  *) echo "Warning: $bin_dir is not in PATH" ;; 
esac

echo "Run: $app_name"
