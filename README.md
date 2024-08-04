# sonicstars

Connect to a subsonic instance, download the list of starred songs, and create a
pls file (or output the list). Does _not_ support starred genres or artists or
elsewise.

## Example

```
make && ./sonicstars --url https://music.wat --user sungo --password awful --output starred pls
```

## Compatibility

### Navidrome

Technically, this work with Navidrome. The output is ... not useful, though, as
Navidrome returns an internal path seemingly based on metadata, not the actual
file path. The file path is of course visible in the web UI, but not in the
subsonic API.

See [#1380](https://github.com/navidrome/navidrome/issues/1380) and
[#1309](https://github.com/navidrome/navidrome/issues/1309), specifically
[this](https://github.com/navidrome/navidrome/issues/1309#issuecomment-912753102)
comment.

# Support / Contributing

This is a personal side project and will get about that much attention, maybe
less. If you have patches, feel free to contact me (see https://sungo.io) but I
make no promise as to when or if I'll respond. But, feel free to fork the code,
respecting the license, and have your way with it.

# Licensing

Licensed under 0BSD. See LICENSE.md for details
