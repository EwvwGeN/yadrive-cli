# CLI application for access to yandex disk

SLI application for using yandex disk

# Table of Contents
- [Installation](#Installation)
- [Application registration](#Application-registration)
- [Yandex disk api](#Yandex-disk-api)
- [Usage](#Usage)

## Installation

You can simply intstall the application by command:</br>
`go install github.com/EwvwGeN/yadrive-cli@latest --ldflags="-X github.com/EwvwGeN/yadrive-cli/cmd.userSecret=<your_own_secret>"`

In command you need to pass your secret, that will be used for encripting/decripting file, which store confidential info.

## Application registration

Visit the `https://oauth.yandex.com/client/new/id` or `https://oauth.yandex.ru/client/new/id` and create new application. Client id and Client secret will be need to get oauth token.

## Yandex disk api

See `https://yandex.com/dev/disk-api/doc/en/`

## Usage

comming soon