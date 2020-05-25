#!/bin/bash

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." >/dev/null 2>&1 && pwd)"
RUNMAKE="false"
OPENAPI_DIRS=()

RED="\033[31m"
CYAN="\033[36m"
GREEN="\033[32m"
BOLD="\033[1m"
DEFAULT="\033[0m"

run() {(
  while IFS= read -d $'\0' -r file ; do
     OPENAPI_DIRS=("${OPENAPI_DIRS[@]}" "$file")
  done < <(find ${PROJECT_DIR} \( -iname openapi -o -iname dist \)  -print0)
  set -e
  # git diff cannot find diff for a several path. That's why we need to iterate and check again
  RUNMAKE="false"
  echo -e "${CYAN}Checking openapi specification directories for changes${DEFAULT}"
  for DIR in "${OPENAPI_DIRS[@]}"; do
    if ! $( git diff --quiet --exit-code --cached -- "$DIR" ) || ! $( git diff --quiet --exit-code -- "$DIR" ); then
      echo -e "${CYAN}Directory has changes: ${DIR}${DEFAULT}"
      RUNMAKE="true"
    fi
  done

  if [[ "${RUNMAKE}" == "true" ]]; then
    echo -e "${CYAN}Run code generation${DEFAULT}";
    make build || {
      echo -e "${RED}${BOLD}Code generation failed${DEFAULT}"
      exit 1
    }
  else
    echo -e "${CYAN}Openapi specification directories do not contain changes.\nGeneration skipped${DEFAULT}";
  fi;

  echo -e "${CYAN}Add all generated files to the git index${DEFAULT}"
  for DIR in "${OPENAPI_DIRS[@]}"; do
      if ! $( git diff --quiet --exit-code --cached -- "$DIR" ) || ! $( git diff --quiet --exit-code -- "$DIR" ) ;
      then
        PARENT_DIR=$( dirname -- $DIR )
        git add --all "${PARENT_DIR}" || {
          echo -e "${RED}${BOLD}Git add failed${DEFAULT}"
          exit 1
        }
      fi
  done
  echo -e "${GREEN}${BOLD}Pre-commit execution successfully${DEFAULT}"
)}

run || {
  echo -e "${RED}${BOLD}Pre-commit execution failed${DEFAULT}"
  exit 1
}
