# S-UI-MODIFIED
**An Advanced Web Panel • Built on SagerNet/Sing-Box**

![](https://img.shields.io/github/v/release/kunori-kiku/s-ui-modified.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kunori-kiku/s-ui-modified)](https://goreportcard.com/report/github.com/kunori-kiku/s-ui-modified)
[![Downloads](https://img.shields.io/github/downloads/kunori-kiku/s-ui-modified/total.svg)](https://img.shields.io/github/downloads/kunori-kiku/s-ui-modified/total.svg)
[![License](https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true)](https://www.gnu.org/licenses/gpl-3.0.en.html)

> **Disclaimer:** This project is only for personal learning and communication, please do not use it for illegal purposes, please do not use it in a production environment
> This is **not** originally produced by kunori-kiku, but a modified version by **alireza0**.
> The original version is at [alireza0/s-ui](https://github.com/alireza0/s-ui)

**If you think this project is helpful to you, you may wish to give a**:star2:

**If you wish to support, please support the original author [alireza0](https://github.com/alireza0/s-ui)**

## Quick Overview
| Features                               |      Enable?       |
| -------------------------------------- | :----------------: |
| Multi-Protocol                         | :heavy_check_mark: |
| Multi-Language                         | :heavy_check_mark: |
| Multi-Client/Inbound                   | :heavy_check_mark: |
| Advanced Traffic Routing Interface     | :heavy_check_mark: |
| Client & Traffic & System Status       | :heavy_check_mark: |
| Subscription Service (link/json + info)| :heavy_check_mark: |
| Dark/Light Theme                       | :heavy_check_mark: |
| API Interface                          | :heavy_check_mark: |

## What are modified?
- Default listen address is set to `127.0.0.1` to prevent security concerns
- Best practice introduced
- Notice message when unsafe method is used
- Added installation script with Github proxy to install `s-ui-modified` on servers that cannot access Github.

## API Documentation

[API-Documentation Wiki](https://github.com/alireza0/s-ui/wiki/API-Documentation)

## Default Installation Information
- Panel Port: 2095
- Panel Path: /s-ui/
- Subscription Port: 2096
- Subscription Path: /sub/
- User/Password: admin

## Best Practice
### Preparation
- [Register Cloudflare Account](https://dash.cloudflare.com), you may follow [**here**](https://developers.cloudflare.com/fundamentals/setup/account/create-account/)
- Buy a domain (suggested: [Namesilo](https://www.namesilo.com/), very cheap, $2 per year for a `.top` domain)
- [Change Your Domain Nameserver To Cloudflare](https://developers.cloudflare.com/dns/zone-setups/full-setup/setup/)
### Initialize ZeroTrust
- Go to the [Zerotrust Dash](https://one.dash.cloudflare.com) and follow the instructions
If you do not have a `visa` or `mastercard`, there are tutorials on how to use random card numbers to pass verifications, **but this may violate the Terms of Use and involve lawsuit, so do not do it! If you did this and caused damage to any entity, it is not my concern as I had warned you!**
**A visa or mastercard is not hard to apply for, so think through!**
- Scroll down to the bottom on the left, go to `Settings-Authentication`
- Scroll down, you will see `Login methods`
- I suggest using `GitHub`, then follow the instructions on the right.
- After saving it, we may start our deployment to the panel
### Create Access Application
- On the left, select `Access-Applications`, click `Add an application`
- Select `Self-hosted`, and name it **without mentioning s-ui**(or our preparations are in vein!), like `auth`, `authentication` or `authentication-germany`
- Click `Add public hostname`, and insert a domain name you would like to access your s-ui (**Must be tertiary!!**, e.g. `de.your-domain.com` but `a.de.your-domain.com` is not permitted!), and again **do not mention the s-ui in your hostname or `s-ui` will appear in your sni requesting your domain!**(And your preparations are in vein). Good examples are `de-1.your-domain.com` or `de-vps.your-domain.com`
- Click `Create new policy` and name it `identity-auth`(preferably). In the selector, select to include `Email`, and insert your email address **to your github account** and save.
- Click `Select existing policies` and select `identity-auth` you just created.
- In the login methods, better to just include `Github`(or your OIDC provider), as it is more succinct.
- Nothing to change in `Experience Settings` section, scroll down and click next
- In `Cross-Origin Resources Sharing (CORS) Settings`, click `Bypass options requests to origin`
- Save the application, but don't close the tab yet.
### Deploy the panel
- Run and better change the port/path settings as you wish
  ```sh
  bash <(curl -Ls https://raw.githubusercontent.com/kunori-kiku/s-ui-modified/master/install.sh)
  ```
### Deploy the tunnel
- In ZeroTrust dash, click `Networks-Tunnels` and click `Create a tunnel`
- Select `Cloudflared` and name it whatever you like. Better only to contain your VPS provider and its region, e.g. `DE-Azure`
- Select suitable operating system and architecture for your VPS in `Choose your environment`, copy the command generated and run it.
  - If your encounter `bash: curl not found`, then install `curl` by your system. If you were debian(ubuntu), run `sudo apt update && sudo apt install -y curl`
- Wait patiently until success
- Click next, insert in the public subdomain you had previously inserted into the creation of Access Application, and insert in the path for your s-ui
- In `type` section, select `http`, and insert in the URL to your s-ui (e.g. `http://localhost:2095/s-ui` by default)
- Click `Save tunnel`, and CONGRATS! You can now access your s-ui via cloudflared tunnel with ZeroTrust Access!

## Install & Upgrade to Latest Version

**Non-proxied(Recommended in most cases)**
```sh
bash <(curl -Ls https://raw.githubusercontent.com/kunori-kiku/s-ui-modified/master/install.sh)
```

**Proxied(Not recommended, only use when necessary)**
```sh
bash <(curl -Ls https://gh-proxy.com/raw.githubusercontent.com/kunori-kiku/s-ui-modified/master/install-proxied.sh)
```

## Install legacy Version

**Step 1:** To install your desired legacy version, add the version to the end of the installation command. e.g., ver `1.0.0`:

```sh
VERSION=1.0.0 && bash <(curl -Ls https://raw.githubusercontent.com/kunori-kiku/s-ui-modified/$VERSION/install.sh) $VERSION
```

## Manual installation

1. Get the latest version of S-UI based on your OS/Architecture from GitHub: [https://github.com/kunori-kiku/s-ui-modified/releases/latest](https://github.com/kunori-kiku/s-ui-modified/releases/latest)
2. **OPTIONAL** Get the latest version of `s-ui.sh` [https://raw.githubusercontent.com/kunori-kiku/s-ui-modified/master/s-ui.sh](https://raw.githubusercontent.com/kunori-kiku/s-ui-modified/master/s-ui.sh)
3. **OPTIONAL** Copy `s-ui.sh` to /usr/bin/ and run `chmod +x /usr/bin/s-ui`.
4. Extract s-ui tar.gz file to a directory of your choice and navigate to the directory where you extracted the tar.gz file.
5. Copy *.service files to /etc/systemd/system/ and run `systemctl daemon-reload`.
6. Enable autostart and start S-UI service using `systemctl enable s-ui --now`
7. Start sing-box service using `systemctl enable sing-box --now`

## Uninstall S-UI

```sh
sudo -i

systemctl disable s-ui  --now

rm -f /etc/systemd/system/sing-box.service
systemctl daemon-reload

rm -fr /usr/local/s-ui
rm /usr/bin/s-ui
```

## Install using Docker

<details>
   <summary>Click for details</summary>

### Usage

**Step 1:** Install Docker

```shell
curl -fsSL https://get.docker.com | sh
```

**Step 2:** Install S-UI

> Build your own image

```shell
git clone https://github.com/kunori-kiku/s-ui-modified
git submodule update --init --recursive
docker build -t s-ui .
```

</details>

## Manual run ( contribution )

<details>
   <summary>Click for details</summary>

### Build and run whole project
```shell
./runSUI.sh
```

### Clone the repository
```shell
# clone repository
git clone https://github.com/kunori-kiku/s-ui-modified
# clone submodules
git submodule update --init --recursive
```


### - Frontend

Visit [s-ui-frontend](https://github.com/kunori-kiku/s-ui-modified-frontend) for frontend code

### - Backend
> Please build frontend once before!

To build backend:
```shell
# remove old frontend compiled files
rm -fr web/html/*
# apply new frontend compiled files
cp -R frontend/dist/ web/html/
# build
go build -o sui main.go
```

To run backend (from root folder of repository):
```shell
./sui
```

</details>

## Languages

- English
- Farsi
- Vietnamese
- Chinese (Simplified)
- Chinese (Traditional)
- Russian

## Features

- Supported protocols:
  - General:  Mixed, SOCKS, HTTP, HTTPS, Direct, Redirect, TProxy
  - V2Ray based: VLESS, VMess, Trojan, Shadowsocks
  - Other protocols: ShadowTLS, Hysteria, Hysteria2, Naive, TUIC
- Supports XTLS protocols
- An advanced interface for routing traffic, incorporating PROXY Protocol, External, and Transparent Proxy, SSL Certificate, and Port
- An advanced interface for inbound and outbound configuration
- Clients’ traffic cap and expiration date
- Displays online clients, inbounds and outbounds with traffic statistics, and system status monitoring
- Subscription service with ability to add external links and subscription
- HTTPS for secure access to the web panel and subscription service (self-provided domain + SSL certificate)
- Dark/Light theme

## Recommended OS

- Ubuntu 20.04+
- Debian 11+
- CentOS 8+
- Fedora 36+
- Arch Linux
- Parch Linux
- Manjaro
- Armbian
- AlmaLinux 9+
- Rocky Linux 9+
- Oracle Linux 8+
- OpenSUSE Tubleweed

## Environment Variables

<details>
  <summary>Click for details</summary>

### Usage

| Variable       |                      Type                      | Default       |
| -------------- | :--------------------------------------------: | :------------ |
| SUI_LOG_LEVEL  | `"debug"` \| `"info"` \| `"warn"` \| `"error"` | `"info"`      |
| SUI_DEBUG      |                   `boolean`                    | `false`       |
| SUI_BIN_FOLDER |                    `string`                    | `"bin"`       |
| SUI_DB_FOLDER  |                    `string`                    | `"db"`        |
| SINGBOX_API    |                    `string`                    | -             |

</details>

## SSL Certificate

<details>
  <summary>Click for details</summary>

### Certbot

```bash
snap install core; snap refresh core
snap install --classic certbot
ln -s /snap/bin/certbot /usr/bin/certbot

certbot certonly --standalone --register-unsafely-without-email --non-interactive --agree-tos -d <Your Domain Name>
```

</details>

## Stargazers over Time
[![Stargazers over time](https://starchart.cc/kunori-kiku/s-ui-modified.svg)](https://starchart.cc/kunori-kiku/s-ui-modified)
