#!/usr/bin/env bash

function usage() {
  echo "$(gettext 'omgind Deployment Management Script')"
  echo
  echo "Usage: "
  echo "  ./ctrl.{mode}.sh [COMMAND] [ARGS...]"
  echo "  ./ctrl.{mode}.sh --help"
  echo
  echo "Management Commands: "
  echo "  start             $(gettext 'Start   JumpServer')"
  echo "  stop              $(gettext 'Stop    JumpServer')"
  echo "  close             $(gettext 'Close   JumpServer')"
  echo "  restart           $(gettext 'Restart JumpServer')"
  echo "  status            $(gettext 'Check   JumpServer')"
  echo "  down              $(gettext 'Offline JumpServer')"
  echo "  uninstall         $(gettext 'Uninstall JumpServer')"
  echo
  echo "More Commands: "
  echo "  show_services     $(gettext 'show all services')"
  echo "  raw               $(gettext 'Execute the original docker-compose command')"
  echo "  tail [service]    $(gettext 'View log')"

}
