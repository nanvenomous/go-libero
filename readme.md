
## Intro

`go-libero` is a simple API-key based guard to any API 

## Motivation
- safe against [timing attacks](https://en.wikipedia.org/wiki/Timing_attack) using `crypto/subtle` from go standard library
- works with any proxy provider that supports [forward_auth](https://docs.goauthentik.io/docs/providers/proxy/forward_auth)

## Usage

build and serve libero
```bash
git clone https://github.com/nanvenomous/go-libero.git && cd go-libero

go build ./...
LIBERO_API_KEY='your-api-key' ./libero
```

libero should now be running on `4444`

however you still need to guard your resource server within the reverse proxy (the following example uses [caddy](https://caddyserver.com/))

assuming you have a resource server running locally on port `4005`

so in your `Caddyfile`:
```Caddyfile
my-domain.com {
  forward_auth :4444 {
    uri /
  }
  reverse_proxy :4005 {
    header_up Host localhost:4005
  }
}
```
- [caddy forward_auth docs](https://caddyserver.com/docs/caddyfile/directives/forward_auth)

## The name
libero is the term from the position in volleyball
go-libero is the gurdian diety of your web api

![nishinoya](https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/631d6f5e-298e-49cf-a4d7-42c35ed82c14/dfkid50-98bacc47-a142-4362-b242-71dd8f5c609d.jpg/v1/fill/w_1280,h_1230,q_75,strp/nishinoya_haikyuu_by_racooncake_dfkid50-fullview.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7ImhlaWdodCI6Ijw9MTIzMCIsInBhdGgiOiJcL2ZcLzYzMWQ2ZjVlLTI5OGUtNDljZi1hNGQ3LTQyYzM1ZWQ4MmMxNFwvZGZraWQ1MC05OGJhY2M0Ny1hMTQyLTQzNjItYjI0Mi03MWRkOGY1YzYwOWQuanBnIiwid2lkdGgiOiI8PTEyODAifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6aW1hZ2Uub3BlcmF0aW9ucyJdfQ.9sglUZnRVsyFhMaOC4ArzUal_9_t5cbFkKRSWsktGYo)
- art by [racooncake on deviantart](https://www.deviantart.com/racooncake/art/Nishinoya-Haikyuu-941441796)

