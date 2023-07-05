# Download Google Fonts file

Download all Google Font files for a given font URL.

```sh
go build -o $GOPATH/bin/google-font-download ./cmd/download
```

Usage:

```sh
mkdir test
google-font-download \
    --url='https://fonts.googleapis.com/css2?family=IBM+Plex+Serif:ital,wght@0,400;0,500;0,700;1,400;1,700&family=Source+Code+Pro:ital,wght@0,400;0,700;1,400;1,700&family=Amaranth&display=swap' \
    --target="$(pwd)/test"
tree test/
```

```
test
└── s
    ├── amaranth
    │   └── v18
    │       └── KtkuALODe433f0j1zPnC.ttf
    ├── ibmplexserif
    │   └── v19
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi2k_hIz.ttf
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi3s-BIz.ttf
    │       ├── jizBREVNn1dOx-zrZ2X3pZvkTiUa2zI.ttf
    │       ├── jizDREVNn1dOx-zrZ2X3pZvkThUY.ttf
    │       └── jizGREVNn1dOx-zrZ2X3pZvkTiUa4442q14.ttf
    └── sourcecodepro
        └── v22
            ├── HI_diYsKILxRpg3hIP6sJ7fM7PqPMcMnZFqUwX28DCuXhM4.ttf
            ├── HI_diYsKILxRpg3hIP6sJ7fM7PqPMcMnZFqUwX28DMyQhM4.ttf
            ├── HI_jiYsKILxRpg3hIP6sJ7fM7PqlOPHYvDP_W9O7GQTTbI1rSQ.ttf
            └── HI_jiYsKILxRpg3hIP6sJ7fM7PqlOPHYvDP_W9O7GQTTi4prSQ.ttf
```
