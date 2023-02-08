#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

PROJECT_DIR=$(dirname "${CURRENT_DIR}")

export CMODE="dev"

echo "current project dir: ${PROJECT_DIR}"

# 倒入文件
. "${PROJECT_DIR}/scripts/ctrl.help.sh"
. "${PROJECT_DIR}/scripts/utils.sh"


action=${1-}
target=${2-}
args=("$@")

#echo "${args}"
#echo "${action}"
#echo "${target}"

EXECUTOR=""

function start() {
#  check_ipv6_iptables_if_need
  ${EXECUTOR} up -d
}

function stop() {
  if [[ -n "${target}" ]]; then
    ${EXECUTOR} stop "${target}" && ${EXECUTOR} rm -f "${target}"
    return
  fi
  services=$(get_docker_compose_services ignore_db)
  for i in ${services}; do
    ${EXECUTOR} stop "${i}"
  done
  for i in ${services}; do
    ${EXECUTOR} rm -f "${i}" >/dev/null
  done

}

function close() {
  if [[ -n "${target}" ]]; then
    ${EXECUTOR} stop "${target}"
    return
  fi
  services=$(get_docker_compose_services ignore_db)
  for i in ${services}; do
    ${EXECUTOR} stop "${i}"
  done
}

function restart() {
  stop
  echo -e "\n"
  start
}


function check_config_file() {
#  echo "config: ${CONFIG_FILE}"
  if [[ ! -f "${CONFIG_FILE}" ]]; then
    echo "$(gettext 'Configuration file not found'): ${CONFIG_FILE}"
    return 3
  fi
  if [[ -f .env ]]; then
    ls -l .env | grep "${CONFIG_FILE}" &>/dev/null
    code="$?"
    if [[ "$code" != "0" ]]; then
      echo ".env $(gettext 'There is a problem with the soft connection, Please update it again')"
      rm -f .env
    fi
  fi

  if [[ ! -f .env ]]; then
    ln -s "${CONFIG_FILE}" .env
  fi
}

function pre_check() {
  check_config_file || return 3
}

function main() {
  if [[ "${action}" == "help" || "${action}" == "h" || "${action}" == "-h" || "${action}" == "--help" ]]; then
    echo ""
  elif [[ "${action}" == "" ]]; then
    echo ""
  else
    pre_check || return 3
    EXECUTOR=$(get_docker_compose_cmd_line)
    echo "EXECUTOR: ${EXECUTOR}"
  fi

  prepare_config

  case "${action}" in
  up)
    if [[ -z "${target}" ]]; then
      start
    else
      ${EXECUTOR} up "${target}" -d
    fi
    ;;
  start)
    start
    ;;
  stop)
    stop
    ;;
  close)
    close
    ;;
  restart)
    restart
    ;;
  status)
    ${EXECUTOR} ps
    ;;
  down)
    if [[ -z "${target}" ]]; then
      ${EXECUTOR} down -v
    else
      ${EXECUTOR} stop "${target}" && ${EXECUTOR} rm -f "${target}"
    fi
    ;;
  show_services)
    get_docker_compose_services
    ;;
  raw)
    ${EXECUTOR} "${args[@]:1}"
    ;;

  help)
    usage
    ;;
  --help)
    usage
    ;;
  h)
    usage
    ;;
  -h)
    usage
    ;;
  *)
    echo "No such command: ${action}"
    usage
    ;;
  esac
}

main "$@"