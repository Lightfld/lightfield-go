# Changelog

## 0.4.1-alpha (2026-04-17)

Full Changelog: [v0.4.0-alpha...v0.4.1-alpha](https://github.com/Lightfld/lightfield-go/compare/v0.4.0-alpha...v0.4.1-alpha)

### Bug Fixes

* **api:** tighten object API schemas and remove recursive JsonValue from OpenAPI ([50eb8ac](https://github.com/Lightfld/lightfield-go/commit/50eb8ac0546d7c2c51ca7b4cf57b5ba3bd27c9e2))

## 0.4.0-alpha (2026-04-01)

Full Changelog: [v0.3.2-alpha...v0.4.0-alpha](https://github.com/Lightfld/lightfield-go/compare/v0.3.2-alpha...v0.4.0-alpha)

### Features

* **internal:** support comma format in multipart form encoding ([6c1ac2b](https://github.com/Lightfld/lightfield-go/commit/6c1ac2b645da5624342ffd6aa13bf2698cb89792))


### Bug Fixes

* prevent duplicate ? in query params ([fee8c86](https://github.com/Lightfld/lightfield-go/commit/fee8c862dee85bd5eca1da91e2b21aceaa9f9f36))


### Chores

* **ci:** skip lint on metadata-only changes ([7fa9b66](https://github.com/Lightfld/lightfield-go/commit/7fa9b669a8092772f68ac76f16f80c392356db41))
* **ci:** support opting out of skipping builds on metadata-only commits ([b04ed47](https://github.com/Lightfld/lightfield-go/commit/b04ed47c747aabf5bd42621263902b157bc730e7))
* **client:** fix multipart serialisation of Default() fields ([bc57e6b](https://github.com/Lightfld/lightfield-go/commit/bc57e6b350de76af5b88c200ca420bf67c01c262))
* **internal:** support default value struct tag ([7815165](https://github.com/Lightfld/lightfield-go/commit/7815165c6a32d54c7eee5824a800210619c3d0ca))
* **internal:** update gitignore ([0652d9f](https://github.com/Lightfld/lightfield-go/commit/0652d9f5502e92e34bbc0b9d2e015d78a6e61362))
* remove unnecessary error check for url parsing ([6d7b968](https://github.com/Lightfld/lightfield-go/commit/6d7b9682aa567df9348a014ea17aebaa0d9a9741))
* **tests:** bump steady to v0.19.4 ([49d9f32](https://github.com/Lightfld/lightfield-go/commit/49d9f3258f142d2af9161ea3606dbd4e4e2ce25b))
* **tests:** bump steady to v0.19.5 ([19a61c4](https://github.com/Lightfld/lightfield-go/commit/19a61c41f15de858dadce3eb14966ea0630efe02))
* **tests:** bump steady to v0.19.6 ([316f8f7](https://github.com/Lightfld/lightfield-go/commit/316f8f783b16e1dd489120a5c8febc04c6e91479))
* **tests:** bump steady to v0.19.7 ([50e6868](https://github.com/Lightfld/lightfield-go/commit/50e6868bea14dfb128c819c03ce370f7b97635df))
* **tests:** bump steady to v0.20.1 ([4a81542](https://github.com/Lightfld/lightfield-go/commit/4a815429f090e538643fded4635a2db1a2b360c4))
* **tests:** bump steady to v0.20.2 ([8d20577](https://github.com/Lightfld/lightfield-go/commit/8d2057778c157efa0d48febbf0523e658856c802))
* update docs for api:"required" ([d3b127d](https://github.com/Lightfld/lightfield-go/commit/d3b127d488f28c4c5107267e9955c26c29d5a148))

## 0.3.2-alpha (2026-03-20)

Full Changelog: [v0.3.1-alpha...v0.3.2-alpha](https://github.com/Lightfld/lightfield-go/compare/v0.3.1-alpha...v0.3.2-alpha)

### Bug Fixes

* **api:** resolve system and custom attribute collision error ([62a0eb5](https://github.com/Lightfld/lightfield-go/commit/62a0eb5a500426dc578f975eb005d21221d6fcea))

## 0.3.1-alpha (2026-03-20)

Full Changelog: [v0.3.0-alpha...v0.3.1-alpha](https://github.com/Lightfld/lightfield-go/compare/v0.3.0-alpha...v0.3.1-alpha)

### Chores

* **api:** add response model refs ([1e3a1ac](https://github.com/Lightfld/lightfield-go/commit/1e3a1acf910aa3e2fb2c46b5ba5ab7b2f9c1ce23))

## 0.3.0-alpha (2026-03-20)

Full Changelog: [v0.2.0...v0.3.0-alpha](https://github.com/Lightfld/lightfield-go/compare/v0.2.0...v0.3.0-alpha)

### Features

* **api:** add field/relationship param descriptions ([5ba8e7d](https://github.com/Lightfld/lightfield-go/commit/5ba8e7d9e6c3b12a8f1298def4b19cb7440c216a))
* **api:** manual updates ([0f6c7b9](https://github.com/Lightfld/lightfield-go/commit/0f6c7b9758907363cf5bd0de1721ad43bc485c59))
* **api:** manual updates ([e10ee46](https://github.com/Lightfld/lightfield-go/commit/e10ee4610017f096b1ece73bc026ab84d3b24c06))
* **api:** member and workflow run ([088de79](https://github.com/Lightfld/lightfield-go/commit/088de79fcd696f449732cba4903a5f416a9ae47a))
* **api:** update field descriptions, plural to singular keys, change READONLY_MARKDOWN type to MARKDOWN ([efb8037](https://github.com/Lightfld/lightfield-go/commit/efb8037fc297fa90dbb55df237bf1d64d4461ebe))


### Chores

* configure new SDK language ([c9184c1](https://github.com/Lightfld/lightfield-go/commit/c9184c137ffa4c21ca94cfd75cb72277a9296fe3))
* **internal:** codegen related update ([21b4113](https://github.com/Lightfld/lightfield-go/commit/21b411308e7fd831a33b437479e04b60a7991fa8))
* **internal:** tweak CI branches ([15774cc](https://github.com/Lightfld/lightfield-go/commit/15774cc3ea3327e98676e6a95e0228c4fa28b872))


### Refactors

* **tests:** switch from prism to steady ([ac726e9](https://github.com/Lightfld/lightfield-go/commit/ac726e9f4ef2b1b3d728808eca159fe6c851fffe))

## 0.2.0 (2026-03-17)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Lightfld/lightfield-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** shorten system attribute prefixes, add attribute definition endpoint ([f96eca1](https://github.com/Lightfld/lightfield-go/commit/f96eca1f0b0fdcc0a935fd03723dda6d9a1cc0da))

## 0.1.0 (2026-03-17)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/Lightfld/lightfield-go/compare/v0.0.2...v0.1.0)

### Features

* **api:** add Go SDK and CLI targets to Stainless config ([891280c](https://github.com/Lightfld/lightfield-go/commit/891280c5773e89eaf407ef656af84dc610ccd359))

## 0.0.2 (2026-03-17)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/Lightfld/lightfield-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([f3c1917](https://github.com/Lightfld/lightfield-go/commit/f3c1917ef35719f42b9cea95565bdcf77924e622))
* configure new SDK language ([c7a2815](https://github.com/Lightfld/lightfield-go/commit/c7a2815d3859ab44d07778d7afd45fce99929d85))
* update SDK settings ([b7c7d2c](https://github.com/Lightfld/lightfield-go/commit/b7c7d2c5110b08db3d0f0595dc4c895e9ebeb800))
