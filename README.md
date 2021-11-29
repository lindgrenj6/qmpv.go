# qmpv.go

A simple little library I use to run mpv in the background with a queue enabled, this way I can just continually run
```bash
$ qmpv https://www.youtube.com/watch?v=rTgj1HxmUbg
```

and the videos will be added to the running mpv instance that:
1. Is on all workspaces
2. Is always on top
3. Will start the first time if necessary

Super handy, can be imported or ran manually.

---

### Installation

A `Makefile` is provided - build with `make`, install with `make install`
