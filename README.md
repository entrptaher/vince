
<p align="center">
  <picture width="640" height="250">
    <source media="(prefers-color-scheme: dark)" srcset="./app/images/logo-darkmode.svg">
    <source media="(prefers-color-scheme: light)" srcset="./app/images/mark.svg">
    <img alt="esbuild: An extremely fast JavaScript bundler" src="logo.svg">
  </picture>
  <br>
  <a href="https://vinceanalytics.com/">Website</a> |
  <a href="https://vinceanalytics.com/blog/deploy-local/">Getting started</a> |
  <a href="https://vinceanalytics.com/tags/api/">API</a>  |
  <a href="https://demo.vinceanalytics.com/v1/share/vinceanalytics.com?auth=Ls9tV4pzqOn7BJ7-&demo=true">Demo</a> 
</p>

**Vince** is a self hosted alternative to Google Analytics.


# Features

- **Automatic TLS** native support for let's encrypt.
- **Drop in replacement for plausible** you can use existing plausible  scripts and just point them to the vince instance (note that vince is lean and only covers features for a single entity self hosting, so it is not our goal to be feature parity with plausible).
- **Outbounds links tracking**
- **File download tracking**
- **404 page tracking**
- **Custom event tracking**
- **Time period comparison**
- **Public dashboards** allow access to the dashoard to anyone(by default all dashboards are private).
- **Unique shared access** generate unique links to dahboards that can be password protected.
- **Zero Dependency**: Ships a single binary with everything in it. No runtime dependency.
- **Easy to operate**: One line commandline flags with env variables is all you need.
- **Unlimited sites**: There is no limit on how many sites you can manage.
- **Unlimited events**: scale according to availbale resources.
- **Privacy friendly**: No cookies and fully compliant with GDPR, CCPA and PECR.


# Installation

Vince ships a single executable without any dependencies.


## Installing

### MacOS and Linux

```
curl -fsSL https://vinceanalytics.com/install.sh | bash
```

### Docker

```
docker pull ghcr.io/vinceanalytics/vince
```

### Helm
```
❯ helm repo add vince http://vinceanalytics.com/charts
❯ helm install vince vince/vince
```

### Download 

[see release page](https://github.com/vinceanalytics/vince/releases)


## Checking installation

```
vince --version
```

## Start vince

***create admin***
```
❯ vince admin --name acme --password 1234
```

***start server***
```
❯ vince serve                            
2024/10/23 15:32:08 [JOB 1] WAL file vince-data/pebble/000002.log with log number 000002 stopped reading at offset: 124; replayed 1 keys in 1 batches
2024/10/23 15:32:08 INFO starting event processing loop
2024/10/23 15:32:08 INFO starting server addr=:8080
```

## Start with docker and docker compose

You can start vince with docker and docker compose.

To start with docker,

```
docker run -it -p 8080:8080 -v ./vince:/data ghcr.io/vinceanalytics/vince serve --data=/data --adminName acme --adminPassword 1234
```

Here is a sample docker-compose.yml file,

```yml
name: <your project name>
services:
    vince:
        stdin_open: true
        tty: true
        ports:
            - 8080:8080
        volumes:
            - ./vince:/data
        image: ghcr.io/vinceanalytics/vince
        command: serve --data=/data --adminName acme --adminPassword 1234
```

# Comparison with Plausible Analytics

| feature |  vince | plausible |
|---------|--------| -----------|
| Entrerprise features | :x:    | :white_check_mark:|
| Hosted offering | :x:    | :white_check_mark:|
| Multi tenant | :x:    | :white_check_mark:|
| Funnels | :x:    | :white_check_mark:|
| Goals Conversion |  :white_check_mark:  | :white_check_mark:|
| Unique visitors |  :white_check_mark:  | :white_check_mark:|
| Total visits |  :white_check_mark:  | :white_check_mark:|
| Page views |  :white_check_mark:  | :white_check_mark:|
| Views per visit |  :white_check_mark:  | :white_check_mark:|
| Visit duration |  :white_check_mark:  | :white_check_mark:|
| Breakdown by **Cities**, **Sources**, **Pages** and **Devices**   |  :white_check_mark:  | :white_check_mark:|
| Self Hosted |  :white_check_mark:  | :white_check_mark:|
| <1KB script |  :white_check_mark:  | :white_check_mark:|
| No Cookies(GDPR, PECR compliant) |  :white_check_mark:  | :white_check_mark:|
| 100% data ownershiip |  :white_check_mark:  | :white_check_mark:|
| Unique shared access to stats|  :white_check_mark:  | :white_check_mark:|
| Outbound links tracking |  :white_check_mark:  | :white_check_mark:|
| File download tracking |  :white_check_mark:  | :white_check_mark:|
| 404 page tracking |  :white_check_mark:  | :white_check_mark:|
| Time period comparisons |  :white_check_mark:  | :white_check_mark:|
| Unlimited sites |  :white_check_mark:  | :x:|
| Unlimited events |  :white_check_mark:  | :x:|
| Zero dependency |  :white_check_mark:  | :x: (needs elixir, clickhouse, postgresql ...etc)|
| Automatic TLS |  :white_check_mark:  | :x:|


# Credit

[Plausible Analytics](https://github.com/plausible/analytics) : `vince` started as a Go port of plausible with a focus on self hosting.

# Note from Author

My name is [Geofrey Ernest](https://github.com/gernest), a senior software engineer based in Arusha, Tanzania.

`vince` is a result of 3 years of research on how to apply Compressed Roaring Bitmap Indexes on web analytics events to achieve storage efficiency and supreme query speeds.

I am currently looking for fulltime remote work, and I'm struggling because Remote now seems to mean Remote US and Remote EU and I'm stuck here in Tanzania. If you are a hiring manager please **give** me a shot, I promise you won't regret it.