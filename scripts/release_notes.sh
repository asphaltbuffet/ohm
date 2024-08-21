#!/bin/bash
set -euo pipefail

if [ $# -ne 1 ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

rm -rf release_notes
mkdir release_notes

chglog="CHANGELOG.md"

# check if the file exists
if [ ! -f "$chglog" ]; then
  echo "File not found: $chglog"
  exit 1
fi

# extract content for specific version
awk -v semver="$1" '
  BEGIN { found=0 }
  {
    if ($0 ~ "^## \\[" semver "\\]") {
      found=1;
      next
    } else if ($0 ~ "^## " || $0 ~ "^\\[") {
      found=0
    }

    if (found) { print }
  }
  END { print "\nSee full changelog at [CHANGELOG.md](CHANGELOG.md)\n"}
' "$chglog" > "release_notes/notes.md"