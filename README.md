# Download Google Font files

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
    │       └── KtkuALODe433f0j1zMnFHdCIwWzX.woff2
    ├── ibmplexserif
    │   └── v19
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi2k_iI0q1vjitOh.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi2k_iI5q1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi2k_iI6q1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi2k_iI7q1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi2k_iIwq1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi3s-CI0q1vjitOh.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi3s-CI5q1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi3s-CI6q1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi3s-CI7q1vjitOh3oc.woff2
    │       ├── jizAREVNn1dOx-zrZ2X3pZvkTi3s-CIwq1vjitOh3oc.woff2
    │       ├── jizBREVNn1dOx-zrZ2X3pZvkTiUa6zETjnTLgNuZ5w.woff2
    │       ├── jizBREVNn1dOx-zrZ2X3pZvkTiUa6zUTjnTLgNs.woff2
    │       ├── jizBREVNn1dOx-zrZ2X3pZvkTiUa6zgTjnTLgNuZ5w.woff2
    │       ├── jizBREVNn1dOx-zrZ2X3pZvkTiUa6zoTjnTLgNuZ5w.woff2
    │       ├── jizBREVNn1dOx-zrZ2X3pZvkTiUa6zsTjnTLgNuZ5w.woff2
    │       ├── jizDREVNn1dOx-zrZ2X3pZvkTiUQ2zcZiVbJsNo.woff2
    │       ├── jizDREVNn1dOx-zrZ2X3pZvkTiUR2zcZiVbJsNo.woff2
    │       ├── jizDREVNn1dOx-zrZ2X3pZvkTiUS2zcZiVbJsNo.woff2
    │       ├── jizDREVNn1dOx-zrZ2X3pZvkTiUb2zcZiVbJsNo.woff2
    │       ├── jizDREVNn1dOx-zrZ2X3pZvkTiUf2zcZiVbJ.woff2
    │       ├── jizGREVNn1dOx-zrZ2X3pZvkTiUa4442m13pjfGj7oaMBg.woff2
    │       ├── jizGREVNn1dOx-zrZ2X3pZvkTiUa4442m1TpjfGj7oaMBg.woff2
    │       ├── jizGREVNn1dOx-zrZ2X3pZvkTiUa4442m1bpjfGj7oaMBg.woff2
    │       ├── jizGREVNn1dOx-zrZ2X3pZvkTiUa4442m1fpjfGj7oaMBg.woff2
    │       └── jizGREVNn1dOx-zrZ2X3pZvkTiUa4442m1npjfGj7oY.woff2
    └── sourcecodepro
        └── v22
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvQlMIFxGC8NAU.woff2
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvUlMIFxGC8.woff2
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvXlMIFxGC8NAU.woff2
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvYlMIFxGC8NAU.woff2
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvZlMIFxGC8NAU.woff2
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvalMIFxGC8NAU.woff2
            ├── HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvblMIFxGC8NAU.woff2
            ├── HI_SiYsKILxRpg3hIP6sJ7fM7PqlM-vWnsUnxlC9.woff2
            ├── HI_SiYsKILxRpg3hIP6sJ7fM7PqlMOvWnsUnxlC9.woff2
            ├── HI_SiYsKILxRpg3hIP6sJ7fM7PqlMevWnsUnxlC9.woff2
            ├── HI_SiYsKILxRpg3hIP6sJ7fM7PqlMuvWnsUnxlC9.woff2
            ├── HI_SiYsKILxRpg3hIP6sJ7fM7PqlOevWnsUnxlC9.woff2
            ├── HI_SiYsKILxRpg3hIP6sJ7fM7PqlPevWnsUnxg.woff2
            └── HI_SiYsKILxRpg3hIP6sJ7fM7PqlPuvWnsUnxlC9.woff2
