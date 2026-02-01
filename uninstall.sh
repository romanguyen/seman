#!/usr/bin/env sh

set -euo pipefail

app_name="seman"
bin_dir="$HOME/.local/bin"
bin_path="$bin_dir/$app_name"
data_root="${XDG_DATA_HOME:-$HOME/.local/share}"
data_dir="$data_root/$app_name"

remove_data="false"
force="false"

usage() {
  echo "Usage: ./uninstall.sh [--data] [--force]"
  echo "  --data   remove user data directory"
  echo "  --force  do not prompt when removing data"
}

for arg in "$@"; do
  case "$arg" in
    --data|-d) remove_data="true" ;;
    --force|-f) force="true" ;;
    --help|-h) usage; exit 0 ;;
    *) echo "Unknown option: $arg"; usage; exit 1 ;;
  esac
done

if [ -f "$bin_path" ]; then
  rm -f "$bin_path"
  echo "Removed binary: $bin_path"
else
  echo "Binary not found: $bin_path"
fi

if [ "$remove_data" = "true" ]; then
  if [ -d "$data_dir" ]; then
    if [ "$force" = "true" ]; then
      rm -rf "$data_dir"
      echo "Removed data directory: $data_dir"
    else
      printf "Remove data directory %s? [y/N] " "$data_dir"
      read -r reply
      case "$reply" in
        y|Y)
          rm -rf "$data_dir"
          echo "Removed data directory: $data_dir"
          ;;
        *)
          echo "Data preserved: $data_dir"
          ;;
      esac
    fi
  else
    echo "Data directory not found: $data_dir"
  fi
else
  echo "Data preserved: $data_dir"
fi
