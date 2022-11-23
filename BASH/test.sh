#!/bin/bash

# setup script practice 

#check for root

check_for_root () {
    if [ "$EUID" -ne 0 ]
      then echo -e "\n\n Script must be run with sudo or as root \n"
      exit
    else
      # 02.19.21 - Kali 2021.1 + MSF 6.0.30-DEV Released
      # Remove any prior hold on metasploit-framework at startup
      eval apt-mark unhold metasploit-framework >/dev/null 2>&1
    fi
    }

apt_update() {
    echo -e "\n running: apt update \n"
    eval apt -y update
    }

apt_upgrade() {
    echo -e "\n running: apt upgrade \n"
    eval apt -y upgrade
    }

apt_autoremove() {
    echo -e "\n running: apt autoremove \n"
    eval apt -y autoremove
    }

apt_update_complete() {
    echo -e "\n apt update - complete"
    }

apt_upgrade_complete() {
    echo -e "\n apt upgrade - complete"
    }

apt_autoremove_complete() {
    echo -e "\n apt autoremove - complete"
    }

apt_install_stuff_i_like() {
    echo -e "\n Installing things I like into Kali \n"
    eval apt -y install seclists curl enum4linux feroxbuster gobuster impacket-scripts nbtscan nikto nmap onesixtyone oscanner redis-tools smbclient smbmap snmp sslscan sipvicious tnscmd10g whatweb wkhtmltopdf ffuf evil-winrm name-that-hash neofetch 
}


check_for_root
apt_update
apt_update_complete
apt_upgrade
apt_update_complete
apt_autoremove
apt_autoremove_complete
apt_install_stuff_i_like