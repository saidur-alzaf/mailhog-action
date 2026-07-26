
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root_dir"

mapfile -t go_files < <(find . -name '*.go' ! -path './vendor/*' ! -path './.git/*' -print)

if [ ${#go_files[@]} -eq 0 ]; then
  echo "No Go files found to check."
  exit 0
fi

unformatted=$(gofmt -l "${go_files[@]}")

if [ -n "$unformatted" ]; then
  echo "The following Go files are not formatted:"
  echo "$unformatted"
  echo "Run 'gofmt -w <file>' to fix them."
  exit 1
fi

echo "All Go files are properly formatted."
